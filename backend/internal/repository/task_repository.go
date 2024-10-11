package repository

import (
	"dev-oleksandrv/taskera-app/internal/model/domain"
	"github.com/google/uuid"
)

type TaskRepository interface {
	GetTasksByListID(listID uuid.UUID) ([]domain.Task, error)
	GetTaskByID(taskID uuid.UUID) (*domain.Task, error)
	Create(task *domain.Task) error
	Update(task *domain.Task) error
	Delete(taskID uuid.UUID) error
	Reorder(ids []uuid.UUID) error
}
