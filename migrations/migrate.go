package migrations

import (
	"gorm.io/gorm"
	"log"
	"medicine-finder-api/schema"
)

func RunMigrations(db *gorm.DB) {
	log.Println("🔧 Running database migrations...")

	err := db.AutoMigrate(
		&schema.User{},
		&schema.Drug{},
	)
	if err != nil {
		log.Fatalf("❌ Failed to run migrations: %v", err)
	}

	log.Println("✅ Migrations completed successfully.")
}
