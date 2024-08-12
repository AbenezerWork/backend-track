package usecase

import (
	"task_managerv2/domain"
	"task_managerv2/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUsecase struct {
	TaskRepository *repository.TaskRepository
}

func InitTaskUsecase(repository *repository.TaskRepository) *TaskUsecase {
	return &TaskUsecase{TaskRepository: repository}
}

func (tu *TaskUsecase) GetTask(sID string) (domain.Task, error) {
	id, err := primitive.ObjectIDFromHex(sID)
	err, task := tu.TaskRepository.GetTask(id)

	if err != nil {
		return domain.Task{}, err
	}
	return task, nil
}

func (tu *TaskUsecase) GetTasks(id string) []domain.Task {
	repository := tu.TaskRepository.AllTasks(id)

	return repository
}

func (tu *TaskUsecase) UpdateTask(sID string, task domain.Task) error {
	id, err := primitive.ObjectIDFromHex(sID)
	if err != nil {
		return err
	}
	return tu.TaskRepository.UpdateTask(id, task)
}

func (tu *TaskUsecase) AddTask(newTask domain.Task) (string, error) {
	id, err := tu.TaskRepository.AddTask(&newTask)
	if err != nil {
		return "", err
	}
	return id.Hex(), nil
}

func (tu *TaskUsecase) DeleteTask(sID string) error {
	id, err := primitive.ObjectIDFromHex(sID)
	if err != nil {
		return err
	}
	return tu.TaskRepository.DeleteTask(id)
}
