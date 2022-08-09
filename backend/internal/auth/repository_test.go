package auth

import (
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/db"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/test"
	"testing"
)

func TestRepo(t *testing.T) {
	logger := log.New()
	db, err := db.New(logger)
	if err != nil {
		panic(err)
	}
	repo := NewRepository(db, logger)

	t.Run("get token", func(t *testing.T) {
		got, err := repo.GetToken("token", -1)
		test.IsNil(t, err)
		test.DeepEqual(t, DbToken{
			userId:       1,
			refreshToken: "token",
		}, got)
		_, err = repo.GetToken("bebra", -1)
		test.NotNil(t, err)
	})

	t.Run("create token", func(t *testing.T) {
		err := repo.CreateToken(2, "test")
		test.IsNil(t, err)
	})

	t.Run("check create token", func(t *testing.T) {
		got, err := repo.GetToken("", 2)
		test.IsNil(t, err)
		test.DeepEqual(t, DbToken{
			userId:       2,
			refreshToken: "test",
		}, got)
	})

	t.Run("update token", func(t *testing.T) {
		err = repo.UpdateToken(2, "ilovepotato")
		test.IsNil(t, err)
	})

	t.Run("check update token", func(t *testing.T) {
		got, err := repo.GetToken("ilovepotato", -1)
		test.IsNil(t, err)
		test.DeepEqual(t, DbToken{
			userId:       2,
			refreshToken: "ilovepotato",
		}, got)
	})
	t.Run("delete token", func(t *testing.T) {
		err = repo.DeleteToken("ilovepotato")
		test.IsNil(t, err)
	})
	t.Run("check delete token", func(t *testing.T) {
		_, err := repo.GetToken("ilovepotato", -1)
		test.NotNil(t, err)
	})
	t.Run("get user by email", func(t *testing.T) {
		got, err := repo.GetUser("test@example.com", -1)

		want := entity.User{
			Id:       1,
			Email:    "test@example.com",
			Nickname: "test",
			Password: "bd73d8db35a186a62c081da14526866c",
		}

		test.IsNil(t, err)
		test.DeepEqual(t, want, got)
	})
}
