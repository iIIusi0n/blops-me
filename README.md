# Blops Me
### Automatic file organizer powered by AI.
As you use your computer every day, you'll end up with lots of files. But organising them and creating the right folder trees for them can be a pain. Gemini's natural language processing capabilities mean it doesn't have to be a problem.

## Deploy
### Dependencies
- Docker

### 1. API keys
You will need Google OAuth credentials for the login system.

You will also need a Gemini API key. (Free is ok, it will automatically limit requests to the given limit).

### 2. Clone the repo
```bash
git clone https://github.com/iIIusi0n/blops-me.git
cd blops-me
```

### 3. Edit configurations
1) Edit environmental variables in docker-compose.yml.
```bash
vim docker-compose.yml
```

### 4. Build and run
1) Run containers using docker-compose.
```bash
docker-compose up -d
```

## TODO
- [x] Add Dockerfile
- [ ] FIX   SPAGHTTI
