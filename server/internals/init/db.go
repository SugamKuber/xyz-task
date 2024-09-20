package init

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var db *pgx.Conn

func InitDB() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL not set in environment")
	}

	var err error
	db, err = pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	createTableSQL := `
        CREATE TABLE IF NOT EXISTS cars (
            id SERIAL PRIMARY KEY,
            image BYTEA,
            type VARCHAR(50),
            color VARCHAR(50),
            make VARCHAR(50),
            model VARCHAR(50),
            caption TEXT
        );
    `
	_, err = db.Exec(context.Background(), createTableSQL)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	fmt.Println("Database initialized and table created if not exists.")
}

func GetDB() *pgx.Conn {
	return db
}
