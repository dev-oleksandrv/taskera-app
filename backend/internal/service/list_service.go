package service

import (
	"dev-oleksandrv/taskera-app/internal/model/domain"
	"github.com/google/uuid"
)

type ListService interface {
	GetListByID(listID uuid.UUID) (*domain.List, error)
	GetAllBySpaceID(spaceID uuid.UUID) ([]domain.List, error)
	Create(list *domain.List) error
	Update(list *domain.List) error
	Delete(listID uuid.UUID) error
	Archive(list *domain.List) error
}
