class Products {

    render() {
        let htmlCatalog = '';
        // eslint-disable-next-line @typescript-eslint/no-unused-vars
        let common = '#C8C8C8';
        // eslint-disable-next-line @typescript-eslint/no-unused-vars
        let rare = '#FFB74D';
        // eslint-disable-next-line @typescript-eslint/no-unused-vars
        let epic = '#F06292';
        // eslint-disable-next-line @typescript-eslint/no-unused-vars
        let legendary = '#26A69A';
        // eslint-disable-next-line @typescript-eslint/no-unused-vars
        let mythical = 'linear-gradient(#40E0D0, #FF8C00, #FF0080)';




        CATALOG_SHOP.forEach(({id, name, description, imageSrc, category, rarity}) => {

            if (category === 'armor' && id === 1) {
                htmlCatalog += `
                <div class="padding-right px-3 padding-left px-3">
                    <h3 class="display-4 text-center" id="clothes">
                        <a href="#top" style = "text-decoration: none">Одежда</a></h3>
                    <br>
                    <div class="row row-cols-lg-6 row-cols-md-4 row-cols-sm-3 row-cols-2 g-4">
                `;
            }
            htmlCatalog += `
                <div class="col">
                    <div class="card h-100">
                        <img src="http://grechkogv.ru:3000/assets/${imageSrc}" class="card-img-top" alt="...">
<!--                           подумать над заданием стиля-->
                        <div class="card-body" style="background: ${eval(`${rarity}`)}">
                            <h5 class="card-title">${name}</h5>
                                <p class="card-text">${description}</p>
                        </div>
                            <a href="" class="btn btn-primary">Buy</a>
                    </div>
                </div>
            `;
            if (category === 'armor' && id === 12) {
                htmlCatalog += `
                        </div>
                    </div>
                    <br>
                `;
            }

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


fetch('https://json.grechkogv.ru/items')
    .then(res => res.json())
    .then(body => {
        CATALOG_SHOP = body;
        productsPage.render();
    })
    .catch(error => {
        console.log(error);
    })

