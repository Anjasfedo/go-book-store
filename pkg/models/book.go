package models

import (
	"github.com/Anjasfedo/go-book-store/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()

	db = config.GetDB()

	db.AutoMigrate(&Book{})
}
