package router

import (
	"task_managerv2/delivery/controllers"
	"task_managerv2/infrastructure"

	"github.com/gin-gonic/gin"
)

type Router struct {
	TaskController *controllers.TaskController
	UserController *controllers.UserController
}

func NewRouter(taskController *controllers.TaskController, userController *controllers.UserController) Router {
	return Router{TaskController: taskController, UserController: userController}
}

func (r *Router) Route() {
	router := gin.Default()

	router.GET("/tasks", infrastructure.UserAuthMiddleware(), r.TaskController.GetTasksHandler)
	router.GET("/tasks/:id", infrastructure.UserAuthMiddleware(), r.TaskController.GetTaskHandler)
	router.PUT("/tasks/:id", infrastructure.UserAuthMiddleware(), r.TaskController.UpdateTaskHandler)
	router.DELETE("/tasks/:id", infrastructure.UserAuthMiddleware(), r.TaskController.DeleteTaskHandler)
	router.POST("/tasks", infrastructure.UserAuthMiddleware(), r.TaskController.AddTaskHandler)

	router.POST("/register", infrastructure.UserAuthMiddleware(), infrastructure.AdminAuthMiddleware(), r.UserController.RegisterHandler)
	router.DELETE("/removeuser/:id", infrastructure.UserAuthMiddleware(), infrastructure.AdminAuthMiddleware(), r.UserController.RemoveUser)

	router.POST("/login", r.UserController.LoginHandler)
	router.Run("localhost:8080")
}
