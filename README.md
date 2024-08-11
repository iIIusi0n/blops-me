# Blops Me
### Automatic file organizer powered by AI.
As you use your computer every day, you'll end up with lots of files. But organising them and creating the right folder trees for them can be a pain. Gemini's natural language processing capabilities mean it doesn't have to be a problem.

## Deploy
### Support systems
- Windows
- MacOS
- Linux

### Dependencies
- Golang 1.20+
- Caddy : You can install it via 'apt'
- Python : **'python' command should be worked in terminal**
- Node.js
- MySQL

### 1. API keys
You will need Google OAuth credentials for the login system.

You will also need a Gemini API key. (Free is ok, it will automatically limit requests to the given tariff).

### 2. Clone the repo
```bash
git clone https://github.com/iIIusi0n/blops-me.git
cd blops-me
```

### 3. Setup MySQL database
Create a new database with 'schemes.sql' and user with privileges.

```sql
CREATE DATABASE my_new_database;

CREATE USER 'new_user'@'localhost' IDENTIFIED BY 'user_password';
GRANT ALL PRIVILEGES ON my_new_database.* TO 'new_user'@'localhost';

FLUSH PRIVILEGES;

SOURCE assets/schemes.sql
```

### 3. Edit configurations
There is an .env file in web/, but it is recommended to make .env.local and use that.

Then edit host in caddyfile.

```bash
# Edit the Caddyfile as shown in the example below
vim Caddyfile

# Edit environmental variables
cp web/.env web/.env.local
vim web/.env.local
```

Set MYSQL_* variables, change APP_URL to url for access. (It should match the caddyfile)

Set OAUTH_CLIENT_ID and OAUTH_CLIENT_SECRET with credentials, also change GEMINI_API_KEY.

You will also need to configure SESSION_SECRET with a 32-byte random string.

Example Caddyfile:
```
example.com {
    reverse_proxy /auth/* localhost:8010
    reverse_proxy /api/* localhost:8010
    reverse_proxy /* localhost:3000
}
```

### 4. Run API server
```bash
go build blops-me/cmd/api-server
./api-server

# It will run as a foreground job, so you can use the screen like this.
screen -S api ./api-server
```

### 5. Run Node.js frontend
```bash
cd web
npm install
npm run build
npm run start

# It will run as a foreground job, so you can use the screen like this.
screen -S web npm run start

cd ..
```

### 6. Run Caddy proxy
```bash
# This will run as a background job. If you want to monitor the status, you can use run instead of start.
caddy start --config Caddyfile
```

## TODO
- [x] Add Dockerfile
- [ ] FIX   SPAGHTTI
