package repository

import (
	"dev-oleksandrv/taskera-app/internal/database"
	"dev-oleksandrv/taskera-app/internal/model/domain"
	"github.com/google/uuid"
)

type TaskRepositoryImpl struct{}

func NewTaskRepository() TaskRepository {
	return &TaskRepositoryImpl{}
}

func (r *TaskRepositoryImpl) GetTasksByListID(listID uuid.UUID) ([]domain.Task, error) {
	var tasks []domain.Task
	if err := database.DB.Model(&domain.Task{}).Where("list_id = ?", listID).Order("tasks.order ASC").Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepositoryImpl) Create(task *domain.Task) error {
	tx := database.DB.Begin()

	var maxOrder int
	if err := tx.Model(&domain.Task{}).Select("MAX(tasks.order)").Scan(&maxOrder).Error; err != nil {
		tx.Rollback()
		return err
	}

	task.Order = maxOrder + 1

	if err := tx.Create(&task).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (r *TaskRepositoryImpl) Update(task *domain.Task) error {
	tx := database.DB.Begin()
	if err := tx.Updates(&task).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *TaskRepositoryImpl) Delete(taskID uuid.UUID) error {
	tx := database.DB.Begin()
	var task domain.Task
	if err := tx.First(&task, taskID).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&task).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *TaskRepositoryImpl) GetTaskByID(taskID uuid.UUID) (*domain.Task, error) {
	var task domain.Task
	if err := database.DB.First(&task, taskID).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepositoryImpl) Reorder(ids []uuid.UUID) error {
	// TODO make reodering of tasks in list
	return nil
}
