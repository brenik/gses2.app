package main

import (
	"github.com/brenik/gses2.app/internal/server"
	"github.com/brenik/gses2.app/internal/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	err = runMigrations()
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	port := os.Getenv("PORT")
	service.StartSending()
	server.Run(port)

}

func runMigrations() error {
	dsn_migration := os.Getenv("DSNM")

	m, err := migrate.New(
		"file://migrations",
		dsn_migration)
	if err != nil {
		return err
	}

	// Застосування міграцій
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
