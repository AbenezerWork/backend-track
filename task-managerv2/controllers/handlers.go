package controllers

import (
	"net/http"
	"strconv"
	"task_managerv2/model"

	"github.com/gin-gonic/gin"
)

func GetTasksHandler(c *gin.Context) {
	rep := GetTasks()
	c.JSON(http.StatusOK, rep)
}

func DeleteTaskHandler(c *gin.Context) {
	id := c.Param("id")

	idd, err := strconv.Atoi(id)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	err = DeleteTask(idd)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.Status(http.StatusNoContent)
}

func GetTaskHandler(c *gin.Context) {
	id := c.Param("id")

	idd, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rep, err := GetTask(idd)
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

	idd, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var newTask model.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	UpdateTask(idd, newTask)
	c.JSON(http.StatusOK, gin.H{"message": "task updated successfully", "id": id})

}
