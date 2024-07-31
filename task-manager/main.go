package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

var tasks = []Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

func putTask(c *gin.Context) {
	id := c.Param("id")

	var updatedTask Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	for _, task := range tasks {
		if id == task.ID {
			if updatedTask.Description != "" {
				task.Description = updatedTask.Description
			}

			if updatedTask.Title != "" {
				task.Title = updatedTask.Title
			}

			if updatedTask.DueDate != (time.Time{}) {
				task.DueDate = updatedTask.DueDate
			}
			c.JSON(http.StatusOK, gin.H{"message": "task updated"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})

}

func deleteHandler(c *gin.Context) {
	id := c.Param("id")

	for i, val := range tasks {
		if val.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func postHandler(c *gin.Context) {
	var newTask Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	tasks = append(tasks, newTask)

	c.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}

func main() {
	router := gin.Default()
	router.GET("/tasks", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"tasks": tasks})
	})
	router.GET("/tasks/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, task := range tasks {
			if task.ID == id {
				c.JSON(http.StatusOK, gin.H{"task": task})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})

	})
	router.PUT("/tasks/:id", putTask)
	router.POST("/tasks", postHandler)

	router.Run("localhost:8080")
}
