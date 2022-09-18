type Equipment = {
    helmet: Item
    leggins: Item
    chest: Item
    weapon: Item
    boots: Item
    pet: Item
}

function getEquipment(items: Item[]): Equipment {
    let equipment: Equipment = {
        helmet: null,
        leggins: null,
        chest: null,
        weapon: null,
        boots: null,
        pet: null,
    }
    items.forEach((el) => {
        if (el.state === 'equipped') {
            switch (el.category) {
                case 'helmet':
                    equipment.helmet = el
                    break
                case 'chest':
                    equipment.chest = el
                    break
                case 'leggins':
                    equipment.leggins = el
                    break
                case 'boots':
                    equipment.boots = el
                    break
                case 'weapon':
                    equipment.weapon = el
                case 'pet':
                    equipment.pet = el
                    break
            }
        }
    })
    return equipment
}

function createEquipHtml(item: Item): HTMLElement {
    let html = document.createElement('img')
    html.setAttribute('class', 'img-fluid')
    if (item.category == 'pet') {
        html.setAttribute('src', 'https://wg.grechkogv.ru/assets/' + item.imageSrc)
    } else {
        html.setAttribute('src', item.imageForHero)
    }
    html.setAttribute('id', item.category)
    return html
}

function setEquipmentImg(item: Item): void {
    if (item.category != 'pet') {
        let equip: HTMLElement = document.getElementById(item.category)!
        if (equip != null) {
            equip.setAttribute('src', item.imageForHero)
        } else {
            let hero: HTMLElement = document.querySelector('.hero') as HTMLElement
            hero.prepend(createEquipHtml(item))
        }
    }
}

function removeEquipmentImg(category: Category) {
    if (category != 'pet') {
        document.getElementById(category)?.remove()
    }
}

function getHeroHtml(e: Equipment): HTMLElement {
    let result: HTMLElement = document.createElement('div')
    let avatar: HTMLElement = document.createElement('img')
    result.setAttribute('class', 'hero')
    avatar.setAttribute('class', 'img-fluid')
    avatar.setAttribute('id', 'avatar')
    avatar.setAttribute('src', 'https://static.wikia.nocookie.net/minecraft_ru_gamepedia/images/3/33/Стив_JE2_BE1.png')
    avatar.setAttribute('alt', 'avatar')
    for (let key of Object.keys(e)) {
        if (e[key] && key != 'pet') {
            result.appendChild(createEquipHtml(e[key]))
        }
    }
    result.appendChild(avatar)
    return result
}

function renderHero(parentElementId: string, equipment: Equipment): void {
    const parentElement: HTMLElement = document.getElementById(parentElementId) as HTMLElement
    if (parentElement != null) {
        parentElement.prepend(getHeroHtml(equipment))
    }
}
