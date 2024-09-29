package repository

import (
	"dev-oleksandrv/taskera-app/internal/database"
	"dev-oleksandrv/taskera-app/internal/model/domain"
	"dev-oleksandrv/taskera-app/internal/utils"
	"errors"
)

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) Register(user *domain.User) error {
	if err := database.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepositoryImpl) Login(user *domain.User) error {
	password := user.Password

	if err := database.DB.Where("email = ?", user.Email).First(&user).Error; err != nil {
		return errors.New("invalid user email")
	}

	if valid := utils.CheckPasswordHash(password, user.Password); !valid {
		return errors.New("invalid user password")
	}

	return nil
}
