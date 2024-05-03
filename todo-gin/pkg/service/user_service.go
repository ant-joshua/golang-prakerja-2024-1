package service

import (
	"todo-gin/pkg/domains/models"
	"todo-gin/pkg/repository"
)

type UserService struct {
	Repository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return UserService{
		Repository: repo,
	}
}

func (u *UserService) GetAllUser() ([]models.User, error) {
	users, err := u.Repository.FindMany()

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserService) GetUserByID(id int) (models.User, error) {
	user, err := u.Repository.FindByID(id)

	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
