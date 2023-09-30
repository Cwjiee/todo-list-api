package main

import (
	"log"
	"github.com/Cwjiee/todo-list-api/src/models"
	"github.com/Cwjiee/todo-list-api/src/routes"
)

func main() {
	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}
	db.DB()

	routes.Routes()
}
