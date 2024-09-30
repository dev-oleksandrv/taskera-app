package repository

import (
	"dev-oleksandrv/taskera-app/internal/model/domain"
	"github.com/google/uuid"
)

type SpaceRepository interface {
	GetAllByUser(userID uuid.UUID) ([]domain.SpaceWithRole, error)
	GetSpaceRoleByUserID(spaceID, userID uuid.UUID) *domain.Role
	Create(space *domain.Space, userID uuid.UUID) error
	Update(space *domain.Space) error
	Delete(spaceID uuid.UUID) error
	CreateSpaceUserRelation(spaceUser *domain.SpaceUser) error
}
