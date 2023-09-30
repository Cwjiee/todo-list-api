package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Cwjiee/todo-list-api/src/controllers"
	"log"
)

func Routes() {
	route := gin.Default()

	log.Printf("route works fine")
	route.POST("/todos", controllers.CreateTodo)
	route.GET("/todos", controllers.GetTodos)
	route.GET("/todos/:id", controllers.GetTodo)
	route.PUT("/todos/:id", controllers.UpdateTodo)
	route.DELETE("/todos/:id", controllers.DeleteTodo)

	route.Run("localhost:8080")
}

