package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func DatabaseConnect() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file")
	}

	dbUrl := buildDbUrl()

	db, err := sql.Open("libsql", dbUrl)
	if err != nil {
		log.Fatalf("Failed to open db %s: %s", dbUrl, err)
	}
	return db
}

func buildDbUrl() string {
	dbToken := os.Getenv("DB_TOKEN")

	return "libsql://mini-mind-adriano-henrique.turso.io?authToken=" + dbToken
}
