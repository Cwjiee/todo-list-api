package main

import (
	"log"
	"github.com/Cwjiee/todo-list-api/src/models"
	"github.com/Cwjiee/todo-list-api/src/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}
	db.DB()

	route := gin.Default()

	log.Printf("route works fine")
	route.POST("/todos", controllers.CreateTodo)
	route.GET("/todos", controllers.GetTodos)
	route.GET("/todos/:id", controllers.GetTodo)
	route.PUT("/todos/:id", controllers.UpdateTodo)
	route.DELETE("/todos/:id", controllers.DeleteTodo)

	route.Run("localhost:8080")
}
