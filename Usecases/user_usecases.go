package usecases

import (
	"github.com/Abzaek/clean-arch/domain"
)

type UserService interface {
	Save(user *domain.User) error
	Delete(userId string) error
	Update(user *domain.User) error
	Find(userId string) (*domain.User, error)
}

type UserUsecase struct {
	repo UserService
}

func NewUserUseCase(repo UserService) *UserUsecase {
	return &UserUsecase{
		repo: repo,
	}
}

func (uc *UserUsecase) SaveUser(user *domain.User) error {
	err := uc.repo.Save(user)
	return err
}

func (uc *UserUsecase) DeleteUser(userId string) error {
	err := uc.repo.Delete(userId)

	return err
}

func (uc *UserUsecase) UpdateUser(user *domain.User) error {
	err := uc.repo.Update(user)

	return err
}

func (uc *UserUsecase) FindUser(userId string) (*domain.User, error) {
	user, err := uc.repo.Find(userId)

	return user, err
}
