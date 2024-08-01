package controllers

import (
	"task_managerv2/data"
	"task_managerv2/model"
)

func GetTask(id int) (model.Task, error) {
	err, task := data.TaskManagerInstance.GetTask(id)

	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func GetTasks() []model.Task {
	data := data.TaskManagerInstance.AllTasks()

	ret := []model.Task{}

	for _, value := range data {
		ret = append(ret, value)
	}

	return ret
}

func UpdateTask(id int, task model.Task) {
	data.TaskManagerInstance.UpdateTask(id, task)
}

func AddTask(newTask model.Task) int {
	id := data.TaskManagerInstance.AddTask(&newTask)
	return id
}

func DeleteTask(id int) error {
	return data.TaskManagerInstance.DeleteTask(id)
}
