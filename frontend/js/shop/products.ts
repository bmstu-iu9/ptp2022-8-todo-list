function renderShop() {
    let htmlCatalog = ''
    let numberBlock = 0
    const namingBlock: string[] = ['Одежда', 'Аксессуары', 'Питомцы', 'Облик']
    let previousCategory = 'helmet'

    catalogShop.forEach((item: Item) => {
        if (item.category != previousCategory && item.category != 'helmet') {
            if (item.category != 'chest' && item.category != 'leggins' && item.category != 'boots') {
                htmlCatalog += `
                            </div>
                        </div>
                        <br>
                    `
            }
            previousCategory = item.category
        }
        if (
            (item.category === 'helmet' && numberBlock === 0) ||
            (item.category === 'weapon' && numberBlock === 1) ||
            (item.category === 'pet' && numberBlock === 2) ||
            (item.category === 'skin' && numberBlock === 3)
        ) {
            htmlCatalog += `
                <div class="padding-right px-3 padding-left px-3"  id="${item.category}">
                    <h3 class="display-4 text-center">
                        <a href="#" style = "text-decoration: none" class="">
                            ${namingBlock[numberBlock]}</a></h3>
                    <br>
                    <div class="row row-cols-lg-6 row-cols-md-4 row-cols-sm-3 row-cols-2 g-4">
                `
            numberBlock += 1
        }
        htmlCatalog += `
                <div class="col idItem=${item.id}">
                    <div class="card h-100">
                        <img src="http://grechkogv.ru:3000/assets/${item.imageSrc}" class="card-img-top" alt="...">
                        <div class="card-body" style="background: ${getRarityColor(item.rarity)}">
                            <h5 class="card-title">${item.name}</h5>
                                <p class="card-text">${item.description}</p>
                        </div>
                            ${buildingBuyButton(item.id, item.state)}
                        
                    </div>
                </div>
            `
    })
    const rootProducts = document.getElementById('products')
    rootProducts.innerHTML = htmlCatalog
}

function alreadyBought(state: string): boolean {
    return state === 'inventoried' || state === 'equipped'
}

function buildingBuyButton(id: number, state: string): string {
    if (alreadyBought(state) === true) {
        return `
            <button type="button" class="for__click btn btn-success  idItem=${id}"
                                data-bs-target="#selling${id}" id="buttonBuy${id}"
                                style="border-radius: 0 0 3px 3px;">
                                Куплено</button>
        `
    } else {
        return `
            <button type="button" class="for__click btn btn-primary idItem=${id}"
                                style="border-radius: 0 0 3px 3px;">
                                Купить</button>
        `
    }
}

// при нажатии на кнопку купить в модальном окне получается покупка

document.addEventListener('click', (e) => {
    const targetBuy = <HTMLElement>e.target
    if (hasParentClass(targetBuy, 'buyButton')) {
        // получаем предмет
        const elementId: number = parseInt(targetBuy.id.replace(/\D+/g, ''))
        const itemBuy: Item = itemsShop.get(elementId)!

        // меняем состояние
        if (balance >= itemBuy.price) {
            if (itemBuy.state == 'store') {
                const footerBuy = <HTMLInputElement>document.getElementById('shopModalFooter')
                balance -= itemBuy.price
                itemBuy.state = 'inventoried'
                sendRequest('PATCH', server + '/users/' + user, JSON.stringify({ balance: balance }))
                sendRequest('PATCH', server + `/items/${itemBuy.id}`, JSON.stringify({ state: itemBuy.state }))
                footerBuy.innerHTML = `
                <button type="button" class="buyButton btn btn-success btn-lg disabled">
                    Предмет куплен
                </button>
                `

                // меняю кнопку в самом магазине
                const buyFullCardInShop = <HTMLElement>document.getElementsByClassName('idItem=' + itemBuy.id)[0]
                const buyCardInShop = <HTMLElement>buyFullCardInShop.getElementsByClassName('card')[0]
                const buyButtonInShop = <HTMLElement>buyCardInShop.getElementsByClassName('for__click')[0]
                buyCardInShop.removeChild(buyButtonInShop)
                buyCardInShop.innerHTML += `
                <button type="button" class="for__click btn btn-success  idItem=${itemBuy.id}" data-bs-toggle="modal"
                    data-bs-target="#selling${itemBuy.id}" id="buttonBuy${itemBuy.id}"
                    style="border-radius: 0 0 3px 3px;">
                    Куплено</button>
                `

                // меняю баланс в самом магазине
                const balanceBuyForShop = <HTMLInputElement>document.getElementById('balance')
                balanceBuyForShop.innerHTML = `
                <a id="balance" class="nav-link fw-bold py-1 px-0 mx-md-3 mx-auto my-auto disabled">
                            Ваш баланс: ${balance} коинов</a>
                `
            } else if (itemBuy.state == 'inventoried' || itemBuy.state == 'equipped') {
                const footerBuy = <HTMLInputElement>document.getElementById('shopModalFooter')
                if (<HTMLInputElement>document.getElementById('duplicateBuyShop') == undefined) {
                    footerBuy.innerHTML += `
                    <p id="duplicateBuyShop" class="text-warning text-center">
                        Предмет уже куплен!
                    </p>
                    `
                }
            }
        } else {
            const footerBuy = <HTMLInputElement>document.getElementById('shopModalFooter')
            if (<HTMLInputElement>document.getElementById('duplicateBuyShop') == undefined) {
                footerBuy.innerHTML += `
                    <p id="duplicateBuyShop" class="text-danger">
                        Недостаточно денег для покупки!
                    </p>
                `
            }
        }
    }
})
