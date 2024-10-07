package config

import (
	"os"
)

func init() {
	DbHost = os.Getenv("MYSQL_HOST")
	DbPort = os.Getenv("MYSQL_PORT")
	DbUser = "blops_me"
	DbPassword = "12345678"
	DbName = "blops_me"

	ServerHost = "0.0.0.0"
	ServerPort = "8010"
	ReleaseMode = os.Getenv("SERVER_MODE")
	LogFile = os.Getenv("SERVER_LOG_PATH")

	ClientId = os.Getenv("OAUTH_CLIENT_ID")
	ClientSecret = os.Getenv("OAUTH_CLIENT_SECRET")
	RedirectUri = os.Getenv("OAUTH_REDIRECT_URI")
	SessionSecret = os.Getenv("SESSION_SECRET")

	GeminiApiKey = os.Getenv("GEMINI_API_KEY")
}
