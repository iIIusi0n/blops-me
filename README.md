# Blops Me
### Automatic file organizer powered by AI.
As you use your computer every day, you'll end up with lots of files. But organising them and creating the right folder trees for them can be a pain. Gemini's natural language processing capabilities mean it doesn't have to be a problem.

## Deploy
### Support systems
- Windows
- MacOS
- Linux

### Dependencies
- Docker

### 1. API keys
You will need Google OAuth credentials for the login system.

You will also need a Gemini API key. (Free is ok, it will automatically limit requests to the given tariff).

### 2. Clone the repo
```bash
git clone https://github.com/iIIusi0n/blops-me.git
cd blops-me
```

### 3. Edit configurations
There is an .env file in web/, but it is recommended to make .env.local and use that.

Please fill following variables and do not touch other if you want to install using docker.
```bash
- APP_DOMAIN

- OAUTH_CLIENT_ID
- OAUTH_CLIENT_SECRET
- SESSION_SECRET

- GEMINI_API_KEY
```

Then edit host in caddyfile.

```bash
# Edit the Caddyfile as shown in the example below
vim Caddyfile

# Edit environmental variables
cp web/.env web/.env.local
vim web/.env.local
```

Example Caddyfile:
```
example.com {
    reverse_proxy /auth/* localhost:8010
    reverse_proxy /api/* localhost:8010
    reverse_proxy /* localhost:3000
}
```

### 4. Build and run
```bash
# Build docker image
docker build -t blops-me .

# Run docker container
docker run -d -p 80:80 -p 443:443 blops-me
```

## TODO
- [x] Add Dockerfile
- [ ] FIX   SPAGHTTI
