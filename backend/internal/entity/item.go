package entity

//Struct Item describing one item
type Item struct {
	ItemId      int    `json:"itemId"`
	ItemName    string `json:"itemName"`
	ImageSrc    string `json:"imageSrc"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
	Rarity      string `json:"rarity"`
}
