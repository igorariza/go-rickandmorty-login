package seed

import (
	"log"

	"github.com/igorariza/go-rickandmorty-login/api/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{}

func Load(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.Character{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
}
