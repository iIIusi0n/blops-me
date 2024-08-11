# Use Ubuntu as the base image
FROM ubuntu:latest

# Set environment variables to avoid prompts during installation
ENV DEBIAN_FRONTEND=noninteractive

# Install dependencies
RUN apt-get update && apt-get install -y \
    golang \
    caddy \
    python3 \
    python3-pip \
    python3-venv \
    python3-virtualenv \
    nodejs \
    npm \
    mysql-server \
    git \
    screen \
    && apt-get clean

# Set the working directory
WORKDIR /app

# Copy all files from the current directory to the container's working directory
COPY . .

RUN apt-get install -y libmagic1 libmagic-dev

# Install Python dependencies
RUN virtualenv /opt/venv
RUN /opt/venv/bin/pip install python-magic
RUN /opt/venv/bin/pip install pypdf
RUN /opt/venv/bin/pip install python-docx
RUN /opt/venv/bin/pip install python-pptx
RUN /opt/venv/bin/pip install xlrd
RUN /opt/venv/bin/pip install openpyxl

# Install Node.js dependencies
WORKDIR /app/web
RUN npm install
RUN npm run build

# Go back to the root directory
WORKDIR /app

RUN go build blops-me/cmd/api-server

# Set up MySQL database
COPY assets/schemes.sql /docker-entrypoint-initdb.d/
RUN service mysql start && \
    mysql -e "CREATE DATABASE blops_me;" && \
    mysql -e "CREATE USER 'blops_me'@'localhost' IDENTIFIED BY '12345678';" && \
    mysql -e "GRANT ALL PRIVILEGES ON blops_me.* TO 'blops_me'@'localhost';" && \
    mysql -e "FLUSH PRIVILEGES;" && \
    mysql blops_me < /docker-entrypoint-initdb.d/schemes.sql

# Expose the ports for API server and Caddy
EXPOSE 80 443

# Run the API server and frontend
CMD service mysql start && \
    screen -dmS api ./api-server && \
    cd web && \
    screen -dmS web npm run start && \
    cd .. && \
    caddy run --config Caddyfile
