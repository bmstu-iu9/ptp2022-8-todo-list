package users

import (
	"testing"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/db"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/test"
)

func TestRepo(t *testing.T) {
	logger := log.New()
	db, err := db.New(logger)
	if err != nil {
		panic(err)
	}

	repo := NewRepository(db, logger)

	t.Run("get", func(t *testing.T) {
		got, err := repo.Get(1)

		want := entity.User{
			Id:       1,
			Email:    "test@example.com",
			Nickname: "test",
			Password: "Test123Test",
		}
		test.IsNil(t, err)
		test.DeepEqual(t, want, got)
	})
	user := &entity.User{
		Id:       0,
		Email:    "slava@example.com",
		Nickname: "slavaruswarrior",
		Password: "Ryudfnsb675",
	}

	t.Run("create", func(t *testing.T) {
		err = repo.Create(user)

		test.IsNil(t, err)
	})

	t.Run("create validate", func(t *testing.T) {
		got, err := repo.Get(user.Id)

		want := *user

		test.IsNil(t, err)
		test.DeepEqual(t, want, got)
		if user.Id != 3 {
			t.Fatalf("expected user.Id: 2, got: %#v", got)
		}
	})

	t.Run("update", func(t *testing.T) {
		user.Email = "example@example.com"
		err = repo.Update(user)

		test.IsNil(t, err)
	})

	t.Run("update validate", func(t *testing.T) {
		got, err := repo.Get(user.Id)

		want := *user

		test.IsNil(t, err)
		test.DeepEqual(t, want, got)
	})

	t.Run("delete", func(t *testing.T) {
		err = repo.Delete(user.Id)

		test.IsNil(t, err)
	})

	t.Run("delete validate", func(t *testing.T) {
		_, err = repo.Get(user.Id)

		test.NotNil(t, err)
	})
}
