package server

import (
	cAuth "blops-me/controllers/auth"
	cStorage "blops-me/controllers/storage"
	"blops-me/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) *gin.Engine {
	auth := r.Group("/auth")
	{
		google := auth.Group("/google")
		{
			google.GET("/login", cAuth.LoginHandler)
			google.GET("/callback", cAuth.CallbackHandler)
		}

		auth.GET("/verify", cAuth.VerifyHandler)
	}

	api := r.Group("/api")
	{
		api.Use(middlewares.ProtectedRouter)

		api.GET("/storage", cStorage.ListStorageHandler)
		api.POST("/storage", cStorage.CreateStorageHandler)
		api.DELETE("/storage", cStorage.DeleteStorageHandler)

		api.GET("/storage/:id")
		api.POST("/storage/:id")
		api.DELETE("/storage/:id")
	}

	return r
}
