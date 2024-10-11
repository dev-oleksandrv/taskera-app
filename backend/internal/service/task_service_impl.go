package service

import (
	"dev-oleksandrv/taskera-app/internal/model/domain"
	"dev-oleksandrv/taskera-app/internal/repository"
	"github.com/google/uuid"
)

type TaskServiceImpl struct {
	taskRepository repository.TaskRepository
}

func (t TaskServiceImpl) GetAllByListID(listID uuid.UUID) ([]domain.Task, error) {
	return t.taskRepository.GetTasksByListID(listID)
}

func (t TaskServiceImpl) Create(task *domain.Task) error {
	return t.taskRepository.Create(task)
}

func (t TaskServiceImpl) Update(task *domain.Task) error {
	//TODO implement me
	panic("implement me")
}

func (t TaskServiceImpl) Delete(taskID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (t TaskServiceImpl) Toggle(taskID uuid.UUID) (*domain.Task, error) {
	//TODO implement me
	panic("implement me")
}

func NewTaskService(taskRepository repository.TaskRepository) TaskService {
	return &TaskServiceImpl{taskRepository}
}
