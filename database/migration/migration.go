package migration

import (
	"fmt"
	"log"

	"github.com/farinchan/golang_backend/config"
	"github.com/farinchan/golang_backend/models"
)

func RunMigration() {
	err := config.DB.AutoMigrate(&models.User{})

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database Migrated")
}
