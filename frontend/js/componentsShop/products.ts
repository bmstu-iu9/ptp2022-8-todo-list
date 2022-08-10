class Products {

    render() {
        let htmlCatalog = '';
        let numberBlock = 0;
        const namingBlock: string[] = ['Одежда', 'Аксессуары', 'Питомцы', 'Облик'];
        let previousCategory = 'helmet';

        CATALOG_SHOP.forEach(({id, name, description, imageSrc, category, rarity, state}) => {

            if (category != previousCategory && category != 'helmet') {
                if (category != 'chest' && category != 'leggins') {
                    htmlCatalog += `
                            </div>
                        </div>
                        <br>
                    `;
                }
                previousCategory = category;
            }
            if ((category === 'armor' || category === 'helmet') && numberBlock === 0 ||
                category === 'weapon' && numberBlock === 1 ||
                category === 'pet' && numberBlock === 2 ||
                category === 'skin' && numberBlock === 3) {

                htmlCatalog += `
                <div class="padding-right px-3 padding-left px-3">
                    <h3 class="display-4 text-center">
                        <a href="#top" style = "text-decoration: none">${namingBlock[numberBlock]}</a></h3>
                    <br>
                    <div class="row row-cols-lg-6 row-cols-md-4 row-cols-sm-3 row-cols-2 g-4">
                `;
                numberBlock += 1;
            }
            htmlCatalog += `
                <div class="col idItem=${id}">
                    <div class="card h-100">
                        <img src="http://grechkogv.ru:3000/assets/${imageSrc}" class="card-img-top" alt="...">
                        <div class="card-body" style="background: ${getRarityColor(rarity)}">
                            <h5 class="card-title">${name}</h5>
                                <p class="card-text">${description}</p>
                        </div>
                            ${buildingBuyButton(id, state)}

                        <!- модальное окно --->
                        
                    </div>
                </div>
            `;
        });

        const html = `
                ${htmlCatalog}
        `;

        const ROOT_PRODUCTS = document.getElementById('products');
        ROOT_PRODUCTS.innerHTML = html;
    }
}

const productsPage = new Products();

let CATALOG_SHOP = [];

function alreadyBought(state: string): boolean {
    return state === 'inventoried' || state === 'equipped';
}

function buildingBuyButton(id: number, state: string): string {
    if (alreadyBought(state) === true) {
        return `
            <button type="button" class="for__click btn btn-success  idItem=${id}" data-bs-toggle="modal"
                                data-bs-target="#selling${id}" id="buttonBuy${id}"
                                style="border-radius: 0 0 3px 3px;">
                                Куплено</button>
        `;
    }
    else {
        return `
            <button type="button" class="for__click btn btn-primary idItem=${id}" data-bs-toggle="modal"
                                data-bs-target="#selling${id}"
                                style="border-radius: 0 0 3px 3px;">
                                Купить</button>
        `;
    }
}

// Загружаю баланс

// для примера возьмем первого
const user = 0;

sendRequest('GET', server + '/users/' + user).then((data) => {
    const balance: number = data.balance;
    const balanceShop = <HTMLInputElement>document.getElementById('balance');
    balanceShop.innerText = 'Ваш баланс: ' + balance.toString() + ' коинов';

})

// при нажатии на кнопку купить в модальном окне получается покупка

const itemsCostShop = new Map<number, Item>();
sendRequest('GET', server + '/items').then((data) => {
    for (let i = 0; i < data.length; i++) {
        const item: Item = data[i];
        itemsCostShop.set(item.price, item);

    }
})

document.addEventListener('click', (e) => {
    const targetBuy = <HTMLElement>e.target;

    if (hasParentClass(targetBuy, 'buyButton')) {
        const costAndString: string = targetBuy.innerHTML;
        const cost: number = parseInt(costAndString.replace(/\D+/g,""));
        const itemBuy: Item = itemsCostShop.get(cost)!;

        // получаем баланс
        // получили предмет

        const balanceShopForBuy: HTMLInputElement = <HTMLInputElement>document.getElementById('balance');
        const balanceShopForBuyAndString: string = balanceShopForBuy.innerHTML;
        let balanceBuy: number = parseInt(balanceShopForBuyAndString.replace(/\D+/g,""));

        // меняем состояние
        if (balanceBuy >= cost && itemBuy.state == 'store') {
            const footerBuy = <HTMLInputElement>document.getElementById('shopModalFooter');
            balanceBuy -= cost;
            itemBuy.state = 'inventoried';
            sendRequest('PATCH', server + '/users/' + user, JSON.stringify({ balance: balanceBuy }))
            sendRequest('PATCH', server + `/items/${itemBuy.id}`, JSON.stringify({ state: itemBuy.state }))
            footerBuy.innerHTML = `
                <button type="button" class="buyButton btn btn-success btn-lg disabled">
                    Предмет куплен
                </button>
                `;

            // меняю кнопку в самом магазине

            const buyFullCardInShop = <HTMLElement>document.getElementsByClassName('idItem=' + itemBuy.id)[0];
            const buyCardInShop = <HTMLElement>buyFullCardInShop.getElementsByClassName('card')[0];
            const buyButtonInShop = <HTMLElement>buyCardInShop.getElementsByClassName('for__click')[0];

            buyCardInShop.removeChild(buyButtonInShop);
            buyCardInShop.innerHTML += `
                <button type="button" class="for__click btn btn-success  idItem=${itemBuy.id}" data-bs-toggle="modal"
                    data-bs-target="#selling${itemBuy.id}" id="buttonBuy${itemBuy.id}"
                    style="border-radius: 0 0 3px 3px;">
                    Куплено</button>
            `;
        }
        else if (balanceBuy >= cost && (itemBuy.state == 'inventoried' || itemBuy.state == 'equipped')) {
            const footerBuy = <HTMLInputElement>document.getElementById('shopModalFooter');
            if (<HTMLInputElement>document.getElementById('duplicateBuyShop') == undefined) {
                footerBuy.innerHTML += `
                <p id="duplicateBuyShop" class="text-warning text-center">
                   Предмет уже куплен!
                </p>
                `;
            }
        }
        else {
            const footerBuy = <HTMLInputElement>document.getElementById('shopModalFooter');
            if (<HTMLInputElement>document.getElementById('duplicateBuyShop') == undefined) {
                footerBuy.innerHTML += `
                <p id="duplicateBuyShop" class="text-danger">
                   Недостаточно денег для покупки!
                </p>
                `;
            }
        }

    }
})

// Получаю предметы пользователя (пока просто предметы)

fetch('https://json.grechkogv.ru/items')
    .then(res => res.json())
    .then(body => {
        CATALOG_SHOP = body;
        productsPage.render();
    })
    .catch(error => {
        console.log(error);
    })
