package controllers

import (
	"task_managerv2/data"
	"task_managerv2/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTask(id primitive.ObjectID) (model.Task, error) {
	err, task := data.TaskManagerInstance.GetTask(id)

	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func GetTasks() []model.Task {
	data := data.TaskManagerInstance.AllTasks()

	return data
}

func UpdateTask(id primitive.ObjectID, task model.Task) {
	data.TaskManagerInstance.UpdateTask(id, task)
}

func AddTask(newTask model.Task) primitive.ObjectID {
	id := data.TaskManagerInstance.AddTask(&newTask)
	return id
}

func DeleteTask(id primitive.ObjectID) error {
	return data.TaskManagerInstance.DeleteTask(id)
}
