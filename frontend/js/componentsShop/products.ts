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
                        <div class="card-body" style="background: ${color(rarity)}">
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

function color(rarity: string): string {
    if (rarity === "common") return "#C8C8C8";
    else if (rarity === 'rare') return "#FFB74D";
    else if (rarity === 'epic') return "#26A69A";
    else if (rarity === 'legendary') return "#F06292";
    return "linear-gradient(#40E0D0, #FF8C00, #FF0080)";
}

function alreadyBought(state: string): boolean {
    return state === 'inventoried' || state === 'equipped';
}

function buildingBuyButton(id: number, state: string): string {
    if (alreadyBought(state) === true) {
        return `
            <button type="button" class="for__click btn btn-success idItem=${id}" data-bs-toggle="modal"
                                data-bs-target="#selling${id}" id="buttonBuy${id}">Куплено</button>
        `;
    }
    else {
        return `
            <button type="button" class="for__click btn btn-primary idItem=${id}" data-bs-toggle="modal"
                                data-bs-target="#selling${id}">Купить</button>
        `;
    }
}

// Загружаю баланс

// для примера возьмем первого
const user = 0;

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
sendRequest('GET', server + '/users/' + user).then((data) => {
    const balance: number = data.balance;
    const balanceShop = <HTMLInputElement>document.getElementById('balance');
    balanceShop.innerText = 'Ваш баланс: ' + balance.toString() + ' коинов';

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
