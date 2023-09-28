package models

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title		string `json:"title"`
	Description	string `json:"description"`
}

func Database() (*gorm.DB, error) {

	db, err := gorm.Open(sqlite.Open("./database.db"), &gorm.Config{})

	if err != nil {
			log.Fatal(err)
	}

	if err = db.AutoMigrate(&Todo{}); err != nil {
			log.Println(err)
	}

	return db, err

}