package users

import (
	"errors"
	"testing"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	. "gopkg.in/check.v1"
)

type UsersTestSuite struct {
	service service
}

type ValidateTestSuite struct {}

func init() {
	Suite(&UsersTestSuite{})
}

func Test(t *testing.T) { TestingT(t) }

func (s *ValidateTestSuite) TestCreateUserRequest_Validate(c *C) {
	req := CreateUserRequest{
		Email: "slava@example.com",
		Nickname: "slavaruswarrior",
		Password: "Asjh2k123",
	}
	c.Check(req.Validate(), IsNil)

	req = CreateUserRequest{
		Email: "email@test.com.",
		Nickname: "slavaruswarrior",
		Password: "Asjh2k123",
	}
	c.Check(req.Validate(), NotNil)

	req = CreateUserRequest{
		Email: "slava@example.com",
		Nickname: "sl",
		Password: "Asjh2k123",
	}
	c.Check(req.Validate(), NotNil)

	req = CreateUserRequest{
		Email: "slava@example.com",
		Nickname: "12345",
		Password: "Asjh2k123",
	}
	c.Check(req.Validate(), IsNil)

	req = CreateUserRequest{
		Email: "slava@example.com",
		Nickname: "-slavaruswarrior",
		Password: "Asjh2k123",
	}
	c.Check(req.Validate(), NotNil)

	req = CreateUserRequest{
		Email: "slava@example.com",
		Nickname: "slavaruswarrior",
		Password: "123",
	}
	c.Check(req.Validate(), NotNil)

	req = CreateUserRequest{
		Email: "slava@example.com",
		Nickname: "slavaruswarrior",
		Password: "dsfkskfhs^3dsfsf",
	}
	c.Check(req.Validate(), NotNil)
}

func (s *ValidateTestSuite) TestUpdateUserRequest_Validate(c *C) {

}

func (s *UsersTestSuite) SetUpTest(c *C) {
	s.service = service{NewMockRerository()}
}

func (s *UsersTestSuite) TestGet(c *C) {
	user, err := s.service.Get(0)
	c.Check(user, DeepEquals,
		User{0, "slava@example.com", "slavaruswarrior"})
	c.Check(err, IsNil)

	user, err = s.service.Get(5)
	c.Check(user, DeepEquals,
		User{5, "geogreck@example.com", "geogreck"})
	c.Check(err, IsNil)

	_, err = s.service.Get(-10)
	c.Check(err, NotNil)

	_, err = s.service.Get(4)
	c.Check(err, NotNil)
}

func (s *UsersTestSuite) TestCreate(c *C) {
	user, err := s.service.Create(&CreateUserRequest{
		Email:    "stewkk@example.com",
		Nickname: "stewkk",
		Password: "oadfahdks",
	})
	c.Check(err, IsNil)
	userGet, err := s.service.Get(user.Id)
	c.Check(user, DeepEquals, userGet)
	c.Check(err, IsNil)

	user, err = s.service.Create(&CreateUserRequest{
		Email:    "stewkkample.com",
		Nickname: "stewkk",
		Password: "oadfahdks",
	})
	c.Check(err, NotNil)

	user, err = s.service.Create(&CreateUserRequest{
		Nickname: "stewkk",
		Password: "oadfahdks",
	})
	c.Check(err, NotNil)
}

func (s *UsersTestSuite) TestDelete(c *C) {
	user, err := s.service.Delete(5)
	c.Check(user, DeepEquals,
		User{5, "geogreck@example.com", "geogreck"})
	c.Check(err, IsNil)

	userGet, err := s.service.Get(5)
	c.Check(userGet, Not(DeepEquals), user)
	c.Check(err, NotNil)

	_, err = s.service.Delete(-123)
	c.Check(err, NotNil)
}

func (s *UsersTestSuite) TestUpdateOK(c *C) {
	email := "test@example.com"
	user, err := s.service.Update(5, &UpdateUserRequest{
		Email:           &email,
		CurrentPassword: "test123test",
	})
	entityUser, _ := s.service.repo.Get(5)
	c.Check(entityUser, DeepEquals, entity.User{
		Id:       5,
		Email:    "test@example.com",
		Nickname: "geogreck",
		Password: "test123test",
	})
	c.Check(err, IsNil)
	userGet, _ := s.service.Get(5)
	c.Check(user, DeepEquals, userGet)

	password := "test321test"
	user, err = s.service.Update(5, &UpdateUserRequest{
		Email:           &email,
		NewPassword:     &password,
		CurrentPassword: "test123test",
	})
	entityUser, _ = s.service.repo.Get(5)
	c.Check(entityUser, DeepEquals, entity.User{
		Id:       5,
		Email:    "test@example.com",
		Nickname: "geogreck",
		Password: "test321test",
	})
	c.Check(err, IsNil)
	userGet, _ = s.service.Get(5)
	c.Check(user, DeepEquals, userGet)

	user, err = s.service.Update(5, &UpdateUserRequest{
		CurrentPassword: "test321test",
	})
	c.Check(err, IsNil)
	userGet, err = s.service.Get(5)
	c.Check(userGet, DeepEquals, user)
	c.Check(err, IsNil)
}

func (s *UsersTestSuite) TestUpdateError(c *C) {
	email := "test@example.com"
	_, err := s.service.Update(5, &UpdateUserRequest{
		Email:           &email,
		CurrentPassword: "wrongPassword",
	})
	c.Check(err, NotNil)

	email = "1234"
	_, err = s.service.Update(5, &UpdateUserRequest{
		Email:           &email,
		CurrentPassword: "test123",
	})
	c.Check(err, NotNil)
}

type mockRepository struct {
	items []entity.User
	id    int64
}

func NewMockRerository() *mockRepository {
	return &mockRepository{
		items: []entity.User{
			{
				Id:       0,
				Email:    "slava@example.com",
				Nickname: "slavaruswarrior",
				Password: "wasdqwertytest",
			},
			{
				Id: 5,
				Email: "geogreck@example.com",
				Nickname: "geogreck",
				Password: "test123test",
			},
		},
		id: 6,
	}
}

func (repo *mockRepository) Create(user *entity.User) error {
	user.Id = repo.id
	repo.id++
	repo.items = append(repo.items, *user)
	return nil
}

func (repo *mockRepository) Get(id int64) (entity.User, error) {
	for _, item := range repo.items {
		if item.Id == id {
			return item, nil
		}
	}
	return entity.User{}, errors.New("repo: can't find User with given id")
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
func (repo *mockRepository) Update(user *entity.User) error {
	for i, item := range repo.items {
		if item.Id == user.Id {
			repo.items[i] = *user
			return nil
		}
	}
	return errors.New("repo: can't find User with given id")
}
