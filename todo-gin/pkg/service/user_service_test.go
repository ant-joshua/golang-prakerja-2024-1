package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"todo-gin/pkg/domains/models"
	"todo-gin/pkg/repository"
)

var userRepository = &repository.UserRepositoryMock{
	Mock: mock.Mock{},
}
var userService = UserService{
	Repository: userRepository,
}

func TestGetAllUser(t *testing.T) {

	mockData := []models.User{
		{
			ID:       1,
			Username: "username",
			Password: "password",
		},
	}

	userRepository.Mock.On("FindMany").Return([]models.User{
		{
			ID:       1,
			Username: "username",
			Password: "password",
		},
	}, nil)

	users, err := userService.GetAllUser()
	assert.Nil(t, err, "GetAllUser() = %v; want nil", err)
	assert.NotEmpty(t, users, "GetAllUser() = %v; want not empty", users)
	assert.Equal(t, users, mockData, "GetAllUser() = %v; want %v", users, mockData)
}

func TestGetUserByID(t *testing.T) {

	mockData := models.User{
		ID:       1,
		Username: "username",
		Password: "password",
	}

	userRepository.Mock.On("FindByID", 1).Return(models.User{
		ID:       1,
		Username: "username",
		Password: "password",
	}, nil)

	user, err := userService.GetUserByID(1)

	assert.Nil(t, err, "GetUserByID(1) = %v; want nil", err)
	assert.NotEmpty(t, user, "GetUserByID(1) = %v; want not empty", user)
	assert.Equal(t, user, mockData, "GetUserByID(1) = %v; want %v", user, mockData)
}
