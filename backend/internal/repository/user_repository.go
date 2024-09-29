package repository

import "dev-oleksandrv/taskera-app/internal/model/domain"

type UserRepository interface {
	Register(user *domain.User) error
	Login(user *domain.User) error
}
