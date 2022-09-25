package tasks

import (
	"testing"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/db"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/test"
)

var task_examples = []entity.Task{
	{ // 0
		Id:                   1,
		UserId:               1,
		Name:                 "test_name",
		Description:          "test_description",
		CreatedOn:            "2000-01-01T00:00:00Z",
		DueDate:              "2000-01-01T00:00:00Z",
		SchtirlichHumorescue: "test_humorescue",
		Labels:               `[{"text":"test_name","color":"#000000"}]`,
		Status:               entity.IN_PROGRESS,
	},
	{ // 1
		Id:                   1,
		UserId:               1,
		Name:                 "test_name_new",
		Description:          "test_description_new",
		CreatedOn:            "2000-01-01T00:00:00Z",
		DueDate:              "2000-01-01T00:00:00Z",
		SchtirlichHumorescue: "test_humorescue_new",
		Labels:               `[{"text":"test_name_new","color":"#ffffff"}]`,
		Status:               entity.DONE,
	},
	{ // 2
		Id:                   2,
		UserId:               1,
		Name:                 "test_name_stranger",
		Description:          "test_description_stranger",
		CreatedOn:            "2000-01-01T00:00:02Z",
		DueDate:              "2000-01-01T00:00:02Z",
		SchtirlichHumorescue: "test_humorescue_stranger",
		Labels:               `[{"text":"test_name_stranger","color":"#00ff00"}]`,
		Status:               entity.DONE,
	},
	{ // 3
		Id:                   2,
		UserId:               1,
		Name:                 "test_name_new_stranger",
		Description:          "test_description_new_stranger",
		CreatedOn:            "2000-01-01T00:00:00Z",
		DueDate:              "2000-01-01T00:00:02Z",
		SchtirlichHumorescue: "test_humorescue_new_stranger",
		Labels:               `[{"text":"test_name_new_stranger","color":"#00ffff"},{"text":"test_name_stranger","color":"#00ff00"}]`,
		Status:               entity.DONE,
	},
	{ // 4
		Id:                   1,
		UserId:               1,
		Name:                 "valid",
		Description:          "valid",
		CreatedOn:            "invalid",
		DueDate:              "invalid",
		SchtirlichHumorescue: "valid",
		Labels:               `[]`,
		Status:               entity.IN_PROGRESS,
	},
	{ // 5
		Id:                   1,
		UserId:               1,
		Name:                 "valid",
		Description:          "valid",
		CreatedOn:            "2000-01-01T00:00:00Z",
		DueDate:              "2000-01-01T00:00:00Z",
		SchtirlichHumorescue: "valid",
		Labels:               `[{"text":"valid","color":"invalid"}]`,
		Status:               entity.DONE,
	},
}

func TestRepo(t *testing.T) {
	logger := log.New()
	db, err := db.NewForTest(logger)
	if err != nil {
		panic(err)
	}

	r := NewRepository(db, logger)

	t.Run("get", func(t *testing.T) {
		got, err := r.Get(1)
		want := []entity.Task{task_examples[0]}

		test.IsNil(t, err)
		test.DeepEqual(t, want, got)
	})

	t.Run("get by id", func(t *testing.T) {
		got, err := r.GetById(1)
		want := task_examples[0]

		test.IsNil(t, err)
		test.DeepEqual(t, want, got)
	})

	t.Run("create", func(t *testing.T) {
		err := r.Create(&task_examples[2])

		test.IsNil(t, err)

		got, err := r.GetById(2)
		want := task_examples[2]

		test.IsNil(t, err)
		test.DeepEqual(t, want, got)
	})

	t.Run("update", func(t *testing.T) {
		err := r.Update(&task_examples[1])

		test.IsNil(t, err)

		got, err := r.GetById(1)
		want := task_examples[1]

		test.IsNil(t, err)
		test.DeepEqual(t, want, got)
	})

	t.Run("delete", func(t *testing.T) {
		err := r.Delete(1)
		test.IsNil(t, err)
		_, err = r.GetById(1)
		test.NotNil(t, err)
	})
}
