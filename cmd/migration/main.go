package main

import (
	"fmt"

	"github.com/velo-project/events-service/adapters/dao/tables"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Connection to Db
	// TODO: Implement Env Variables to connect to the database
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost", "postgres", "1234", "velo_project", "5432",
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
