type Item = {
    id: number
    name: string
    imageSrc: string
    description: string
    price: number
    category: Category
    rarity: Rarity
    state: ItemState
}

type Rarity = 'common' | 'rare' | 'epic' | 'legendary'

function color(rarity: Rarity): string {
    if (rarity === "common") return "#C8C8C8"
    else if (rarity === 'rare') return "#2bfff4"
    else if (rarity === 'epic') return "#f04dff"
    return "#linear-gradient(#40E0D0, #91e047, #fff456, #fff456, #ffa856, #e64f4f)"
}

type ItemState = "store" | "equipped" | "inventoried"
type Category = 'helmet' | 'chest' | 'leggins' | 'boots' | 'weapon' | 'pet' | 'armor'
