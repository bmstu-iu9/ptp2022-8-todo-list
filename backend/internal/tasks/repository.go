package tasks

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/errors"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
)

// Repository encapsulates the logic to access tasks from the data source.
type Repository interface {
	// Get returns all tasks for user with specified id
	Get(user_id int64) ([]entity.Task, error)
	// GetById returns single task with specified id
	GetById(task_id int64) (entity.Task, error)
	// Create saves new task
	Create(task_data *entity.Task) error
	// Update modifies task
	Update(task_data *entity.Task) error
	// Delete removes task with specified id
	Delete(task_id int64) error
}

// repository persists tasks in database.
type repository struct {
	db     *sql.DB
	logger log.Logger
}

// NewRepository creates a new tasks repository.
func NewRepository(db *sql.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads all tasks with specified user id from database.
func (r repository) Get(user_id int64) ([]entity.Task, error) {
	q := "SELECT id, user_id, name, description, created_on, due_date, schtirlich_humorescue, labels, cur_status FROM tasks WHERE user_id = $1;"

	rows, err := r.db.Query(q, user_id)
	defer rows.Close()

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, fmt.Errorf("%w: %v", errors.ErrNotFound, err)
		default:
			return nil, fmt.Errorf("%w: %v", errors.ErrDb, err)
		}
	}

	tasks := make([]entity.Task, 0)
	task := entity.Task{}

	var task_created_on string
	var task_due_date string

	for rows.Next() {
		err = rows.Scan(
			&task.Id,
			&task.UserId,
			&task.Name,
			&task.Description,
			&task_created_on,
			&task_due_date,
			&task.SchtirlichHumorescue,
			&task.Labels,
			&task.Status,
		)

		if err != nil {
			return nil, fmt.Errorf("%w: %v", errors.ErrDb, err)
		}

		task.CreatedOn = entity.Date(task_created_on)
		task.DueDate = entity.Date(task_due_date)

		tasks = append(tasks, task)
	}

	return tasks, nil
}

// Get reads the task with specified id from database.
func (r repository) GetById(task_id int64) (entity.Task, error) {
	q := "SELECT id, user_id, name, description, created_on, due_date, schtirlich_humorescue, labels, cur_status FROM tasks WHERE id = $1;"

	task := entity.Task{}
	var task_created_on string
	var task_due_date string

	err := r.db.QueryRow(q, task_id).Scan(
		&task.Id,
		&task.UserId,
		&task.Name,
		&task.Description,
		&task_created_on,
		&task_due_date,
		&task.SchtirlichHumorescue,
		&task.Labels,
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

	task.CreatedOn = entity.Date(task_created_on)
	task.DueDate = entity.Date(task_due_date)

	return task, err
}

// Create saves a new task in repository and sets an id field of task_data argument
// to the id of saved task.
func (r repository) Create(task_data *entity.Task) error {
	q := `INSERT INTO tasks (user_id, name, description, created_on, due_date, schtirlich_humorescue, labels, cur_status) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;`

	err := r.db.QueryRow(
		q,
		task_data.UserId,
		task_data.Name,
		task_data.Description,
		task_data.CreatedOn,
		task_data.DueDate,
		task_data.SchtirlichHumorescue,
		task_data.Labels,
		task_data.Status).Scan(&task_data.Id)

	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrDb, err)
	}

	return nil
}

// Update saves changes to a task from database.
func (r repository) Update(task_data *entity.Task) error {
	q := "UPDATE tasks SET (name, description, due_date, schtirlich_humorescue, labels, cur_status) = ($1, $2, $3, $4, $5, $6) WHERE id = $7;"
	_, err := r.db.Exec(q, task_data.Name, task_data.Description, task_data.DueDate, task_data.SchtirlichHumorescue, task_data.Labels, task_data.Status, task_data.Id)

	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrDb, err)
	}

	return nil
}

// Delete removes a task with specified id from database.
func (r repository) Delete(task_id int64) error {
	q := "DELETE FROM tasks WHERE id = $1;"
	_, err := r.db.Exec(q, task_id)
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrDb, err)
	}
	return nil
}
