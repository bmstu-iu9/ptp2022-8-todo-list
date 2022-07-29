package items

import (
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/db"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
	. "gopkg.in/check.v1"
)

type RepositoryTestSuite struct {
	repo Repository
}

func init() {
	Suite(&RepositoryTestSuite{})
}

func (r *RepositoryTestSuite) SetUpTest(c *C) {
	logger := log.New()
	db, err := db.New(logger)
	if err != nil {
		panic(err)
	}
	r.repo = NewRepository(db, logger)
}

func (r *RepositoryTestSuite) TestGetAll(c *C) {
	testItems, err := r.repo.GetAll()
	c.Check(err, IsNil)
	c.Check(testItems, DeepEquals, []Item{
		{
			ItemId:   1,
			ItemName: "testItem1",
		},
		{
			ItemId:   2,
			ItemName: "testItem2",
		},
	})
}

func (r *RepositoryTestSuite) TestGetOne(c *C) {
	testItem, err := r.repo.GetOne(1, 1)
	c.Check(err, IsNil)
	c.Check(testItem, Equals, Item{
		ItemId:   1,
		ItemName: "Sasha",
	})
	testItem, err = r.repo.GetOne(2, 2)
	c.Check(err, NotNil)
	testItem, err = r.repo.GetOne(1, 3)
	c.Check(err, NotNil)

	testItem, err = r.repo.GetOne(1, 2)
	c.Check(err, IsNil)
	c.Check(testItem, Equals, Item{
		ItemId:   2,
		ItemName: "testItem2",
	})
}

func (r *RepositoryTestSuite) TestUpdate(c *C) {
	item := Item{
		ItemId:   1,
		ItemName: "Sasha",
	}
	err := r.repo.Update(item)
	c.Check(err, IsNil)
	testItems, err := r.repo.GetAll()
	c.Check(err, IsNil)
	c.Check(testItems, DeepEquals, []Item{
		{
			ItemId:   1,
			ItemName: "Sasha",
		},
		{
			ItemId:   2,
			ItemName: "testItem2",
		},
	})
}
