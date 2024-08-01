package data

import (
	"errors"
	"log"
	"task_managerv2/model"
	"time"
)

type TaskManager struct {
	tasks  map[int]model.Task
	nextID int
}

// Initialize the TaskManagerInstance
var TaskManagerInstance = TaskManager{}

func InitTaskManager() *TaskManager {
	TaskManagerInstance = TaskManager{}
	TaskManagerInstance.tasks = make(map[int]model.Task)
	return &TaskManagerInstance
}

func (tm *TaskManager) AddTask(task *model.Task) int {
	task.ID = tm.nextID
	tm.tasks[tm.nextID] = *task
	tm.nextID++
	return tm.nextID - 1
}

func (tm *TaskManager) DeleteTask(id int) error {
	_, exists := tm.tasks[id]
	if !exists {
		return errors.New("The task you're trying to delete doesn't exist.")
	}
	delete(tm.tasks, id)
	return nil
}

func (tm *TaskManager) GetTask(id int) (error, model.Task) {
	task, exists := tm.tasks[id]
	if !exists {
		return errors.New("The task you're trying to delete doesn't exist."), model.Task{}
	}
	return nil, task
}

func (tm *TaskManager) UpdateTask(id int, task model.Task) error {
	oldTask, exists := tm.tasks[id]
	if !exists {
		return errors.New("The task you're trying to delete doesn't exist.")
	}
	if task.Title != "" {
		oldTask.Title = task.Title
	}
	if task.Description != "" {
		oldTask.Description = task.Description
	}
	if task.Status != "" {
		oldTask.Status = task.Description
	}
	if task.DueDate != (time.Time{}) {
		oldTask.DueDate = task.DueDate

	}
	tm.tasks[task.ID] = oldTask
	return nil
}

func (tm *TaskManager) AllTasks() map[int]model.Task {
	log.Println(tm.tasks)
	return tm.tasks
}
