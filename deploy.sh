#!/bin/bash

# Function to handle errors
handle_error() {
    echo -e "\e[31mERROR: $1\e[0m" >&2
    exit 1
}

# Function to check command success
check_command() {
    if [ $? -ne 0 ]; then
        handle_error "$1"
    fi
}

# Function for verbose output
verbose() {
    echo -e "\e[34m[$(date '+%Y-%m-%d %H:%M:%S')] $1\e[0m"
}

# Trap to catch unexpected errors
trap 'handle_error "An unexpected error occurred. Exiting..."' ERR

verbose "Starting deployment process..."

# Step 1: Install dependencies
verbose "Step 1: Installing dependencies"

verbose "Installing Go dependencies..."
go mod tidy
check_command "Failed to install Go dependencies"
verbose "Go dependencies installed successfully"

verbose "Installing Node.js dependencies..."
cd ./web || handle_error "Failed to change directory to ./web"
npm install
check_command "Failed to install Node.js dependencies"
verbose "Node.js dependencies installed successfully"
cd ..

verbose "Checking Caddy installation..."
if ! command -v caddy &> /dev/null; then
    verbose "Caddy not found. Installing Caddy..."
    sudo apt-get update && sudo apt-get install -y debian-keyring debian-archive-keyring apt-transport-https
    curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/gpg.key' | sudo gpg --dearmor -o /usr/share/keyrings/caddy-stable-archive-keyring.gpg
    curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/debian.deb.txt' | sudo tee /etc/apt/sources.list.d/caddy-stable.list
    sudo apt-get update
    sudo apt-get install caddy
    check_command "Failed to install Caddy"
    verbose "Caddy installed successfully"
else
    verbose "Caddy is already installed"
fi

# Step 2: Get configuration by user input
verbose "Step 2: Collecting configuration information"
read -p "Enter MySQL host: " MYSQL_HOST
read -p "Enter MySQL port: " MYSQL_PORT
read -p "Enter MySQL database name: " MYSQL_DB
read -p "Enter MySQL user: " MYSQL_USER
read -sp "Enter MySQL password: " MYSQL_PASSWORD
echo
read -p "Enter OAuth key: " OAUTH_KEY
read -p "Enter Client ID: " CLIENT_ID
read -p "Enter domain: " DOMAIN

# Validate inputs
[[ -z "$MYSQL_HOST" || -z "$MYSQL_PORT" || -z "$MYSQL_DB" || -z "$MYSQL_USER" || -z "$MYSQL_PASSWORD" || -z "$OAUTH_KEY" || -z "$CLIENT_ID" || -z "$DOMAIN" ]] && handle_error "All fields are required"
verbose "All configuration information collected successfully"

# Step 3: Setup database DDL
verbose "Step 3: Setting up database"
verbose "Executing SQL schema..."
mysql -h "$MYSQL_HOST" -P "$MYSQL_PORT" -u "$MYSQL_USER" -p"$MYSQL_PASSWORD" "$MYSQL_DB" < ./schemes.sql
check_command "Failed to set up database"
verbose "Database setup completed successfully"

# Step 4: Change configurations by inputs
verbose "Step 4: Updating configuration files"

# Function to safely update configuration files
update_config() {
    local file=$1
    local search=$2
    local replace=$3
    if [ ! -f "$file" ]; then
        handle_error "Configuration file $file not found"
    fi
    verbose "Updating $file: Replacing $search with $replace"
    sed -i "s|$search|$replace|g" "$file"
    check_command "Failed to update $file"
    verbose "$file updated successfully"
}

update_config "./config/server.go" "<MYSQL_HOST>" "$MYSQL_HOST"
update_config "./config/server.go" "<MYSQL_PORT>" "$MYSQL_PORT"
update_config "./config/server.go" "<MYSQL_DB>" "$MYSQL_DB"
update_config "./config/server.go" "<MYSQL_USER>" "$MYSQL_USER"
update_config "./config/server.go" "<MYSQL_PASSWORD>" "$MYSQL_PASSWORD"

update_config "./config/oauth.go" "<OAUTH_KEY>" "$OAUTH_KEY"
update_config "./config/oauth.go" "<CLIENT_ID>" "$CLIENT_ID"

update_config "./Caddyfile" "<DOMAIN>" "$DOMAIN"

verbose "All configuration files updated successfully"

# Step 5: Build backend Golang and register to service, start it
verbose "Step 5: Building and deploying backend"
verbose "Building backend..."
cd ./cmd/api-server || handle_error "Failed to change directory to ./cmd/api-server"
go build -o api-server main.go
check_command "Failed to build backend"
verbose "Backend built successfully"

verbose "Moving api-server to /usr/local/bin/"
sudo mv api-server /usr/local/bin/
check_command "Failed to move api-server to /usr/local/bin/"
verbose "api-server moved successfully"
cd ../..

verbose "Creating systemd service for backend..."
sudo tee /etc/systemd/system/api-server.service > /dev/null <<EOF
[Unit]
Description=API Server

[Service]
ExecStart=/usr/local/bin/api-server
Restart=always
User=nobody
Group=nogroup
Environment=PATH=/usr/bin:/usr/local/bin
Environment=GO_ENV=production
WorkingDirectory=/usr/local/bin

[Install]
WantedBy=multi-user.target
EOF
check_command "Failed to create api-server systemd service"
verbose "api-server systemd service created successfully"

verbose "Starting api-server service..."
sudo systemctl daemon-reload
sudo systemctl enable api-server
sudo systemctl start api-server
check_command "Failed to start api-server service"
verbose "api-server service started successfully"

# Step 6: Build Next.js and run as service, start it
verbose "Step 6: Building and deploying frontend"
verbose "Building frontend..."
cd ./web || handle_error "Failed to change directory to ./web"
npm run build
check_command "Failed to build frontend"
verbose "Frontend built successfully"

verbose "Creating systemd service for frontend..."
sudo tee /etc/systemd/system/nextjs.service > /dev/null <<EOF
[Unit]
Description=Next.js Frontend

[Service]
ExecStart=/usr/bin/npm start
Restart=always
User=nobody
Group=nogroup
Environment=PATH=/usr/bin:/usr/local/bin
Environment=NODE_ENV=production
WorkingDirectory=$(pwd)

[Install]
WantedBy=multi-user.target
EOF
check_command "Failed to create nextjs systemd service"
verbose "nextjs systemd service created successfully"

verbose "Starting nextjs service..."
sudo systemctl daemon-reload
sudo systemctl enable nextjs
sudo systemctl start nextjs
check_command "Failed to start nextjs service"
verbose "nextjs service started successfully"
cd ..

# Step 7: Start Caddy with Caddyfile
verbose "Step 7: Starting Caddy server"
verbose "Starting Caddy..."
sudo caddy start --config ./Caddyfile
check_command "Failed to start Caddy"
verbose "Caddy started successfully"

verbose "Deployment completed successfully!"
