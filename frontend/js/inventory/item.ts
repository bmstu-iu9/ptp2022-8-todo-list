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

function getRarityColor(rarity: Rarity): string {
    switch (rarity) {
        case 'common':
            return '#C8C8C8'
        case 'rare':
            return '#0b9ccf'
        case 'epic':
            return '#d461f7'
        case 'legendary':
            return '#ff9c00'
        default:
            return 'linear-gradient(#40E0D0, #91e047, #fff456, #fff456, #ffa856, #e64f4f)'
    }
}

type ItemState = 'store' | 'equipped' | 'inventoried'
type Category = 'helmet' | 'chest' | 'leggins' | 'boots' | 'weapon' | 'pet'
