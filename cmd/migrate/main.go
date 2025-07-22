package main

import (
	"context"
	"log"
	"medicine-finder-api/config"
	"medicine-finder-api/migrations"
)

func main() {
	ctx := context.Background()

	dbClient, err := config.NewClientFromConfig(ctx)
	if err != nil {
		log.Fatalf("❌ DB connection failed: %v", err)
	}

	migrations.RunMigrations(dbClient.GetGormDB())
}
