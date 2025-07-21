package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type SqlClient struct {
	db *sql.DB
}

func mySQLConnector(ctx context.Context, dbType, host, port, username, password, dbName string) (*SqlClient, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4",
		username, password, host, port, dbName)

	//log.Printf("Connecting to MySQL with DSN: %s", connectionString)

	db, err := sql.Open(dbType, connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to create sql connection: %w", err)
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to connect to sql: %w", err)
	}

	log.Println("Connected to database successfully.")

	return &SqlClient{db: db}, nil
}

// NewClientFromConfig creates a new SqlClient using package config variables
func NewClientFromConfig(ctx context.Context) (*SqlClient, error) {
	return mySQLConnector(ctx, DB_TYPE, DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)
}

// GetDatabaseDetails returns the MySQL version string
func (client *SqlClient) GetDatabaseDetails() string {
	var result string
	err := client.db.QueryRowContext(context.Background(), "SELECT @@version").Scan(&result)
	if err != nil {
		return fmt.Sprintf("Database scan failed: %v", err)
	}
	return strings.TrimSpace(result)
}

// GetDB exposes the underlying *sql.DB for direct queries if needed
func (client *SqlClient) GetDB() *sql.DB {
	return client.db
}
