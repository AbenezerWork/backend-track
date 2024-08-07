package controllers

import (
	"fmt"
	"task_managerv2/data"
	"task_managerv2/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func GetTask(id primitive.ObjectID) (model.Task, error) {
	err, task := data.TaskManagerInstance.GetTask(id)

	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func GetTasks(id string) []model.Task {
	data := data.TaskManagerInstance.AllTasks(id)

	return data
}

func UpdateTask(id primitive.ObjectID, task model.Task) {
	data.TaskManagerInstance.UpdateTask(id, task)
}

func AddTask(newTask model.Task) (primitive.ObjectID, error) {
	id, err := data.TaskManagerInstance.AddTask(&newTask)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return id, nil
}

func DeleteTask(id primitive.ObjectID) error {
	return data.TaskManagerInstance.DeleteTask(id)
}

func addUser(user model.User) (primitive.ObjectID, error) {
	id, err := data.TaskManagerInstance.AddUser(&user)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return id, nil
}

func UserLogin(user *model.User) error {
	userByEmail, err := data.TaskManagerInstance.GetUserByEmail(user.Email)
	if err != nil {
		fmt.Println("userByEmail")
		return err
	}
	fmt.Println(bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost))
	fmt.Println([]byte(userByEmail.Password))

	err = bcrypt.CompareHashAndPassword([]byte(userByEmail.Password), []byte(user.Password))
	fmt.Println(err, "hello")
	if err != nil {
		fmt.Println("CompareHashAndPassword")
		return err
	}
	user.ID = userByEmail.ID
	return nil
}
