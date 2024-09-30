package service

import "dev-oleksandrv/taskera-app/internal/model/domain"

type UserService interface {
	Register(user *domain.User) error
	Login(user *domain.User) error
	GetUserByEmail(email string) (*domain.User, error)
}
