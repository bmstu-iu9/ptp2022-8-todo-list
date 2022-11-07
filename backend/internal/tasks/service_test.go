package tasks

import (
	"testing"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/errors"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/test"
)

type mockRepository struct {
	items  []entity.Task
	taskId int64
}

func (r *mockRepository) Get(userId int64) ([]entity.Task, error) {
	res := []entity.Task{}
	for _, item := range r.items {
		if item.UserId == userId {
			res = append(res, item)
		}
	}

	if len(res) == 0 {
		return nil, errors.ErrNotFound
	}

	return res, nil
}

func (r *mockRepository) GetById(userId, taskId int64) (entity.Task, error) {
	for _, item := range r.items {
		if item.Id == taskId {
			return item, nil
		}
	}
	return entity.Task{}, errors.ErrNotFound
}

func (r *mockRepository) Create(task *entity.Task) error {
	task.Id = r.taskId
	r.taskId++
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

func (r *mockRepository) Delete(userId, taskId int64) error {
	for i, item := range r.items {
		if item.Id == taskId {
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
			taskExamples[0],
		},
		taskId: 2,
	}}

	t.Run("get", func(t *testing.T) {
		got, err := s.Get(taskExamples[0].UserId)
		want := []entity.Task{taskExamples[0]}
		test.IsNil(t, err)
		test.DeepEqual(t, want, got)
	})

	t.Run("get_by_id", func(t *testing.T) {
		got, err := s.GetById(taskExamples[0].UserId, taskExamples[0].Id)
		want := taskExamples[0]

		test.IsNil(t, err)
		test.DeepEqual(t, want, got)
	})

	t.Run("create", func(t *testing.T) {
		got, err := s.Set(&SetTaskRequest{
			Mode:                 CREATE,
			UserId:               1,
			Name:                 "valid",
			Description:          "valid",
			CreatedOn:            "2000-01-01T00:00:00Z",
			DueDate:              "2000-01-01T00:00:00Z",
			SchtirlichHumorescue: "valid",
			Labels:               `[]`,
			Status:               entity.COMPLETED,
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
			Status:               entity.COMPLETED,
		}

		test.IsNil(t, err)
		test.DeepEqual(t, want, got)
	})

	t.Run("update", func(t *testing.T) {
		got, err := s.Update(&UpdateTaskRequest{
			TaskId:               2,
			Name:                 Name(taskExamples[3].Name),
			Description:          Text(taskExamples[3].Description),
			DueDate:              Date(taskExamples[3].DueDate),
			SchtirlichHumorescue: Text(taskExamples[3].SchtirlichHumorescue),
			Labels:               Labels(taskExamples[3].Labels),
			Status:               Status(taskExamples[3].Status),
		})

		want := taskExamples[3]

		test.IsNil(t, err)
		test.DeepEqual(t, want, got)
	})

	t.Run("delete", func(t *testing.T) {
		got, err := s.Delete(taskExamples[3].UserId, taskExamples[3].Id)
		want := taskExamples[3]
		test.IsNil(t, err)
		test.DeepEqual(t, want, got)
		_, err = s.GetById(taskExamples[3].UserId, taskExamples[3].Id)
		test.NotNil(t, err)
	})

	t.Run("get error", func(t *testing.T) {
		_, err := s.Get(2)
		test.NotNil(t, err)
	})

	t.Run("get by id error", func(t *testing.T) {
		_, err := s.GetById(taskExamples[3].UserId, taskExamples[3].Id)
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
			Status:               entity.COMPLETED,
		})

		test.NotNil(t, err)
	})

	t.Run("create invalid date", func(t *testing.T) {
		_, err := s.Set(&SetTaskRequest{
			Mode:                 CREATE,
			UserId:               1,
			Name:                 "valid",
			Description:          "valid",
			CreatedOn:            "cringe",
			DueDate:              "cringe",
			SchtirlichHumorescue: "valid",
			Labels:               `[]`,
			Status:               entity.COMPLETED,
		})

		test.NotNil(t, err)
	})

	t.Run("create invalid label", func(t *testing.T) {
		_, err := s.Set(&SetTaskRequest{
			Mode:                 CREATE,
			UserId:               1,
			Name:                 "valid",
			Description:          "valid",
			CreatedOn:            "2000-01-01T00:00:00Z",
			DueDate:              "2000-01-01T00:00:00Z",
			SchtirlichHumorescue: "valid",
			Labels:               `[{"text":"valid","color":"cringe"}]`,
			Status:               entity.COMPLETED,
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
			Status:               entity.COMPLETED,
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
			Status:               entity.COMPLETED,
		})

		test.NotNil(t, err)
	})
}
