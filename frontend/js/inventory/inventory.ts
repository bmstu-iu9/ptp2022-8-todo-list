const modalInventory = new bootstrap.Modal(<HTMLFormElement>document.getElementById('inventoryModal'))
var itemsInventory = new Map<number, Item>()
let Equipped = {
    helmet: -1,
    chest: -1,
    leggins: -1,
    boots: -1,
    weapon: -1,
    pet: -1,
    skin: -1
}

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
        let userImg = document.getElementById("inventory__user-img")
        let id = Equipped.skin
        let item = itemsInventory.get(id)!
        userImg!.setAttribute('src', `http://grechkogv.ru:3000/assets/${item.imageSrc}`)
    }
})

document.addEventListener('click', (e) => {
    const target = <HTMLElement>e.target

    if (hasParentClass(target, 'inventory__box-title')) {
        let regexp = /idItem=\d+/
        let strId = findID(target, regexp)
        let id = parseInt(strId!.substring(7))
        let item = itemsInventory.get(id)!
        modalInventory.show()
        let titleModal = <HTMLInputElement>document.getElementById('inventoryModalTitle')
        titleModal.innerHTML = item.name!
        let description = <HTMLInputElement>document.getElementById('inventoryModalBoby')
        description.innerHTML = item.description! !== '' ? item.description! : 'Описание отсутствует'
    } else if (hasParentClass(target, 'inventory__item-btn')) {
        let regexp = /idItem=\d+/
        let strId = findID(target, regexp)!
        let id = parseInt(strId!.substring(7))
        let item = itemsInventory.get(id)!
        if (item.state === 'equipped') {
            unEquip(item)
            unEquipped(item)
        } else {
            let idE = idEquipped(item)
            if (idE !== -1) {
                let itemEquipped = itemsInventory.get(idE)!
                unEquip(itemEquipped)
            }
            equip(item)
        }

    }
})

function createInventoryHTMLBlock(item: Item) {

    let str: string
    str = `<div class="col my-2 idItem=${item.id}">
        <div class="card inventory__item ${item.state === 'equipped' ? 'bg-success' : 'bg-secondary'} p-2 h-100">
                <img src="http://grechkogv.ru:3000/assets/${item.imageSrc}"
                    class="card-img-top bg-light inventory__box-img p-2">
            <div class="card-body inventory__box-title btn text-dark" title="Посмотреть описание"
                style="background: ${color(item.rarity)}">
                <p class="card-title lh-1 align-middle m-0 p-0">${item.name}</p>
            </div>
            <div class="card-footer p-0">
            ${item.state === 'equipped' ?
            `<a class="inventory__item-btn btn btn-success w-100 h-100">Снять</a>`
            : `<a class="inventory__item-btn btn btn-primary w-100 h-100">Надеть</a>`}
            </div>
        </div>
    </div>`
    return str
}

function toInventoryHTMLBlock(item: Item) {
    let str = createInventoryHTMLBlock(item)
    let buf = document.querySelector('.inventory__box__items')!.innerHTML
    document.querySelector('.inventory__box__items')!.innerHTML = buf.concat(str)
}

function equipped(item: Item) {
    if (item.category === 'helmet') Equipped.helmet = item.id
    else if (item.category === 'chest') Equipped.chest = item.id
    else if (item.category === 'leggins') Equipped.leggins = item.id
    else if (item.category === 'boots') Equipped.boots = item.id
    else if (item.category === 'weapon') Equipped.weapon = item.id
    else if (item.category === 'pet') Equipped.pet = item.id
    else {
        Equipped.skin = item.id
        let userImg = document.getElementById("inventory__user-img")
        //userImg!.setAttribute('src', `http://grechkogv.ru:3000/assets/${item.imageSrc}`)
    }
}

function unEquipped(item: Item) {
    if (item.category === 'helmet') Equipped.helmet = -1
    else if (item.category === 'chest') Equipped.chest = -1
    else if (item.category === 'leggins') Equipped.leggins = -1
    else if (item.category === 'boots') Equipped.boots = -1
    else if (item.category === 'weapon') Equipped.weapon = -1
    else if (item.category === 'pet') Equipped.pet = -1
    else {
        Equipped.skin = -1
        let userImg = document.getElementById("inventory__user-img")
        //userImg!.setAttribute('src', `http://grechkogv.ru:3000/assets/`)
    }
}

function equip(item: Item) {
    item.state = 'equipped'
    sendRequest('PATCH', server + `/items/${item.id}`, JSON.stringify({ state: item.state }))
    itemsInventory.set(item.id, item)
    equipHTML(item)
    equipped(item)
}

function unEquip(item: Item) {
    item.state = 'inventoried'
    sendRequest('PATCH', server + `/items/${item.id}`, JSON.stringify({ state: item.state }))
    itemsInventory.set(item.id, item)
    equipHTML(item, true)
}

function idEquipped(item: Item): number {
    if (item.category === 'helmet') return Equipped.helmet
    else if (item.category === 'chest') return Equipped.chest
    else if (item.category === 'leggins') return Equipped.leggins
    else if (item.category === 'boots') return Equipped.boots
    else if (item.category === 'weapon') return Equipped.weapon
    else if (item.category === 'pet') return Equipped.pet
    else return Equipped.skin
}

function equipHTML(item: Item, unEquip: boolean = false) {
    let buf = <HTMLElement>document.getElementsByClassName(`idItem=${item.id}`)[0].childNodes[1]
    buf.classList.toggle('bg-success')
    buf.classList.toggle('bg-secondary')
    let btn = buf.getElementsByClassName('inventory__item-btn')[0]
    if (unEquip) btn.innerHTML = 'Надеть'
    else btn.innerHTML = 'Снять'
    btn.classList.toggle('btn-success')
    btn.classList.toggle('btn-primary')
}

