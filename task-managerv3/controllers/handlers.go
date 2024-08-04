package controllers

import (
	"log"
	"net/http"
	"task_managerv2/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTasksHandler(c *gin.Context) {
	rep := GetTasks()
	c.JSON(http.StatusOK, rep)
}

func DeleteTaskHandler(c *gin.Context) {
	id := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	err = DeleteTask(objID)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.Status(http.StatusNoContent)
}

func GetTaskHandler(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	rep, err := GetTask(objID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rep)
}

func AddTaskHandler(c *gin.Context) {
	var newTask model.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := AddTask(newTask)
	c.JSON(http.StatusCreated, gin.H{"message": "task created successfully", "id": id})
}

func UpdateTaskHandler(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var newTask model.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	UpdateTask(objID, newTask)
	c.JSON(http.StatusOK, gin.H{"message": "task updated successfully", "id": id})

}
