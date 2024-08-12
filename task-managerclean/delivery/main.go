package main

import (
	"github.com/joho/godotenv"
	"task_managerv2/bootstrap"
	"task_managerv2/delivery/controllers"
	"task_managerv2/delivery/router"
	"task_managerv2/repository"
	"task_managerv2/usecase"
)

func main() {
	godotenv.Load()

	//get the database
	database := bootstrap.InitMongoDB()

	taskRepository := repository.InitTaskRepository(database)
	userRepository := repository.InitUserRepository(database)

	taskUsecase := usecase.InitTaskUsecase(taskRepository)

	userUsecase := usecase.InitUserUsecase(userRepository)

	taskController := controllers.InitTaskController(taskUsecase)
	userController := controllers.InitUserController(userUsecase)
	route := router.NewRouter(taskController, userController)
	route.Route()

	// router := InitializeRouter()
	// router.Route()
}
