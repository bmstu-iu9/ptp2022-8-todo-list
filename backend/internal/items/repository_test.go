package items

import (
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/db"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
	. "gopkg.in/check.v1"
	"net/http"
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
	req, _ := http.NewRequest("GET", "/user/1/items?state=inventoried", nil)
	testItems, err := r.repo.GetAll(1, NewFilter(req))
	c.Check(err, IsNil)
	c.Check(testItems, DeepEquals, []entity.Item{
		{
			Id:           1,
			Name:         "testItem1",
			ImageSrc:     "test.png",
			ImageForHero: "test.png",
			Description:  "test1",
			Price:        65,
			Rarity:       "rare",
			Category:     "armor",
			State:        entity.Inventoried,
			Armor:        10,
			Damage:       10,
		},
	})
	req, _ = http.NewRequest("GET", "/user/1/items?rarity=rare&category=weapon", nil)
	testItems, err = r.repo.GetAll(1, NewFilter(req))
	c.Check(err, IsNil)
	c.Check(testItems, DeepEquals, []entity.Item{})
	req, _ = http.NewRequest("GET", "/user/1/items?category=armor", nil)
	testItems, err = r.repo.GetAll(1, NewFilter(req))
	c.Check(err, IsNil)
	c.Check(testItems, DeepEquals, []entity.Item{
		{
			Id:           1,
			Name:         "testItem1",
			ImageSrc:     "test.png",
			ImageForHero: "test.png",
			Description:  "test1",
			Price:        65,
			Rarity:       "rare",
			Category:     "armor",
			State:        entity.Inventoried,
			Armor:        10,
			Damage:       10,
		},
	})
}

func (r *RepositoryTestSuite) TestGetOne(c *C) {
	testItem, err := r.repo.GetOne(1, 1)
	c.Check(err, IsNil)
	c.Check(testItem, Equals, entity.Item{
		Id:           1,
		Name:         "testItem1",
		ImageSrc:     "test.png",
		ImageForHero: "test.png",
		Description:  "test1",
		Price:        65,
		Rarity:       "rare",
		Category:     "armor",
		State:        entity.Inventoried,
		Armor:        10,
		Damage:       10,
	})
	_, err = r.repo.GetOne(2, 2)
	c.Check(err, NotNil)
	_, err = r.repo.GetOne(1, 3)
	c.Check(err, NotNil)

	testItem, err = r.repo.GetOne(1, 2)
	c.Check(err, IsNil)
	c.Check(testItem, Equals, entity.Item{
		Id:           2,
		Name:         "testItem2",
		ImageSrc:     "test2.png",
		ImageForHero: "test2.png",
		Description:  "test2",
		Price:        62,
		Rarity:       "epic",
		Category:     "weapon",
		State:        entity.Equipped,
		Damage:       5,
	})
}

func (r *RepositoryTestSuite) TestUpdate(c *C) {
	item := entity.Item{
		Id:    1,
		State: entity.Equipped,
	}
	err := r.repo.Update(1, &item)
	c.Check(err, IsNil)
	testItems, err := r.repo.GetOne(1, 1)
	c.Check(err, IsNil)
	c.Check(testItems, Equals, entity.Item{
		Id:           1,
		Name:         "testItem1",
		ImageSrc:     "test.png",
		ImageForHero: "test.png",
		Description:  "test1",
		Price:        65,
		Rarity:       "rare",
		Category:     "armor",
		State:        entity.Equipped,
		Armor:        10,
		Damage:       10,
	})
}
