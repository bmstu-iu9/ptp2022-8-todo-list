package users

import (
	"testing"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/errors"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/test"
)

type CRUDTestCase struct {
	Name  string
	Input interface{}
	Want  User
	IsOK  bool
}

type GetTestCase struct {
	Id int64
}

type CreateTestCase struct {
	Data *CreateUserRequest
}

type DeleteTestCase struct {
	Id int64
}

type UpdateTestCase struct {
	Id   int64
	Data *UpdateUserRequest
}

func TestCRUD(t *testing.T) {
	service := service{&mockRepository{
		items: []entity.User{
			{
				Id:       0,
				Email:    "slava@example.com",
				Nickname: "slavaruswarrior",
				Password: "wasdqwertytest",
			},
			{
				Id:       5,
				Email:    "geogreck@example.com",
				Nickname: "geogreck",
				Password: "test123test",
			},
		},
		id: 6,
	}}

	tests := []CRUDTestCase{
		{
			Name:  "get OK",
			Input: GetTestCase{0},
			Want:  User{0, "slava@example.com", "slavaruswarrior"},
			IsOK:  true,
		},
		{
			Name:  "get OK",
			Input: GetTestCase{5},
			Want:  User{5, "geogreck@example.com", "geogreck"},
			IsOK:  true,
		},
		{
			Name:  "get negative id",
			Input: GetTestCase{-10},
			Want:  User{},
			IsOK:  false,
		},
		{
			Name:  "get non-existing id",
			Input: GetTestCase{4},
			Want:  User{},
			IsOK:  false,
		},
		{
			Name: "create OK",
			Input: CreateTestCase{
				&CreateUserRequest{
					Email:    "stewkk@example.com",
					Nickname: "stewkk",
					Password: "oadfahdks",
				},
			},
			IsOK: true,
			Want: User{
				Id:       6,
				Email:    "stewkk@example.com",
				Nickname: "stewkk",
			},
		},
		{
			Name:  "create validate",
			Input: GetTestCase{6},
			Want:  User{6, "stewkk@example.com", "stewkk"},
			IsOK:  true,
		},
		{
			Name: "create validate error",
			Input: CreateTestCase{
				&CreateUserRequest{
					Email:    "stewkk",
					Nickname: "oadfahdks",
					Password: "",
				},
			},
			IsOK: false,
			Want: User{},
		},
		{
			Name:  "delete OK",
			Input: DeleteTestCase{5},
			Want:  User{5, "geogreck@example.com", "geogreck"},
			IsOK:  true,
		},
		{
			Name:  "delete validate",
			Input: GetTestCase{5},
			Want:  User{},
			IsOK:  false,
		},
		{
			Name:  "delete negative id",
			Input: DeleteTestCase{-123},
			Want:  User{},
			IsOK:  false,
		},
		{
			Name:  "delete non-existing id",
			Input: DeleteTestCase{123},
			Want:  User{},
			IsOK:  false,
		},
		{
			Name: "update email OK",
			Input: UpdateTestCase{
				Id: 0,
				Data: &UpdateUserRequest{
					Email:           newEmail("test@example.com"),
					CurrentPassword: "wasdqwertytest",
				},
			},
			Want: User{
				Id:       0,
				Email:    "test@example.com",
				Nickname: "slavaruswarrior",
			},
			IsOK: true,
		},
		{
			Name:  "update validate",
			Input: GetTestCase{0},
			Want: User{
				Id:       0,
				Email:    "test@example.com",
				Nickname: "slavaruswarrior",
			},
			IsOK: true,
		},
		{
			Name: "update password OK",
			Input: UpdateTestCase{
				Id: 0,
				Data: &UpdateUserRequest{
					NewPassword:     newPassword("test321test"),
					CurrentPassword: "wasdqwertytest",
				},
			},
			Want: User{
				Id:       0,
				Email:    "test@example.com",
				Nickname: "slavaruswarrior",
			},
			IsOK: true,
		},
		{
			Name: "update password validate",
			Input: UpdateTestCase{
				Id: 0,
				Data: &UpdateUserRequest{
					Nickname:        newNickname("example"),
					CurrentPassword: "test321test",
				},
			},
			Want: User{
				Id:       0,
				Email:    "test@example.com",
				Nickname: "example",
			},
			IsOK: true,
		},
		{
			Name: "update empty",
			Input: UpdateTestCase{
				Id: 0,
				Data: &UpdateUserRequest{
					CurrentPassword: "test321test",
				},
			},
			Want: User{
				Id:       0,
				Email:    "test@example.com",
				Nickname: "example",
			},
			IsOK: true,
		},
		{
			Name: "update wrong password",
			Input: UpdateTestCase{
				Id: 0,
				Data: &UpdateUserRequest{
					Nickname:        newNickname("example123"),
					CurrentPassword: "wrongPassword",
				},
			},
			Want: User{},
			IsOK: false,
		},
		{
			Name: "update validation error",
			Input: UpdateTestCase{
				Id: 0,
				Data: &UpdateUserRequest{
					Email:           newEmail("example123"),
					CurrentPassword: "test321test",
				},
			},
			Want: User{},
			IsOK: false,
		},
		{
			Name:  "update validate",
			Input: GetTestCase{0},
			Want: User{
				Id:       0,
				Email:    "test@example.com",
				Nickname: "example",
			},
			IsOK: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			var (
				got User
				err error
			)
			switch tc.Input.(type) {
			case GetTestCase:
				got, err = service.Get(tc.Input.(GetTestCase).Id)
			case CreateTestCase:
				got, err = service.Create(tc.Input.(CreateTestCase).Data)
			case DeleteTestCase:
				got, err = service.Delete(tc.Input.(DeleteTestCase).Id)
			case UpdateTestCase:
				input := tc.Input.(UpdateTestCase)
				got, err = service.Update(input.Id, input.Data)
			}

			if tc.IsOK {
				test.IsNil(t, err)
			} else {
				test.NotNil(t, err)
			}
			test.DeepEqual(t, tc.Want, got)
		})
	}
}

type mockRepository struct {
	items []entity.User
	id    int64
}

func (repo *mockRepository) InitUserInventory(id int64) error {
	return nil
}

func (repo *mockRepository) CleanUserInventory(id int64) error {
	return nil
}

func (repo *mockRepository) Create(user *entity.User) error {
	user.Id = repo.id
	repo.id++
	repo.items = append(repo.items, *user)
	return nil
}

func (repo mockRepository) Get(id int64) (entity.User, error) {
	for _, item := range repo.items {
		if item.Id == id {
			return item, nil
		}
	}
	return entity.User{}, errors.ErrNotFound
}

func (repo *mockRepository) Delete(id int64) error {
	for i, item := range repo.items {
		if item.Id == id {
			repo.items[i] = repo.items[len(repo.items)-1]
			repo.items = repo.items[:len(repo.items)-1]
			return nil
		}
	}
	return nil
}
func (repo mockRepository) Update(user *entity.User) error {
	for i, item := range repo.items {
		if item.Id == user.Id {
			repo.items[i] = *user
			return nil
		}
	}
	return errors.ErrNotFound
}
