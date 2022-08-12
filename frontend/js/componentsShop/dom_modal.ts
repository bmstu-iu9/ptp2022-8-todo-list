function emptyDescription(description: string): string {
    return description !== '' ? description : `Описание отсутствует`;
}

document.addEventListener('click', (e) => {
    const target = <HTMLElement>e.target;
    // получение из хранилища предмета, на карту которого кликнули
    const regexp = /idItem=\d+/;
    const strId = findID(target, regexp);
    const id = parseInt(strId?.substring(7));
    const item = itemsShop.get(id)!;

    if (hasParentClass(target, 'for__click')) {
        // просмотр описания карты
        const windowForm = <HTMLElement>document.getElementsByClassName('shop__form')[0];
        const buf = (<HTMLElement>document.getElementsByClassName(`idItem=${id}`)[0]).getBoundingClientRect();

        // делать по центру одинаково
        windowForm.style.width = `calc(${buf.right - buf.x}*1.2px + 3vw + 3.8rem)`;
        modalShop.show();

        const titleModal = <HTMLInputElement>document.getElementById('ShopModalTitle');
        titleModal.innerHTML = item.name!;
        const description = <HTMLInputElement>document.getElementById('shopModalBody');
        description.innerHTML = `
            <img src="https://wg.grechkogv.ru/assets/${item.imageSrc}"
            class="px-3 pb-3 img-fluid" alt="${item.imageSrc}">
            <div class="card-body" style="background: ${getRarityColor(item.rarity)}">
            <p class="card-text">${emptyDescription(item.description)}</p>
            `;
        const footer = <HTMLInputElement>document.getElementById('shopModalFooter');
        if (item.state == 'store') {
            footer.innerHTML = `
                <button id="buyButtonDuplicateModal${item.id}" type="button" class="buyButton btn btn-primary btn-lg">
                    Купить за ${item.price} todoкоинов?
                </button>
                `;
        } else if (item.state == 'inventoried' || item.state == 'equipped') {
            footer.innerHTML = `
               <button id="buyButtonDuplicateModal${item.id}" type="button" class="buyButton btn btn-success btn-lg disabled">
                   Предмет куплен
               </button>
            `;
        }
    }
});