package repository

import (
	"context"
	"database/sql"
	"github.com/joaquimsnjunior/todo-api/config"
	"github.com/joaquimsnjunior/todo-api/schemes"
)

var db *sql.DB

func init() {
	db = config.GetDB()
}

// Task singular → representa uma task
func GetTasksByID(ctx context.Context, id int) (*schemes.Tasks, error) {
	var task schemes.Tasks
	row := db.QueryRowContext(ctx, "SELECT id, title, description, completed, created_at FROM tasks WHERE id=$1", id)
	err := row.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &task, nil
}

func CreateTasks(ctx context.Context, task *schemes.Tasks) error {
	_, err := db.ExecContext(ctx, "INSERT INTO tasks (title, description, completed) VALUES ($1, $2, $3)",
		task.Title, task.Description, task.Completed)
	return err
}

