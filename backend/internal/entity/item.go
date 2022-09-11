package entity

// Item represents single in-game item.
type Item struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	ImageSrc     string `json:"imageSrc"`
	ImageForHero string `json:"imageForHero"`
	Description  string `json:"description"`
	Price        int    `json:"price"`
	Category     string `json:"category"`
	Armor        int    `json:"armor"`
	Damage       int    `json:"damage"`
	Rarity       string `json:"rarity"`

	State ItemState `json:"state"`
}

// State represents state.
type State string

// ItemState represents item state.
type ItemState State

const (
	// Unknown shows that the ItemState is unknown for server.
	Unknown ItemState = ""
	// Equipped shows that user is using this item.
	Equipped ItemState = "equipped"
	// Inventoried shows that item is in inventory.
	Inventoried ItemState = "inventoried"
	// Store shows that item in a store
	Store ItemState = "store"
)
