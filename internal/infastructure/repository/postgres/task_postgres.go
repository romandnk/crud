package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/romandnk/crud/internal/entities/task"
	"time"
)

type TaskPostgres struct {
	db *pgx.Conn
}

func NewTaskPostgres(db *pgx.Conn) *TaskPostgres {
	return &TaskPostgres{
		db: db,
	}
}

func (t *TaskPostgres) Create(ctx context.Context, task task.Task) (int, error) {
	var id int

	query := `INSERT INTO task (creation_time, message) VALUES ($1, $2) RETURNING id`

	row := t.db.QueryRow(ctx, query,
		time.Now().Format("01/02/2000 12:00:00"),
		time.Now().Format("01/02/2000 12:00:00"),
		task.Message,
	)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (t *TaskPostgres) GetAll(ctx context.Context) ([]task.Task, error) {
	var tasks []task.Task

	query := `SELECT id, creation_time, updating_time message FROM task`

	rows, err := t.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var currentTask task.Task

		err = rows.Scan(&currentTask.Id, &currentTask.CreationTime, &currentTask.Message)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, currentTask)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t *TaskPostgres) GetById(ctx context.Context, id int) (task.Task, error) {
	var selectedTask task.Task

	query := `SELECT id, creation_time, message FROM task WHERE id = $1`

	row := t.db.QueryRow(ctx, query, id)
	if err := row.Scan(&selectedTask.Id, &selectedTask.CreationTime, &selectedTask.Message); err != nil {
		return selectedTask, nil
	}

	return selectedTask, nil
}

func (t *TaskPostgres) Update(ctx context.Context, id int, task task.Task) (task.Task, error) {
	actualTask, err := t.GetById(ctx, id)
	if err != nil {
		return actualTask, err
	}

	query := `UPDATE task SET updating_time = $1, message = $2 WHERE id = $3`

	_, err = t.db.Exec(ctx, query,
		time.Now().Format("01/02/2000 12:00:00"),
		task.Message,
		id,
	)
	if err != nil {
		return actualTask, err
	}

	return t.GetById(ctx, id)
}

func (t *TaskPostgres) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM task WHERE id = $1`
	_, err := t.db.Exec(ctx, query, id)
	return err
}
