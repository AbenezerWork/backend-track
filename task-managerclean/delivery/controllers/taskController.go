package controllers

import (
	"net/http"
	"task_managerv2/domain"
	"task_managerv2/infrastructure"
	"task_managerv2/usecase"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskUsecase *usecase.TaskUsecase
}

func InitTaskController(taskUsecase *usecase.TaskUsecase) *TaskController {
	return &TaskController{TaskUsecase: taskUsecase}
}

func (tc *TaskController) GetTasksHandler(c *gin.Context) {
	user := c.MustGet("claims").(*infrastructure.MyCustomClaims)
	var rep []domain.Task
	if user.Role == "admin" {
		rep = tc.TaskUsecase.GetTasks("")
	} else {
		user_id := user.UserID
		rep = tc.TaskUsecase.GetTasks(user_id)
	}
	c.JSON(http.StatusOK, rep)
}

func (tc *TaskController) DeleteTaskHandler(c *gin.Context) {
	user := c.MustGet("claims").(*infrastructure.MyCustomClaims)
	user_id := user.UserID
	id := c.Param("id")

	rep, err := tc.TaskUsecase.GetTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	if user_id != rep.UserID && user.Role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You don't have access to the requested resource"})
		return
	}

	err = tc.TaskUsecase.DeleteTask(id)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.Status(http.StatusNoContent)
}

func (tc *TaskController) GetTaskHandler(c *gin.Context) {
	user := c.MustGet("claims").(*infrastructure.MyCustomClaims)
	user_id := user.UserID
	id := c.Param("id")

	rep, err := tc.TaskUsecase.GetTask(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user_id != rep.UserID && user.Role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You don't have access to the requested resource"})
		return

	}

	c.JSON(http.StatusOK, rep)
}

func (tc *TaskController) AddTaskHandler(c *gin.Context) {
	user := c.MustGet("claims").(*infrastructure.MyCustomClaims)
	user_id := user.UserID
	var newTask domain.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTask.UserID = user_id
	id, err := tc.TaskUsecase.AddTask(newTask)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "couldn't add task"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "task created successfully", "id": id})
}

func (tc *TaskController) UpdateTaskHandler(c *gin.Context) {
	user := c.MustGet("claims").(*infrastructure.MyCustomClaims)
	user_id := user.UserID
	id := c.Param("id")

	var newTask domain.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rep, err := tc.TaskUsecase.GetTask(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user_id != rep.UserID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You don't have access to the requested resource"})
		return
	}
	newTask.UserID = rep.UserID

	tc.TaskUsecase.UpdateTask(id, newTask)
	c.JSON(http.StatusOK, gin.H{"message": "task updated successfully", "id": id})

}
