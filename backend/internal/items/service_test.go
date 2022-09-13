package items

import (
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/test"
	"testing"
)

type CRUDTestCase struct {
	Name    string
	Input   interface{}
	Want    entity.Item
	WantAll []entity.Item
	IsOK    bool
}

type GetAllTest struct {
	userId  int
	filters ItemFilter
}

type GetOneTest struct {
	userId int
	itemId int
}

type UpdateItemStateTest struct {
	userId int
	itemId int
	data   *UpdateItemStateRequest
}

func TestCRUD(t *testing.T) {
	service := NewService(NewMockRerository())
	tests := []CRUDTestCase{
		{
			Name: "get all ok", Input: GetAllTest{1, defaultItemFilter},
			WantAll: []entity.Item{
				{
					Id:       10,
					Name:     "sword",
					Rarity:   "rare",
					Category: "weapon",
					State:    entity.Equipped,
				},
				{
					Id:       6,
					Name:     "head",
					Rarity:   "legendary",
					Category: "armor",
					State:    entity.Equipped,
				},
			},
			IsOK: true,
		},
		{
			Name: "get all ok rarity filter", Input: GetAllTest{1, ItemFilter{RarityFilter: []string{"rare"}}},
			WantAll: []entity.Item{
				{
					Id:       10,
					Name:     "sword",
					Rarity:   "rare",
					Category: "weapon",
					State:    entity.Equipped,
				},
			},
			IsOK: true,
		},
		{
			Name: "get all ok filter", Input: GetAllTest{1, ItemFilter{RarityFilter: []string{"rare"},
				CategoryFilter: []string{"weapon"}}},
			WantAll: []entity.Item{
				{
					Id:       10,
					Name:     "sword",
					Rarity:   "rare",
					Category: "weapon",
					State:    entity.Equipped,
				},
			},
			IsOK: true,
		},
		{
			Name: "get all ok filter", Input: GetAllTest{1, ItemFilter{RarityFilter: []string{"rare"},
				CategoryFilter: []string{"armor"}}},
			WantAll: nil,
			IsOK:    true,
		},
		{
			Name: "get all fail filter", Input: GetAllTest{1,
				ItemFilter{RarityFilter: []string{"fail"}}},
			IsOK: false,
		},
		{
			Name: "get one ok", Input: GetOneTest{1, 6},
			Want: entity.Item{
				Id:       6,
				Name:     "head",
				Rarity:   "legendary",
				Category: "armor",
				State:    entity.Equipped,
			},
			IsOK: true,
		},
		{
			Name: "get one fail itemid", Input: GetOneTest{1, 7},
			IsOK: false,
		},
		{
			Name: "get one fail userid", Input: GetOneTest{2, 6},
			IsOK: false,
		},
		{
			Name: "update state ok", Input: UpdateItemStateTest{1, 6,
				&UpdateItemStateRequest{ItemState: entity.Store}},
			Want: entity.Item{
				Id:       6,
				Name:     "head",
				Rarity:   "legendary",
				Category: "armor",
				State:    entity.Store,
			},
			IsOK: true,
		},
		{
			Name: "update state fail", Input: UpdateItemStateTest{1, 6,
				&UpdateItemStateRequest{ItemState: "fail"}},
			IsOK: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			var (
				got    entity.Item
				gotArr []entity.Item
				err    error
			)
			switch tc.Input.(type) {
			case GetAllTest:
				gotArr, err = service.GetAll(tc.Input.(GetAllTest).userId, tc.Input.(GetAllTest).filters)
			case GetOneTest:
				got, err = service.GetOne(tc.Input.(GetOneTest).userId, tc.Input.(GetOneTest).itemId)
			case UpdateItemStateTest:
				got, err = service.UpdateItemState(tc.Input.(UpdateItemStateTest).userId,
					tc.Input.(UpdateItemStateTest).itemId, tc.Input.(UpdateItemStateTest).data)
			}
			if tc.IsOK {
				test.IsNil(t, err)
			} else {
				test.NotNil(t, err)
			}
			if tc.WantAll != nil {
				test.DeepEqual(t, tc.WantAll, gotArr)
			} else {
				test.DeepEqual(t, tc.Want, got)
			}
		})
	}
}
