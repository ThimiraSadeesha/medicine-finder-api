package main

import (
	"fmt"
	"log"

	"medicine-finder-api/config"
	"medicine-finder-api/middleware"
	"medicine-finder-api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	addr := fmt.Sprintf("%s:%s", config.HOST, config.PORT)
	router := gin.Default()

	if err := router.SetTrustedProxies([]string{}); err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	router.Use(middleware.ResponseInterceptor())
	routes.RegisterRoutes(router)

	log.Printf("Starting server on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
