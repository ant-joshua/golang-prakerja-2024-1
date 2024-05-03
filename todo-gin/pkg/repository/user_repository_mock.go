package repository

import (
	"github.com/stretchr/testify/mock"
	"todo-gin/pkg/domains/models"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (r *UserRepositoryMock) FindByID(id int) (models.User, error) {
	args := r.Mock.Called(id)

	return args.Get(0).(models.User), args.Error(1)
}

func (r *UserRepositoryMock) FindMany() ([]models.User, error) {
	args := r.Mock.Called()

	return args.Get(0).([]models.User), args.Error(1)
}
