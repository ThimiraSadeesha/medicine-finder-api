package middleware

import (
	"github.com/gin-gonic/gin"
	_ "net/http"
	"time"
)

func ResponseInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if c.Writer.Written() || len(c.Errors) > 0 {
			return
		}
		status := c.Writer.Status()
		data, exists := c.Get("responseData")
		if !exists {
			data = gin.H{}
		}

		c.JSON(status, gin.H{
			"response":  status,
			"path":      c.Request.URL.Path,
			"timestamp": time.Now().Format("2006-01-02 15:04:05"),
			"data":      data,
		})
	}
}
