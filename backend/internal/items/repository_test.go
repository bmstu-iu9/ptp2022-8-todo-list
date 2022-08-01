package items

import (
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/db"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
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
	c.Check(testItems, DeepEquals, []entity.Item{
		{
			ItemId:      1,
			ItemName:    "testItem1",
			ImageSrc:    "test.png",
			Description: "test1",
			Price:       65,
			Rarity:      "rare",
			Category:    "armor",
		},
		{
			ItemId:      2,
			ItemName:    "testItem2",
			ImageSrc:    "test2.png",
			Description: "test2",
			Price:       62,
			Rarity:      "epic",
			Category:    "weapon",
		},
	})
}

func (r *RepositoryTestSuite) TestGetOne(c *C) {
	testItem, err := r.repo.GetOne(2, 1)
	c.Check(err, IsNil)
	c.Check(testItem, Equals, entity.Item{
		ItemId:      1,
		ItemName:    "testItem1",
		ImageSrc:    "test.png",
		Description: "test1",
		Price:       65,
		Rarity:      "rare",
		Category:    "armor",
		ItemState:   entity.Equipped,
	})
	_, err = r.repo.GetOne(2, 2)
	c.Check(err, NotNil)
	_, err = r.repo.GetOne(1, 3)
	c.Check(err, NotNil)

	testItem, err = r.repo.GetOne(1, 2)
	c.Check(err, IsNil)
	c.Check(testItem, Equals, entity.Item{
		ItemId:      2,
		ItemName:    "testItem2",
		ImageSrc:    "test2.png",
		Description: "test2",
		Price:       62,
		Rarity:      "epic",
		Category:    "weapon",
		ItemState:   entity.Inventoried,
	})
}

func (r *RepositoryTestSuite) TestUpdate(c *C) {
	item := entity.Item{
		ItemId:    1,
		ItemState: entity.Equipped,
	}
	err := r.repo.Update(1, &item)
	c.Check(err, IsNil)
	testItems, err := r.repo.GetOne(1, 1)
	c.Check(err, IsNil)
	c.Check(testItems, Equals, entity.Item{
		ItemId:      1,
		ItemName:    "testItem1",
		ImageSrc:    "test.png",
		Description: "test1",
		Price:       65,
		Rarity:      "rare",
		Category:    "armor",
		ItemState:   entity.Equipped,
	})
}
