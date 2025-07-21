package main

import (
	"fmt"
	"log"
	"medicine-finder-api/config"
	"medicine-finder-api/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	addr := fmt.Sprintf("%s:%s", config.HOST, config.PORT)
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	router.Use(middleware.ResponseInterceptor())
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello from Go API!"})
	})
	router.GET("/example", func(c *gin.Context) {
		result := gin.H{"message": "Hello!"}
		c.Set("responseData", result)
	})

	log.Printf("Starting server on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
