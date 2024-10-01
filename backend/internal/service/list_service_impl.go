package service

import (
	"dev-oleksandrv/taskera-app/internal/model/domain"
	"dev-oleksandrv/taskera-app/internal/repository"
	"github.com/google/uuid"
)

type ListServiceImpl struct {
	listRepository repository.ListRepository
}

func NewListService(listRepository repository.ListRepository) ListService {
	return &ListServiceImpl{listRepository}
}

func (s *ListServiceImpl) GetListByID(listID uuid.UUID) (*domain.List, error) {
	return s.listRepository.GetListByID(listID)
}

func (s *ListServiceImpl) GetAllBySpaceID(spaceID uuid.UUID) ([]domain.List, error) {
	return s.listRepository.GetAllBySpaceID(spaceID)
}

func (s *ListServiceImpl) Create(list *domain.List) error {
	return s.listRepository.Create(list)
}

func (s *ListServiceImpl) Update(list *domain.List) error {
	return s.listRepository.Update(list)
}

func (s *ListServiceImpl) Delete(listID uuid.UUID) error {
	return s.listRepository.Delete(listID)
}

func (s *ListServiceImpl) Archive(list *domain.List) error {
	return s.listRepository.Archive(list)
}
