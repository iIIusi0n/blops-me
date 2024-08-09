package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load("web/.env.local", "web/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DbHost = os.Getenv("MYSQL_HOST")
	DbPort = os.Getenv("MYSQL_PORT")
	DbUser = os.Getenv("MYSQL_USER")
	DbPassword = os.Getenv("MYSQL_PASSWORD")
	DbName = os.Getenv("MYSQL_DB_NAME")

	ServerHost = os.Getenv("SERVER_HOST")
	ServerPort = os.Getenv("SERVER_PORT")
	ReleaseMode = os.Getenv("SERVER_MODE")
	LogFile = os.Getenv("SERVER_LOG_PATH")

	ClientId = os.Getenv("OAUTH_CLIENT_ID")
	ClientSecret = os.Getenv("OAUTH_CLIENT_SECRET")
	RedirectUri = os.Getenv("OAUTH_REDIRECT_URI")
}
