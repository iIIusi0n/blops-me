package middlewares

import (
	"blops-me/internal/gemini"
	"github.com/gin-gonic/gin"
)

func AddGeminiClientQueueToContext(cq *gemini.ClientQueue) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("geminiClient", cq)
		c.Next()
	}
}
