package repository

import (
	"context"
	"todoapp/internal/core/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TaskRepository struct {
	db *pgxpool.Pool
}

func NewTaskRepository(db *pgxpool.Pool) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (r *TaskRepository) Add(t domain.Task) error {
	query := `INSERT INTO tasks (title, completed) VALUES ($1, $2);`

	_, err := r.db.Exec(context.Background(), query, t.Title, t.IsDone)

	return err
}

func (r *TaskRepository) GetAll() ([]domain.Task, error) {
	query := `SELECT id, title, completed FROM tasks ORDER BY id;`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []domain.Task

	for rows.Next() {
		var task domain.Task

		err := rows.Scan(&task.ID, &task.Title, &task.IsDone)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *TaskRepository) Complete(id int64) error {
	query := `UPDATE tasks SET completed = true WHERE id = $1;`

	_, err := r.db.Exec(context.Background(), query, id)
	return err
}

func (r *TaskRepository) Delete(id int64) error {
	query := `DELETE FROM tasks WHERE id = $1;`

	_, err := r.db.Exec(context.Background(), query, id)
	return err
}
