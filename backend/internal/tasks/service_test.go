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
			r.items[i].Status = task.Status

			for _, lb := range task.Labels {
				if lb.Id == 0 {
					lb.Id = r.label_id
					r.label_id++
					r.items[i].Labels = append(r.items[i].Labels, lb)
				} else {
					for j, dlb := range r.items[i].Labels {
						if dlb.Id == lb.Id {
							r.items[i].Labels[j] = r.items[i].Labels[len(r.items[i].Labels)-1]
							r.items[i].Labels = r.items[i].Labels[:len(r.items[i].Labels)-1]
						}
					}
				}
			}
			return nil
		}
	}
	return errors.ErrNotFound
}

func toLabels(labels entity.Labels) Labels {
	lbs := Labels{}
	for _, lb := range labels {
		lbs = append(lbs, Label{
			Id:     lb.Id,
			TaskId: lb.TaskId,
			Name:   Name(lb.Name),
			Color:  Color(lb.Color),
		})
	}
	return lbs
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
			UserId:               task_examples[3].UserId,
			Name:                 Name(task_examples[3].Name),
			Description:          Text(task_examples[3].Description),
			CreatedOn:            Date(task_examples[3].CreatedOn),
			DueDate:              Date(task_examples[3].DueDate),
			SchtirlichHumorescue: Text(task_examples[3].SchtirlichHumorescue),
			Labels:               toLabels(task_examples[3].Labels),
			Status:               Status(task_examples[3].Status),
		})
		want := task_examples[3]

		test.IsNil(t, err)
		test.DeepEqual(t, want, got)
	})

	t.Run("update", func(t *testing.T) {
		got, err := s.Update(&UpdateTaskRequest{
			TaskId:               1,
			Name:                 Name(task_examples[2].Name),
			Description:          Text(task_examples[2].Description),
			DueDate:              Date(task_examples[2].DueDate),
			SchtirlichHumorescue: Text(task_examples[2].SchtirlichHumorescue),
			Labels:               toLabels(task_examples[2].Labels),
			Status:               Status(task_examples[2].Status),
		})
		want := task_examples[1]

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
}
