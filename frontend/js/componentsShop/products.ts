class Products {

    render() {
        let htmlCatalog = '';
        let numberBlock = 0;
        const namingBlock: string[] = ['Одежда', 'Аксессуары', 'Питомцы', 'Облик'];
        let previousCategory = 'helmet';

        CATALOG_SHOP.forEach(({name, description, imageSrc, category, rarity}) => {

            if (category != previousCategory && category != 'helmet') {
                if (category != 'armor') {
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
                    <h3 class="display-4 text-center" id="${category}">
                        <a href="#top" style = "text-decoration: none">${namingBlock[numberBlock]}</a></h3>
                    <br>
                    <div class="row row-cols-lg-6 row-cols-md-4 row-cols-sm-3 row-cols-2 g-4">
                `;
                numberBlock += 1;
            }
            htmlCatalog += `
                <div class="col">
                    <div class="card h-100">
                        <img src="http://grechkogv.ru:3000/assets/${imageSrc}" class="card-img-top" alt="...">
                        <div class="card-body" style="background: ${color(rarity)}"">
                            <h5 class="card-title">${name}</h5>
                                <p class="card-text">${description}</p>
                        </div>
                            <a href="" class="btn btn-primary">Buy</a>
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

function color(rarity): string {
    if (rarity === "common") return "#C8C8C8"
    else if (rarity === 'rare') return "#2bfff4"
    else if (rarity === 'epic') return "#f04dff"
    return "#linear-gradient(#40E0D0, #91e047, #fff456, #fff456, #ffa856, #e64f4f)"
}

fetch('https://json.grechkogv.ru/items')
    .then(res => res.json())
    .then(body => {
        CATALOG_SHOP = body;
        productsPage.render();
    })
    .catch(error => {
        console.log(error);
    })
