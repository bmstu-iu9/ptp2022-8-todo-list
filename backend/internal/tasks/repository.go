package tasks

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/errors"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
)

type Repository interface {
	Get(user_id int64) ([]entity.Task, error)
	GetById(task_id int64) (entity.Task, error)
	// GetLabels(task_id int64) ([]entity.TaskLabel, error)
	Create(task_data *entity.Task) error
	Update(task_data *entity.Task) error
	Delete(task_id int64) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return repository{db}
}

func (r repository) Get(user_id int64) ([]entity.Task, error) {
	q := "SELECT id, user_id, name, description, created_on, due_date, schtirlich_humorescue, cur_status FROM tasks WHERE user_id = $1;"

	rows, err := r.db.Query(q, user_id)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, fmt.Errorf("%w: %v", errors.ErrNotFound, err)
		default:
			return nil, fmt.Errorf("%w: %v", errors.ErrDb, err)
		}
	}

	defer rows.Close()

	tasks := make([]entity.Task, 0)
	task := entity.Task{}

	for rows.Next() {
		err = rows.Scan(
			&task.Id,
			&task.UserId,
			&task.Name,
			&task.Description,
			&task.CreatedOn,
			&task.DueDate,
			&task.SchtirlichHumorescue,
			&task.Status,
		)

		if err != nil {
			return nil, fmt.Errorf("%w: %v", errors.ErrDb, err)
		}

		task.Labels, err = r.getLabels(task.Id)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r repository) GetById(task_id int64) (entity.Task, error) {
	q := "SELECT id, user_id, name, description, created_on, due_date, schtirlich_humorescue, cur_status FROM tasks WHERE id = $1;"

	task := entity.Task{}
	err := r.db.QueryRow(q, task_id).Scan(
		&task.Id,
		&task.UserId,
		&task.Name,
		&task.Description,
		&task.CreatedOn,
		&task.DueDate,
		&task.SchtirlichHumorescue,
		&task.Status,
	)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return entity.Task{}, fmt.Errorf("%w: %v", errors.ErrNotFound, err)
		default:
			return entity.Task{}, fmt.Errorf("%w: %v", errors.ErrDb, err)
		}
	}

	task.Labels, err = r.getLabels(task_id)

	return task, err
}

func (r repository) getLabels(task_id int64) ([]entity.TaskLabel, error) {
	q := "SELECT id, name, color FROM task_labels WHERE task_id = $1;"

	label := entity.TaskLabel{}
	labels := make([]entity.TaskLabel, 0)
	rows, err := r.db.Query(q, task_id)

	if err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrDb, err)
	}

	for rows.Next() {
		label.TaskId = task_id
		err = rows.Scan(&label.Id, &label.Name, &label.Color)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", errors.ErrDb, err)
		}

		labels = append(labels, label)
	}

	return labels, nil
}

func (r repository) Create(task_data *entity.Task) error {
	q := `INSERT INTO tasks (user_id, name, description, created_on, due_date, schtirlich_humorescue, cur_status) 
	VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;`

	err := r.db.QueryRow(
		q,
		task_data.UserId,
		task_data.Name,
		task_data.Description,
		task_data.CreatedOn,
		task_data.DueDate,
		task_data.SchtirlichHumorescue,
		task_data.Status).Scan(&task_data.Id)

	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrDb, err)
	}

	for _, label := range task_data.Labels {
		q := "INSERT INTO task_labels (task_id, name, color) VALUES ($1, $2, $3);"
		_, err := r.db.Exec(q, task_data.Id, label.Name, label.Color)

		if err != nil {
			return fmt.Errorf("%w: %v", errors.ErrDb, err)
		}
	}

	return nil
}

func (r repository) Update(task_data *entity.Task) error {
	q := "UPDATE tasks SET (name, description, due_date, schtirlich_humorescue, cur_status) = ($1, $2, $3, $4, $5) WHERE id = $6;"
	_, err := r.db.Exec(q, task_data.Name, task_data.Description, task_data.DueDate, task_data.SchtirlichHumorescue, task_data.Status, task_data.Id)

	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrDb, err)
	}

	logger := log.New()

	for _, label := range task_data.Labels {
		logger.Debug("Label update", label)
		if label.Id != 0 {
			q = "DELETE FROM task_labels WHERE id = $1"
			_, err = r.db.Exec(q, label.Id)
		} else {
			q = "INSERT INTO task_labels (task_id, name, color) VALUES ($1, $2, $3);"
			_, err = r.db.Exec(q, task_data.Id, label.Name, label.Color)
		}

		if err != nil {
			return fmt.Errorf("%w: %v", errors.ErrDb, err)
		}
	}

	return nil
}

func (r repository) Delete(task_id int64) error {
	q := "DELETE FROM tasks WHERE id = $1;"
	_, err := r.db.Exec(q, task_id)
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrDb, err)
	}
	return nil
}
