package router

import (
	"task_managerv2/controllers"

	"github.com/gin-gonic/gin"
)

func Route() {
	router := gin.Default()

	router.GET("/tasks", controllers.GetTasksHandler)
	router.GET("/tasks/:id", controllers.GetTaskHandler)
	router.PUT("/tasks/:id", controllers.UpdateTaskHandler)
	router.DELETE("/tasks/:id", controllers.DeleteTaskHandler)
	router.POST("/tasks", controllers.AddTaskHandler)
	router.Run("localhost:8080")
}
