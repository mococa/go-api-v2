package db

import (
	"fmt"
	"log"

	"github.com/mococa/go-api-v2/pkg/db/models"
	"github.com/mococa/go-api-v2/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() *gorm.DB {
	env, err := utils.LoadConfig()
	if err != nil {
		panic("ENV error")
	}

	DB_URI := env.DB_URI
	db, err := gorm.Open(postgres.Open(DB_URI), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database")

	runMigrations(db)

	return db
}

func GetDB() *gorm.DB {
	return db
}

func runMigrations(db *gorm.DB) {
	fmt.Println("Running migrations...")

	db.AutoMigrate(
		/* -------------- Books -------------- */
		&models.Book{},

		/* -------------- Authors -------------- */
		&models.Author{},
	)
}
