package repository

import (
	"dev-oleksandrv/taskera-app/internal/database"
	"dev-oleksandrv/taskera-app/internal/model/domain"
	"github.com/google/uuid"
)

type ListRepositoryImpl struct{}

func NewListRepository() ListRepository {
	return &ListRepositoryImpl{}
}

func (r *ListRepositoryImpl) GetListByID(listID uuid.UUID) (*domain.List, error) {
	var list domain.List
	if err := database.DB.Model(&domain.List{}).Where("id = ?", listID).First(&list).Error; err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *ListRepositoryImpl) GetAllBySpaceID(spaceID uuid.UUID) ([]domain.List, error) {
	var lists []domain.List
	if err := database.DB.Model(&domain.List{}).Where("space_id = ?", spaceID).Find(&lists).Error; err != nil {
		return nil, err
	}
	return lists, nil
}

func (r *ListRepositoryImpl) Create(list *domain.List) error {
	if err := database.DB.Create(&list).Error; err != nil {
		return err
	}
	return nil
}

func (r *ListRepositoryImpl) Update(list *domain.List) error {
	if err := database.DB.Updates(&list).Error; err != nil {
		return err
	}
	return nil
}

func (r *ListRepositoryImpl) Delete(listID uuid.UUID) error {
	var list domain.List
	if err := database.DB.First(&list, listID).Error; err != nil {
		return err
	}

	if err := database.DB.Delete(&list).Error; err != nil {
		return err
	}

	return nil
}

func (r *ListRepositoryImpl) Archive(list *domain.List) error {
	// TODO
	return nil
}
