package repository

import (
	"dev-oleksandrv/taskera-app/internal/model/domain"
	"github.com/google/uuid"
)

type ListRepository interface {
	GetAllBySpaceID(spaceID uuid.UUID) ([]domain.List, error)
	Create(list *domain.List, userID uuid.UUID) error
	Update(list *domain.List) error
	Delete(listID uuid.UUID) error
	Archive(list *domain.List) error
}
