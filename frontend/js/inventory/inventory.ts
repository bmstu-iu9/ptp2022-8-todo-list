// модальная форма просмотра описания карты
const modalInventory = new bootstrap.Modal(<HTMLFormElement>document.getElementById('inventoryModal'))
// хранилище предметов
var itemsInventory = new Map<number, Item>()
// хранилище айди надетых предметов по категориям(если предмет какой-то категории не надет, то значение -1)
let Equipped = {
    helmet: -1,
    chest: -1,
    leggins: -1,
    boots: -1,
    weapon: -1,
    pet: -1,
    skin: -1,
}

// получение предметов с сервера
sendRequest('GET', server + '/items').then((data) => {
    for (let i = 0; i < data.length; i++) {
        let item: Item = data[i]
        if (item.state !== 'store') {
            toInventoryHTMLBlock(item)
            itemsInventory.set(item.id, item)
            if (item.state === 'equipped') {
                equipped(item)
            }
        }
    }
    if (Equipped.skin !== -1) {
        let userImg = document.getElementById('inventory__user-img')
        let id = Equipped.skin
        let item = itemsInventory.get(id)!
        //userImg!.setAttribute('src', `https://wg.grechkogv.ru/assets/${item.imageSrc}`)
    }
})

// Общая обработка кликов по странице
document.addEventListener('click', (e) => {
    const target = <HTMLElement>e.target
    // получение из хранилища предмета, на карту которого кликнули
    let regexp = /idItem=\d+/
    let strId = findID(target, regexp)
    let id = parseInt(strId!.substring(7))
    let item = itemsInventory.get(id)!

    if (hasParentClass(target, 'inventory__box-title')) {
        // просмотр описания карты
        let windowForm = <HTMLElement>document.getElementsByClassName('inventory__form')[0]
        let buf = (<HTMLElement>document.getElementsByClassName(`idItem=${id}`)[0]).getBoundingClientRect()
        windowForm.style.top = `${buf.y}px`
        windowForm.style.left = `calc(${buf.x}px - 1vw - 1.4rem)`
        windowForm.style.width = `calc(${buf.right - buf.x}px + 2vw + 2.8rem)`
        windowForm.style.borderColor = `${getRarityColor(item.rarity)}`
        modalInventory.show()
        let titleModal = <HTMLInputElement>document.getElementById('inventoryModalTitle')
        titleModal.innerHTML = item.name!
        let description = <HTMLInputElement>document.getElementById('inventoryModalBoby')
        description.innerHTML = `<img src="https://wg.grechkogv.ru/assets/${item.imageSrc}"
    class="px-3 pb-3 img-fluid">`
        description.innerHTML += item.description! !== '' ? item.description! : `Описание отсутствует`
    } else if (hasParentClass(target, 'inventory__item-btn')) {
        // надеть/снять предмет
        if (item.state === 'equipped') {
            takeOff(item)
            unEquipped(item)
        } else {
            let idE = idEquipped(item)
            if (idE !== -1) {
                let itemEquipped = itemsInventory.get(idE)!
                takeOff(itemEquipped)
            }
            putOn(item)
        }
    }

})
// создание карты предмета в HTML
function createInventoryHTMLBlock(item: Item) {
    let str: string
    str = `<div class="col my-2 idItem=${item.id}">
        <div class="card inventory__item ${item.state === 'equipped' ? 'bg-success' : 'bg-secondary'} p-2 h-100">
                <img src="https://wg.grechkogv.ru/assets/${item.imageSrc}"
                    class="card-img-top bg-light inventory__box-img p-2">
            <div class="card-body inventory__box-title btn text-dark" title="Посмотреть описание"
                style="background: ${getRarityColor(item.rarity)}">
                <p class="card-title lh-1 align-middle m-0 p-0">${item.name}</p>
            </div>
            <div class="card-footer p-0">
            ${
                item.state === 'equipped'
                    ? `<a class="inventory__item-btn btn btn-success w-100 h-100">Снять</a>`
                    : `<a class="inventory__item-btn btn btn-primary w-100 h-100">Надеть</a>`
            }
            </div>
        </div>
    </div>`
    return str

}
// добавление HTML-блока карты на страницу
function toInventoryHTMLBlock(item: Item) {
    let str = createInventoryHTMLBlock(item)
    let buf = document.querySelector('.inventory__box__items')!.innerHTML
    document.querySelector('.inventory__box__items')!.innerHTML = buf.concat(str)
}

// добавление айди предмета item в хранилище айди надетых предметов
function equipped(item: Item) {
    switch (item.category) {
        case 'helmet':
            Equipped.helmet = item.id
            break
        case 'chest':
            Equipped.chest = item.id
            break
        case 'leggins':
            Equipped.leggins = item.id
            break
        case 'boots':
            Equipped.boots = item.id
            break
        case 'weapon':
            Equipped.weapon = item.id
            break
        case 'pet':
            Equipped.pet = item.id
            break
        case 'skin': {
            Equipped.skin = item.id
            let userImg = document.getElementById('inventory__user-img')
            //userImg!.setAttribute('src', `http://grechkogv.ru:3000/assets/${item.imageSrc}`)
            break
        }
    }
}

// удаление предмета нужной категории из хранилища айди надетых предметов 
function unEquipped(item: Item) {
    switch (item.category) {
        case 'helmet':
            Equipped.helmet = -1
            break
        case 'chest':
            Equipped.chest = -1
            break
        case 'leggins':
            Equipped.leggins = -1
            break
        case 'boots':
            Equipped.boots = -1
            break
        case 'weapon':
            Equipped.weapon = -1
            break
        case 'pet':
            Equipped.pet = -1
            break
        case 'skin': {
            Equipped.skin = -1
            let userImg = document.getElementById('inventory__user-img')
            //userImg!.setAttribute('src', `http://grechkogv.ru:3000/assets/`) базовый скин
            break
        }
    }
}

// функция надевания
function putOn(item: Item) {
    item.state = 'equipped'
    sendRequest('PATCH', server + `/items/${item.id}`, JSON.stringify({ state: item.state }))
    itemsInventory.set(item.id, item)
    equipItemHTML(item)
    equipped(item)

}
// функция снятия 
function takeOff(item: Item) {
    item.state = 'inventoried'
    sendRequest('PATCH', server + `/items/${item.id}`, JSON.stringify({ state: item.state }))
    itemsInventory.set(item.id, item)
    equipItemHTML(item, true)
}

// получение айди надетого предмета нужной категории
function idEquipped(item: Item): number {
    switch (item.category) {
        case 'helmet':
            return Equipped.helmet
        case 'chest':
            return Equipped.chest
        case 'leggins':
            return Equipped.leggins
        case 'boots':
            return Equipped.boots
        case 'weapon':
            return Equipped.weapon
        case 'pet':
            return Equipped.pet
        default:
            return Equipped.skin
    }

}
// отрисовка состояния предмета в HTML
function equipItemHTML(item: Item, unEquip: boolean = false) {
    let buf = <HTMLElement>document.getElementsByClassName(`idItem=${item.id}`)[0].childNodes[1]
    buf.classList.toggle('bg-success')
    buf.classList.toggle('bg-secondary')
    let btn = buf.getElementsByClassName('inventory__item-btn')[0]
    btn.innerHTML = unEquip ? 'Надеть' : 'Снять'
    btn.classList.toggle('btn-success')
    btn.classList.toggle('btn-primary')
}
