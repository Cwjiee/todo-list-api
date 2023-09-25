package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title		string `json:"title"`
	Description	string `json:"description"`
}

func main() {

	router := gin.Default()

	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Todo{})

	router.POST("/todos", func(ctx *gin.Context) {
		var todo Todo
		if err := ctx.ShouldBindJSON(&todo); err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid JSON data"})
			return
		}

		db.Create(&todo)
		ctx.JSON(200, todo)
	})

	router.GET("/todos", func(ctx *gin.Context) {
		var todos []Todo

		db.Find(&todos)

		ctx.JSON(200, todos)
	})

	router.GET("/todos/:id", func(ctx *gin.Context) {
		var todo Todo
		todoID := ctx.Param("id")

		result := db.First(&todo, todoID)
		if result.Error != nil {
			ctx.JSON(404, gin.H{"error": "data not found"})
			return
		}

		ctx.JSON(200, todo)
	})

	router.PUT("/todos/:id", func(ctx *gin.Context) {
		var todo Todo
		todoID := ctx.Param("id")

		result := db.Find(&todo, todoID)
		if result.Error != nil {
			ctx.JSON(404, gin.H{"error": "data not found"})
			return
		}

		var updatedTodo Todo
		if err := ctx.ShouldBindJSON(&updatedTodo); err != nil {
			ctx.JSON(400, gin.H{"error": "invalid JSON data"})
			return
		}

		todo.Title = updatedTodo.Title
		todo.Description = updatedTodo.Description
		db.Save(&todo)

		ctx.JSON(200, todo)
	})

	router.DELETE("/todos/:id", func(ctx *gin.Context) {
		var todo Todo
		todoId := ctx.Param("id")

		result := db.First(&todo, todoId)
		if result.Error != nil {
			ctx.JSON(404, gin.H{"error": "data not found"})
			return
		}

		db.Delete(&todo)

		ctx.JSON(200, gin.H{"message": fmt.Sprintf("Todo with ID %s is deleted", todoId)})
	})

	router.Run("localhost:8080")
}