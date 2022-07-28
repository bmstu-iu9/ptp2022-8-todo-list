package users

import (
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/db"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
	. "gopkg.in/check.v1"
)

type RepoTestSuite struct {
	repo Repository
}

func init() {
	Suite(&RepoTestSuite{})
}

func (s *RepoTestSuite) SetUpSuite(c *C) {
	db, err := db.New()
	if err != nil {
		panic(err)
	}

	s.repo = NewRepository(db, log.New())
}

func (s *RepoTestSuite) TestRepo(c *C) {
	testUser, err := s.repo.Get(1)
	c.Check(err, IsNil)
	c.Check(testUser, DeepEquals, entity.User{
		Id: 1,
		Email: "test@example.com",
		Nickname: "test",
		Password: "Test123Test",
	})

	user := &entity.User{
		Email:    "slava@example.com",
		Nickname: "slavaruswarrior",
		Password: "Ryudfnsb675",
	}
	err = s.repo.Create(user)
	c.Check(err, IsNil)
	got, err := s.repo.Get(user.Id)
	c.Check(err, IsNil)
	c.Check(*user, DeepEquals, got)
	c.Check(user.Id, Equals, int64(2))

	user.Email = "example@example.com"
	err = s.repo.Update(user)
	c.Check(err, IsNil)
	got, err = s.repo.Get(user.Id)
	c.Check(err, IsNil)
	c.Check(got.Email, Equals, "example@example.com")

	err = s.repo.Delete(user.Id)
	c.Check(err, IsNil)
	_, err = s.repo.Get(user.Id)
	c.Check(err, NotNil)
}
