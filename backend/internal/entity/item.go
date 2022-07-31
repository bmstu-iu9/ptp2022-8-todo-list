package entity

//Struct Item describing one item
type Item struct {
	ItemId      int    `json:"item_id"`
	ItemName    string `json:"item_name"`
	ImageSrc    string `json:"image_src"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
	Rarity      string `json:"rarity"`

	IsInInventory bool `json:"is_in_inventory"`
	IsEquipped    int  `json:"is_equipped"`
}
