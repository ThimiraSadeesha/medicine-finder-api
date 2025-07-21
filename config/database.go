package config

import (
	"context"
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// SqlClient Configurations for the database connection
type SqlClient struct {
	db     *sql.DB
	GormDB *gorm.DB
}

func NewClientFromConfig(ctx context.Context) (*SqlClient, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)

	sqlDB, err := sql.Open(DB_TYPE, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to create sql connection: %w", err)
	}

	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gorm: %w", err)
	}

	log.Println("✅ Connected to database successfully.")

	return &SqlClient{
		db:     sqlDB,
		GormDB: gormDB,
	}, nil
}

// raw *sql.DB
func (client *SqlClient) GetDB() *sql.DB {
	return client.db
}

// GetGormDB GORM *gorm.DB
func (client *SqlClient) GetGormDB() *gorm.DB {
	return client.GormDB
}

// GetDatabaseDetails
func (client *SqlClient) GetDatabaseDetails() string {
	var version string
	err := client.db.QueryRow("SELECT @@version").Scan(&version)
	if err != nil {
		return "Error fetching DB version: " + err.Error()
	}
	return version
}
