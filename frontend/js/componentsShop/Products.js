class Products {

    render() {
        // eslint-disable-next-line no-undef
        CATALOG.forEach((element) => {
            console.log(element)
            })
    }
}

const productsPage = new Products();
productsPage.render();