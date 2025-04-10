package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/velo-project/events-service/adapters/dao/tables"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Connection to Db
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	postgresServer := os.Getenv("POSTGRES_URI")
	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPwd := os.Getenv("POSTGRES_PASSWORD")
	postgresDb := os.Getenv("POSTGRES_DATABASE")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		postgresServer, postgresUser, postgresPwd, postgresDb, postgresPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&tables.EventTable{})
	if err != nil {
		panic(err)
	}
}
