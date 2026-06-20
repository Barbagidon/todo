package service

import (
	"errors"
	"todoapp/internal/core/domain"
	"todoapp/internal/features/tasks/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(title string) error {
	if title == "" {
		return errors.New("название задачи не может быть пустым")
	}

	task := domain.Task{
		Title:  title,
		IsDone: false,
	}

	return s.repo.Add(task)

}

func (s *TaskService) GetTasks() []domain.Task {
	return s.repo.GetAll()
}

func (s *TaskService) CompleteTask(index int) error {
	tasks := s.repo.GetAll()
	if index < 0 || index >= len(tasks) {
		return errors.New("неверный номер задачи")
	}

	return s.repo.Complete(index)

}
