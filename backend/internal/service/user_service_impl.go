package service

import (
	"dev-oleksandrv/taskera-app/internal/model/domain"
	"dev-oleksandrv/taskera-app/internal/repository"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{userRepository}
}

func (s *UserServiceImpl) Register(user *domain.User) error {
	if err := s.userRepository.Register(user); err != nil {
		return err
	}

	return nil
}

func (s *UserServiceImpl) Login(user *domain.User) error {
	if err := s.userRepository.Login(user); err != nil {
		return err
	}

	return nil
}

func (s *UserServiceImpl) GetUserByEmail(email string) (*domain.User, error) {
	user, err := s.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
