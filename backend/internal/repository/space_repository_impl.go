package repository

import (
	"dev-oleksandrv/taskera-app/internal/database"
	"dev-oleksandrv/taskera-app/internal/model/domain"
	"fmt"
	"github.com/google/uuid"
)

type SpaceRepositoryImpl struct{}

func NewSpaceRepository() SpaceRepository {
	return &SpaceRepositoryImpl{}
}

func (s *SpaceRepositoryImpl) GetSpaceRoleByUserID(spaceID, userID uuid.UUID) *domain.Role {
	var spaceUser domain.SpaceUser
	if err := database.DB.Where("space_id = ? AND user_id = ?", spaceID, userID).First(&spaceUser).Error; err != nil {
		fmt.Println(err, "ERROR")
		return nil
	}
	return &spaceUser.Role
}

func (s *SpaceRepositoryImpl) GetAllByUser(userID uuid.UUID) ([]domain.SpaceWithRole, error) {
	var spacesWithRole []domain.SpaceWithRole
	err := database.DB.Table("space_users su").
		Select("s.*", "su.role").
		Joins("JOIN spaces s ON su.space_id = s.id").
		Where("su.user_id = ?", userID).
		Scan(&spacesWithRole).Error
	if err != nil {
		return nil, err
	}
	return spacesWithRole, nil
}

func (s *SpaceRepositoryImpl) Create(space *domain.Space, userID uuid.UUID) error {
	tx := database.DB.Begin()

	if err := tx.Create(space).Error; err != nil {
		tx.Rollback()
		return err
	}

	spaceUserConn := domain.SpaceUser{
		UserID:  userID,
		SpaceID: space.ID,
		Role:    domain.Owner,
	}

	if err := tx.Create(&spaceUserConn).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (s *SpaceRepositoryImpl) Update(space *domain.Space) error {
	if err := database.DB.Updates(&space).Error; err != nil {
		return err
	}
	return nil
}

func (s *SpaceRepositoryImpl) Delete(spaceID uuid.UUID) error {
	tx := database.DB.Begin()

	var space domain.Space
	if err := tx.First(&space, spaceID).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&space).Association("Users").Clear(); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&space).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (s *SpaceRepositoryImpl) CreateSpaceUserRelation(spaceUser *domain.SpaceUser) error {
	if err := database.DB.Create(&spaceUser).Error; err != nil {
		return err
	}
	return nil
}
