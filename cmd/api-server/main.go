package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	c "blops-me/config"
	"blops-me/data"
	"blops-me/middlewares"
	"blops-me/server"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetLogOutput(logFilePath string) (*os.File, error) {
	logDir := filepath.Dir(logFilePath)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, err
	}

	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	log.SetOutput(gin.DefaultWriter)

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("[GIN-debug] %v |%v %v %v %v\n", time.Now().Format("2006/01/02 - 15:04:05"), httpMethod, absolutePath, handlerName, nuHandlers)
	}

	return logFile, nil
}

func main() {
	gin.SetMode(c.ReleaseMode)

	logFile, err := SetLogOutput(c.LogFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {

		}
	}(logFile)
	log.Println("Log file created")

	dbConn, err := data.GetDatabaseConn(data.DBConfig{
		Host:     c.DbHost,
		Port:     c.DbPort,
		User:     c.DbUser,
		Password: c.DbPassword,
		DBName:   c.DbName,
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer func(dbConn *sql.DB) {
		err := dbConn.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(dbConn)

	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8MB

	r.Use(middlewares.AddDatabaseConnToContext(dbConn))
	log.Println("Database connection added to context")

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8080"}
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "storage-id")
	r.Use(cors.New(corsConfig))

	r.Use(middlewares.AuthMiddleware())

	server.SetupRouter(r)

	log.Fatalln(r.Run(fmt.Sprintf("%s:%s", c.ServerHost, c.ServerPort)))
}
