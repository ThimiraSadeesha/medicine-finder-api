package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to Medicine Finder API"})
	})

	// Add your module routes here
	//controllers.RegisterUserRoutes(api)
	//controllers.RegisterProductRoutes(api)
	// ... add other route registrations
}
