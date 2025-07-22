package main

import (
	"context"
	"fmt"
	"log"
	"medicine-finder-api/config"
	"medicine-finder-api/middleware"
	//"medicine-finder-api/migrations"
	"medicine-finder-api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	ctx := context.Background()

	// Connect to database on startup
	dbClient, err := config.NewClientFromConfig(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("DB Version:", dbClient.GetDatabaseDetails())

	// Run migrations BEFORE starting the server
	//migrations.Run(dbClient.GetGormDB())
	//
	addr := fmt.Sprintf("%s:%s", config.HOST, config.PORT)

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

	// Register your routes, pass dbClient (with both gormDB and sqlDB)
	routes.RegisterRoutes(router, dbClient)

	log.Printf("Starting server on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
