package controllers

import (
	"fmt"
	"log"
	"net/http"
	"github.com/Cwjiee/todo-list-api/src/models"
	"github.com/gin-gonic/gin"
)

type todoRequest struct {
	Title        string `json:"title"`
	Description string `json:"description"`
}

func CreateTodo(ctx *gin.Context) {
	var data todoRequest

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid JSON data"})
		return
	}

	todo := models.Todo{}
	todo.Title = data.Title
	todo.Description = data.Description

	db.Create(&todo)
	ctx.JSON(http.StatusCreated, todo)
}

func GetTodos(ctx *gin.Context) {
	var todos []models.Todo

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}

	db.Find(&todos)
	ctx.JSON(http.StatusOK, todos)
}

func GetTodo(ctx *gin.Context) {
	todo := models.Todo{}
	todoID := ctx.Param("id")

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}

	result := db.First(&todo, todoID)
	if result.Error != nil {
		ctx.JSON(404, gin.H{"error": "data not found"})
		return
	}

	ctx.JSON(200, todo)
}

func UpdateTodo(ctx *gin.Context) {
	todo := models.Todo{}
	todoID := ctx.Param("id")

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}

	result := db.Find(&todo, todoID)
	if result.Error != nil {
		ctx.JSON(404, gin.H{"error": "data not found"})
		return
	}

	var updatedTodo todoRequest
	if err := ctx.ShouldBindJSON(&updatedTodo); err != nil {
		ctx.JSON(400, gin.H{"error": "invalid JSON data"})
		return
	}

	todo.Title = updatedTodo.Title
	todo.Description = updatedTodo.Description
	db.Save(&todo)

	ctx.JSON(200, todo)
}

func DeleteTodo(ctx *gin.Context) {
	todo := models.Todo{}
	todoId := ctx.Param("id")

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}

	result := db.First(&todo, todoId)
	if result.Error != nil {
		ctx.JSON(404, gin.H{"error": "data not found"})
		return
	}

	db.Delete(&todo)

	ctx.JSON(200, gin.H{"message": fmt.Sprintf("Todo with ID %s is deleted", todoId)})
}