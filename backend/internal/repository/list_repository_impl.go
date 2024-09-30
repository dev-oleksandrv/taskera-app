package repository

import (
	"dev-oleksandrv/taskera-app/internal/model/domain"
	"github.com/google/uuid"
)

type ListRepositoryImpl struct{}

func NewListRepository() ListRepository {
	return &ListRepositoryImpl{}
}

func (r *ListRepositoryImpl) GetAllBySpaceID(spaceID uuid.UUID) ([]domain.List, error) {
	return nil, nil
}

func (r *ListRepositoryImpl) Create(list *domain.List, userID uuid.UUID) error {
	return nil
}

func (r *ListRepositoryImpl) Update(list *domain.List) error {
	return nil
}

func (r *ListRepositoryImpl) Delete(listID uuid.UUID) error {
	return nil
}

func (r *ListRepositoryImpl) Archive(list *domain.List) error {
	return nil
}
