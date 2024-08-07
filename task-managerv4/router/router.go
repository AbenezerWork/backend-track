package router

import (
	"task_managerv2/controllers"
	"task_managerv2/middleware"

	"github.com/gin-gonic/gin"
)

func Route() {
	router := gin.Default()

	router.GET("/tasks", middleware.UserAuthMiddleware(), controllers.GetTasksHandler)
	router.GET("/tasks/:id", middleware.UserAuthMiddleware(), controllers.GetTaskHandler)
	router.PUT("/tasks/:id", middleware.UserAuthMiddleware(), controllers.UpdateTaskHandler)
	router.DELETE("/tasks/:id", middleware.UserAuthMiddleware(), controllers.DeleteTaskHandler)
	router.POST("/tasks", middleware.UserAuthMiddleware(), controllers.AddTaskHandler)
	router.POST("/register", controllers.RegisterHandler)
	router.POST("/login", controllers.LoginHandler)
	router.Run("localhost:8080")
}
