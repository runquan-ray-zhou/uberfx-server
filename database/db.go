package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	InitDB,
)

var DB *sql.DB

func InitDB() *sql.DB {
	var err error

	// Option 1: use DATABASE_URL from environment (e.g., Render)
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// Option 2: use individual connection values from .env
		host := os.Getenv("PG_HOST")
		port := os.Getenv("PG_PORT")
		user := os.Getenv("PG_USER")
		database := os.Getenv("PG_DATABASE")
		dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", // change to sslmode=require in production
			host, port, user, database)
	}

	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Could not open database: %v", err)
	}

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := DB.PingContext(ctx); err != nil {
		log.Fatalf("Could not ping database: %v", err)
	}

	// Set connection pool options
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	// Run schema migration or setup
	if err := createTables(DB); err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}

	return DB
}

func createTables(DB *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		created_at TIMESTAMPTZ DEFAULT now()
	);
	`
	_, err := DB.Exec(query)
	return err
}
