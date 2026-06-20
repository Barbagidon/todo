package repository

import (
	"encoding/json"
	"os"
	"todoapp/internal/core/domain"
)

type TaskRepository struct {
	tasks    []domain.Task
	filePath string
}

func NewTaskRepository(path string) *TaskRepository {

	repo := &TaskRepository{
		tasks:    make([]domain.Task, 0, 10),
		filePath: path,
	}

	data, err := os.ReadFile(path)

	if err != nil {
		return repo
	}

	json.Unmarshal(data, &repo.tasks)

	return repo
}

func (r *TaskRepository) Add(t domain.Task) error {
	r.tasks = append(r.tasks, t)
	return r.saveToFile()
}

func (r *TaskRepository) GetAll() []domain.Task {
	return r.tasks
}

func (r *TaskRepository) Complete(index int) error {
	if index >= 0 && index < len(r.tasks) {
		r.tasks[index].IsDone = true
		return r.saveToFile()
	}
	return nil
}

func (r *TaskRepository) saveToFile() error {
	// 1. Превращаем слайс r.tasks в байты JSON
	data, err := json.Marshal(r.tasks)
	if err != nil {
		return err
	}

	// 2. Записываем эти байты в файл
	// Функция os.WriteFile принимает: (путь_к_файлу, данные_в_байтах, права_доступа)
	err = os.WriteFile(r.filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
