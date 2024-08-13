package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskRepository interface {
	AllTasks(id string) []Task
	AddTask(user *Task) (primitive.ObjectID, error)
	DeleteTask(id primitive.ObjectID) error
	GetTask(id primitive.ObjectID) (error, Task)
	UpdateTask(primitive.ObjectID, Task) error
}
