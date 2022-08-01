class Products {

    render() {
        let htmlCatalog = '';
        // eslint-disable-next-line no-undef
        CATALOG_SHOP.forEach(({id, name, description, imageSrc, category}) => {

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
                        <img src="${imageSrc}" class="card-img-top" alt="...">
<!--                           подумать над заданием стиля-->
<!--                           подумать над добавлением заголовков -->
                        <div class="card-body" style="background: #C8C8C8">
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
productsPage.render();