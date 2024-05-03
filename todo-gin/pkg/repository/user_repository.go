package repository

import (
	"gorm.io/gorm"
	"todo-gin/pkg/domains/models"
)

type UserRepository interface {
	FindMany() ([]models.User, error)
	FindByID(id int) (models.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (u *UserRepositoryImpl) FindMany() ([]models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepositoryImpl) FindByID(id int) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}
