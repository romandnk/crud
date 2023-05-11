package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/romandnk/crud/internal/entities"
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

func (t *TaskPostgres) Create(ctx context.Context, task entities.Task) (int, error) {
	var id int

	query := fmt.Sprintf(`INSERT INTO %s (creation_time, updating_time, message) VALUES ($1, $2, $3) RETURNING id`, tasksTable)

	row := t.db.QueryRow(ctx, query,
		time.Now(),
		time.Now(),
		task.Message,
	)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (t *TaskPostgres) GetAll(ctx context.Context) ([]entities.Task, error) {
	var tasks []entities.Task

	query := fmt.Sprintf(`SELECT id, creation_time, updating_time, message FROM %s`, tasksTable)

	rows, err := t.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			currentTask    entities.Task
			creationTimeTS pgtype.Timestamp
			updatingTimeTS pgtype.Timestamp
		)

		err = rows.Scan(&currentTask.Id, &creationTimeTS, &updatingTimeTS, &currentTask.Message)
		if err != nil {
			return nil, err
		}
		currentTask.CreationTime = creationTimeTS.Time
		currentTask.UpdatingTime = updatingTimeTS.Time

		tasks = append(tasks, currentTask)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t *TaskPostgres) GetById(ctx context.Context, id int) (entities.Task, error) {
	var (
		selectedTask   entities.Task
		creationTimeTS pgtype.Timestamp
		updatingTimeTS pgtype.Timestamp
	)

	query := fmt.Sprintf(`SELECT id, creation_time, updating_time, message FROM %s WHERE id = $1`, tasksTable)

	row := t.db.QueryRow(ctx, query, id)
	if err := row.Scan(&selectedTask.Id, &creationTimeTS, &updatingTimeTS, &selectedTask.Message); err != nil {
		return selectedTask, err
	}
	if id == 0 {
		return selectedTask, fmt.Errorf("no task with such an id")
	}

	selectedTask.CreationTime = creationTimeTS.Time
	selectedTask.UpdatingTime = updatingTimeTS.Time

	return selectedTask, nil
}

func (t *TaskPostgres) Update(ctx context.Context, id int, task entities.Task) (entities.Task, error) {
	actualTask, err := t.GetById(ctx, id)
	if err != nil {
		return actualTask, err
	}

	query := fmt.Sprintf(`UPDATE %s SET updating_time = $1, message = $2 WHERE id = $3`, tasksTable)

	_, err = t.db.Exec(ctx, query,
		time.Now(),
		task.Message,
		id,
	)
	if err != nil {
		return actualTask, err
	}

	return t.GetById(ctx, id)
}

func (t *TaskPostgres) Delete(ctx context.Context, id int) error {
	_, err := t.GetById(ctx, id)
	if err != nil {
		return err
	}

	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, tasksTable)
	_, err = t.db.Exec(ctx, query, id)
	return err
}
