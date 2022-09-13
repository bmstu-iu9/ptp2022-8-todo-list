package items

import (
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/db"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/test"
	"net/http"
	"testing"
)

var testsDbValues = []entity.Item{
	{
		Id:           1,
		Name:         "testItem1",
		ImageSrc:     "test.png",
		ImageForHero: "test.png",
		Description:  "test1",
		Price:        65,
		Category:     "armor",
		Rarity:       "rare",
		Damage:       10,
		Armor:        10,
		State:        entity.Inventoried,
	},
	{
		Id:           1,
		Name:         "testItem1",
		ImageSrc:     "test.png",
		ImageForHero: "test.png",
		Description:  "test1",
		Price:        65,
		Category:     "armor",
		Rarity:       "rare",
		Damage:       10,
		Armor:        10,
		State:        entity.Equipped,
	},
	{
		Id:           2,
		Name:         "testItem2",
		ImageSrc:     "test2.png",
		ImageForHero: "test2.png",
		Description:  "test2",
		Price:        62,
		Category:     "weapon",
		Rarity:       "epic",
		Damage:       5,
		Armor:        0,
		State:        entity.Equipped,
	},
}

var defaultItemFilter = ItemFilter{
	RarityFilter:   []string{"common", "rare", "epic", "legendary"},
	CategoryFilter: []string{"armor", "weapon", "pet", "skin"},
	StateFilter:    []entity.ItemState{"equipped", "inventoried", "store"},
}

func TestRepo(t *testing.T) {
	logger := log.New()
	db, err := db.New(logger)
	if err != nil {
		panic(err)
	}
	repo := NewRepository(db, logger)

	t.Run("get all ok", func(t *testing.T) {
		got, err := repo.GetAll(1, defaultItemFilter)
		test.IsNil(t, err)
		test.DeepEqual(t, []entity.Item{
			testsDbValues[0], testsDbValues[2],
		}, got)
	})

	t.Run("get all ok with filters", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/user/1/items?state=inventoried", nil)
		got, err := repo.GetAll(1, NewFilter(req))
		test.IsNil(t, err)
		test.DeepEqual(t, []entity.Item{
			testsDbValues[0],
		}, got)

		req, _ = http.NewRequest("GET", "/user/1/items?state=inventoried&rarity=legendary", nil)
		got, err = repo.GetAll(1, NewFilter(req))
		test.IsNil(t, err)
		test.DeepEqual(t, []entity.Item{}, got)
	})

	t.Run("get one", func(t *testing.T) {
		got, err := repo.GetOne(1, 1)
		test.IsNil(t, err)
		test.DeepEqual(t, testsDbValues[0], got)
		_, err = repo.GetOne(1, 3)
		test.NotNil(t, err)
		_, err = repo.GetOne(3, 1)
		test.NotNil(t, err)
	})

	t.Run("modify state", func(t *testing.T) {
		testsDbValues[0].State = entity.Store
		err = repo.Update(1, &testsDbValues[0])
		test.IsNil(t, err)
		got, err := repo.GetOne(1, 1)
		test.IsNil(t, err)
		test.DeepEqual(t, testsDbValues[0], got)
	})
}
