package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		InitDB,
	),
	fx.Invoke(func(db *sql.DB) {
		fmt.Println("Database initialized")
	}),
)

var DB *sql.DB

func InitDB() (*sql.DB, error) {
	var err error

	// load .env file with godotenv
	err = godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("no .env file found: %w", err)
	}

	// connect to render when database is online
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// use individual connection values from .env
		host := os.Getenv("PG_HOST")
		port := os.Getenv("PG_PORT")
		user := os.Getenv("PG_USER")
		database := os.Getenv("PG_DATABASE")
		dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", // change to sslmode=require in production
			host, port, user, database)
	}

	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("could not open database: %w", err)
	}

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = DB.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}

	// Set connection pool options
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	// createTable
	err = createTables(DB)
	if err != nil {
		return nil, fmt.Errorf("failed to create tables: %w", err)
	}

	return DB, nil
}

func createTables(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		created_at TIMESTAMPTZ DEFAULT now()
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to exec query: %w", err)
	}
	return nil
}
