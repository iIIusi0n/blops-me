package middlewares

import (
	"blops-me/controllers/auth"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")
		if err != nil {
			c.Set("authorized", false)
			c.Set("user", "")
			c.Next()
			return
		}

		ok, user, err := auth.VerifyToken(cookie)
		if err != nil {
			c.Set("authorized", false)
			c.Set("user", "")
			c.Next()
			return
		}

		c.Set("authorized", ok)
		c.Set("user", user)
		c.Next()
	}
}

func ProtectedRouter(c *gin.Context) {
	authorized := c.GetBool("authorized")
	if !authorized {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	c.Next()
}
