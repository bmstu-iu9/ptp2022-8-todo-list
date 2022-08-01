package entity

// A Item represents single in-game item.
type Item struct {
	ItemId      int    `json:"item_id"`
	ItemName    string `json:"item_name"`
	ImageSrc    string `json:"image_src"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
	Rarity      string `json:"rarity"`

	ItemState State `json:"item_state"`
}

//State represents ItemsState.
type State string

const (
	Unknown     State = ""
	Equipped    State = "equipped"
	Inventoried State = "inventoried"
	Store       State = "store"
)

type Filter struct {
	StateFilter    State
	RarityFilter   string
	CategoryFilter string
}
