package items

import (
	. "gopkg.in/check.v1"
)

type ServiceTestSuite struct {
	service Service
}

func init() {
	Suite(&ServiceTestSuite{})
}

func (s *ServiceTestSuite) SetUpTest(c *C) {
	s.service = NewService(NewMockRerository())
}

func (s *ServiceTestSuite) TestGetAll(c *C) {
	items, err := s.service.GetAll()
	c.Check(err, IsNil)
	c.Check(items, DeepEquals, []Item{
		{
			ItemId:   10,
			ItemName: "sword",
		},
		{
			ItemId:   6,
			ItemName: "head",
		},
	})
}

func (s *ServiceTestSuite) TestGetOne(c *C) {
	item, err := s.service.GetOne(1, 10)
	c.Check(err, IsNil)
	c.Check(item, Equals, Item{
		ItemId:   10,
		ItemName: "sword",
	})
	item, err = s.service.GetOne(1, 6)
	c.Check(err, IsNil)
	c.Check(item, Equals, Item{
		ItemId:   6,
		ItemName: "head",
	})
	_, err = s.service.GetOne(10, 6)
	c.Check(err, NotNil)
	_, err = s.service.GetOne(1, 5)
	c.Check(err, NotNil)
}

func (s *ServiceTestSuite) TestModify(c *C) {
	newItem := UpdateItemRequest{
		ItemName: "Sasha",
	}
	for i := 0; i < 2; i++ {
		item, err := s.service.Modify(1, 10, &newItem)
		c.Check(err, IsNil)
		c.Check(item, Equals, Item{
			ItemId:   10,
			ItemName: "Sasha",
		})
	}
	newItem.ItemName = ""
	item, err := s.service.Modify(1, 10, &newItem)
	c.Check(err, IsNil)
	c.Check(item, Equals, Item{
		ItemId:   10,
		ItemName: "Sasha",
	})
}
