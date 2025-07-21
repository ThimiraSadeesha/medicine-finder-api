package config

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

type SqlClient struct {
	db *sql.DB
}

func mySql(ctx context.Context, dbType, host, port, username, password, dbName string) (*SqlClient, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbName)

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

func (client *SqlClient) GetDatabaseDetails() string {
	var result string
	err := client.db.QueryRowContext(context.Background(), "SELECT @@version").Scan(&result)
	if err != nil {
		return fmt.Sprintf("Database scan failed: %v", err)
	}
	return strings.TrimSpace(result)
}

func (client *SqlClient) GetDB() *sql.DB {
	return client.db
}
