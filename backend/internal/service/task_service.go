package service

import (
	"dev-oleksandrv/taskera-app/internal/model/domain"
	"github.com/google/uuid"
)

type TaskService interface {
	GetAllByListID(listID uuid.UUID) ([]domain.Task, error)
	Create(task *domain.Task) error
	Update(task *domain.Task) error
	Delete(taskID uuid.UUID) error
	Toggle(taskID uuid.UUID) (*domain.Task, error)
}
