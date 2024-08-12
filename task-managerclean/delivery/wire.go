// +build wireinject

package main

import (
	"task_managerv2/bootstrap"
	"task_managerv2/delivery/controllers"
	"task_managerv2/delivery/router"
	"task_managerv2/repository"
	"task_managerv2/usecase"

	"github.com/google/wire"
)

func InitializeRouter() router.Router {
	wire.Build(bootstrap.InitMongoDB, repository.InitTaskRepository, repository.InitUserRepository, usecase.InitTaskUsecase, usecase.InitUserUsecase, controllers.InitTaskController, controllers.InitUserController, router.NewRouter)
	return router.Router{}
}
