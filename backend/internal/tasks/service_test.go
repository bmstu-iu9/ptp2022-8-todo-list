package tasks

import (
	"testing"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/errors"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/test"
)

type mockRepository struct {
	items    []entity.Task
	task_id  int64
	label_id int64
}

func (r *mockRepository) Get(user_id int64) ([]entity.Task, error) {
	res := []entity.Task{}
	for _, item := range r.items {
		if item.UserId == user_id {
			res = append(res, item)
		}
	}

	if len(res) == 0 {
		return nil, errors.ErrNotFound
	}

	return res, nil
}

func (r *mockRepository) GetById(task_id int64) (entity.Task, error) {
	for _, item := range r.items {
		if item.Id == task_id {
			return item, nil
		}
	}
	return entity.Task{}, errors.ErrNotFound
}

func (r *mockRepository) Create(task *entity.Task) error {
	task.Id = r.task_id
	r.task_id++
	r.label_id += int64(len(task.Labels))
	r.items = append(r.items, *task)
	return nil
}

func (r *mockRepository) Update(task *entity.Task) error {
	for i := 0; i < len(r.items); i++ {
		if r.items[i].Id == task.Id {
			r.items[i].Name = task.Name
			r.items[i].Description = task.Description
			r.items[i].DueDate = task.DueDate
			r.items[i].SchtirlichHumorescue = task.SchtirlichHumorescue
			r.items[i].Labels = task.Labels
			r.items[i].Status = task.Status
			return nil
		}
	}
	return errors.ErrNotFound
}

func (r *mockRepository) Delete(task_id int64) error {
	for i, item := range r.items {
		if item.Id == task_id {
			r.items[i] = r.items[len(r.items)-1]
			r.items = r.items[:len(r.items)-1]
			return nil
		}
	}
	return errors.ErrNotFound
}

func TestService(t *testing.T) {
	s := service{&mockRepository{
		items: []entity.Task{
			task_examples[0],
		},
		task_id:  2,
		label_id: 2,
	}}

	t.Run("get", func(t *testing.T) {
		got, err := s.Get(1)
		want := []entity.Task{task_examples[0]}
		test.IsNil(t, err)
		test.DeepEqual(t, want, got)
	})

	t.Run("get_by_id", func(t *testing.T) {
		got, err := s.GetById(1)
		want := task_examples[0]

		test.IsNil(t, err)
		test.DeepEqual(t, want, got)
	})

	t.Run("create", func(t *testing.T) {
		got, err := s.Create(&CreateTaskRequest{
			UserId:               1,
			Name:                 "valid",
			Description:          "valid",
			CreatedOn:            "2000-01-01T00:00:00Z",
			DueDate:              "2000-01-01T00:00:00Z",
			SchtirlichHumorescue: "valid",
			Labels:               `[]`,
			Status:               entity.DONE,
		})
		want := entity.Task{
			Id:                   2,
			UserId:               1,
			Name:                 "valid",
			Description:          "valid",
			CreatedOn:            "2000-01-01T00:00:00Z",
			DueDate:              "2000-01-01T00:00:00Z",
			SchtirlichHumorescue: "valid",
			Labels:               `[]`,
			Status:               entity.DONE,
		}

		test.IsNil(t, err)
		test.DeepEqual(t, want, got)
	})

	t.Run("update", func(t *testing.T) {
		got, err := s.Update(&UpdateTaskRequest{
			TaskId:               2,
			Name:                 Name(task_examples[3].Name),
			Description:          Text(task_examples[3].Description),
			DueDate:              Date(task_examples[3].DueDate),
			SchtirlichHumorescue: Text(task_examples[3].SchtirlichHumorescue),
			Labels:               Labels(task_examples[3].Labels),
			Status:               Status(task_examples[3].Status),
		})
		want := task_examples[3]

		test.IsNil(t, err)
		test.DeepEqual(t, want, got)
	})

	t.Run("delete", func(t *testing.T) {
		got, err := s.Delete(2)
		want := task_examples[3]
		test.IsNil(t, err)
		test.DeepEqual(t, want, got)
		_, err = s.GetById(2)
		test.NotNil(t, err)
	})

	t.Run("get error", func(t *testing.T) {
		_, err := s.Get(2)
		test.NotNil(t, err)
	})

	t.Run("get by id error", func(t *testing.T) {
		_, err := s.GetById(2)
		test.NotNil(t, err)
	})

	t.Run("update not existent", func(t *testing.T) {
		_, err := s.Update(&UpdateTaskRequest{
			TaskId:               2,
			Name:                 "valid",
			Description:          "valid",
			DueDate:              "2000-01-01T00:00:00Z",
			SchtirlichHumorescue: "valid",
			Labels:               `[]`,
			Status:               entity.DONE,
		})

		test.NotNil(t, err)
	})

	t.Run("create invalid date", func(t *testing.T) {
		_, err := s.Create(&CreateTaskRequest{
			UserId:               1,
			Name:                 "valid",
			Description:          "valid",
			CreatedOn:            "cringe",
			DueDate:              "cringe",
			SchtirlichHumorescue: "valid",
			Labels:               `[]`,
			Status:               entity.DONE,
		})

		test.NotNil(t, err)
	})

	t.Run("create invalid label", func(t *testing.T) {
		_, err := s.Create(&CreateTaskRequest{
			UserId:               1,
			Name:                 "valid",
			Description:          "valid",
			CreatedOn:            "2000-01-01T00:00:00Z",
			DueDate:              "2000-01-01T00:00:00Z",
			SchtirlichHumorescue: "valid",
			Labels:               `[{"text":"valid","color":"cringe"}]`,
			Status:               entity.DONE,
		})

		test.NotNil(t, err)
	})

	t.Run("update invalid date", func(t *testing.T) {
		_, err := s.Update(&UpdateTaskRequest{
			TaskId:               1,
			Name:                 "valid",
			Description:          "valid",
			DueDate:              "cringe",
			SchtirlichHumorescue: "valid",
			Labels:               `[]`,
			Status:               entity.DONE,
		})

		test.NotNil(t, err)
	})

	t.Run("update invalid lable", func(t *testing.T) {
		_, err := s.Update(&UpdateTaskRequest{
			TaskId:               1,
			Name:                 "valid",
			Description:          "valid",
			DueDate:              "2000-01-01T00:00:00Z",
			SchtirlichHumorescue: "valid",
			Labels:               `[{"text":"valid","color":"cringe"}]`,
			Status:               entity.DONE,
		})

		test.NotNil(t, err)
	})
}
