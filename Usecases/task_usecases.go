package usecases

import (
	domain "github.com/Abzaek/clean-arch/domain"
)

type TaskUseCase struct {
	repo TaskService
}

type TaskService interface {
	Update(task *domain.Task) error
	Save(task *domain.Task) error
	Delete(taskId string) error
	GetById(taskId string) (*domain.Task, error)
	GetAll() ([]*domain.Task, error)
}

func NewTaskUseCase(repo TaskService) *TaskUseCase {
	return &TaskUseCase{
		repo: repo,
	}
}

func (uc *TaskUseCase) GetAllTasks() ([]*domain.Task, error) {
	tasks, err := uc.repo.GetAll()

	return tasks, err
}

func (uc *TaskUseCase) GetTaskById(taskId string) (*domain.Task, error) {
	task, err := uc.repo.GetById(taskId)

	return task, err
}

func (uc *TaskUseCase) DeleteTask(taskId string) error {
	err := uc.repo.Delete(taskId)

	return err
}

func (uc *TaskUseCase) SaveTask(task *domain.Task) error {
	err := uc.repo.Save(task)

	return err
}

func (uc *TaskUseCase) UpdateTask(task *domain.Task) error {
	err := uc.repo.Update(task)

	return err
}
