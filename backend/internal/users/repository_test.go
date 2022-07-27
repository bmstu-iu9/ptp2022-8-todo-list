package users

import (
	"database/sql"
	"log"
	"os"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	. "gopkg.in/check.v1"
)

type RepoTestSuite struct {
	repo repository
}

func init() {
	Suite(&RepoTestSuite{})
}

func (s *RepoTestSuite) SetUpSuite(c *C) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres dbname=slavatidika password=example sslmode=disable")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
DROP TABLE IF EXISTS users;
CREATE TABLE users (
       id serial PRIMARY KEY,
       email varchar(255) UNIQUE NOT NULL,
       nickname varchar(20) NOT NULL,
       password varchar(100) NOT NULL
);
INSERT INTO users(email, nickname, password)
VALUES('test@example.com', 'test', 'Test123Test');
`)
	if err != nil {
		panic(err)
	}

	s.repo = NewRepository(db, log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile))
}

func (s *RepoTestSuite) TestRepo(c *C) {
	tmp := s.repo.db.Ping()
	c.Check(tmp, IsNil)

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
