package domain

type TaskUseCase interface {
	GetTask(sID string) (Task, error)
	GetTasks(id string) []Task
	UpdateTask(sID string, task Task) error
	AddTask(newTask Task) (string, error)
	DeleteTask(sID string) error
}
