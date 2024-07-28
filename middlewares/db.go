package middlewares

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func AddDatabaseConnToContext(conn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", conn)
		c.Next()
	}
}
