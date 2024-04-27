package config

import (
	"app/migrations"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf("host=%s  user=%s  password=%s  dbname=%s  port=%s ",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	if err := migrations.Migrate(db); err != nil {
		return err
	}

	DB = db
	log.Print("Connected")
	return nil
}

func Disconnect() error {
	if DB != nil {
		db, err := DB.DB()
		if err != nil {
			return err
		}
		if err := db.Close(); err != nil {
			return err
		}
	}
	log.Print("Disconnected")
	return nil
}
