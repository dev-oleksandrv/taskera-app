package service

import (
	"dev-oleksandrv/taskera-app/internal/model/domain"
	"dev-oleksandrv/taskera-app/internal/repository"
	"github.com/google/uuid"
)

type SpaceServiceImpl struct {
	spaceRepository repository.SpaceRepository
}

func NewSpaceService(spaceRepository repository.SpaceRepository) SpaceService {
	return &SpaceServiceImpl{spaceRepository}
}

func (s *SpaceServiceImpl) GetAllByUser(userID uuid.UUID) ([]domain.SpaceWithRole, error) {
	return s.spaceRepository.GetAllByUser(userID)
}

func (s *SpaceServiceImpl) GetSpaceRoleByUserID(spaceID, userID uuid.UUID) *domain.Role {
	return s.spaceRepository.GetSpaceRoleByUserID(spaceID, userID)
}

func (s *SpaceServiceImpl) Create(space *domain.Space, userID uuid.UUID) error {
	if err := s.spaceRepository.Create(space, userID); err != nil {
		return err
	}
	return nil
}

func (s *SpaceServiceImpl) Update(space *domain.Space) error {
	if err := s.spaceRepository.Update(space); err != nil {
		return err
	}
	return nil
}

func (s *SpaceServiceImpl) Delete(spaceID uuid.UUID) error {
	if err := s.spaceRepository.Delete(spaceID); err != nil {
		return err
	}
	return nil
}

func (s *SpaceServiceImpl) CreateSpaceUserRelation(spaceUser *domain.SpaceUser) error {
	if err := s.spaceRepository.CreateSpaceUserRelation(spaceUser); err != nil {
		return err
	}
	return nil
}
