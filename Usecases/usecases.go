package usecases

import "github.com/Abzaek/clean-arch/domain"

type Usecases interface {
	UpdateTask(task *domain.Task) error
	GetAllTasks() ([]*domain.Task, error)
	GetTaskById(taskId string) (*domain.Task, error)
	DeleteTask(taskId string) error
	SaveTask(task *domain.Task) error

	FindUser(userId string) (*domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUser(userId string) error
	SaveUser(user *domain.User) error
}
