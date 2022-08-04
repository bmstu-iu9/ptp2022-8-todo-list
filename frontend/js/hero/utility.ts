type Item = {
    ItemId: number,
    ItemName: string,
    ImageSrc:  string,
    ImageForHero: string;
    Description: string,
    Price: number,
    Rarity: string,
    Category: string,
    ItemState: string,
}


type Equipment = {
    helmet: Item;
    leggings: Item;
    chestplate: Item;
    boots: Item;
    pet: Item;
}