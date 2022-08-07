type Item = {
    id: number;
    name: string;
    imageSrc: string;
    description: string;
    price: number;
    category: Category;
    rarity: Rarity;
    state: ItemState;
}

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
const itemsInventory = new Map<number, Item>();

type Rarity = 'common' | 'rare' | 'epic' | 'legendary';

function getRarityColor(rarity: Rarity): string {
    switch (rarity) {
        case 'common':
            return '#C8C8C8';
        case 'rare':
            return '#0b9ccf';
        case 'epic':
            return '#d461f7';
        case 'legendary':
            return '#ff9c00';
        default:
            return 'linear-gradient(#40E0D0, #91e047, #fff456, #fff456, #ffa856, #e64f4f)';
    }
}

type ItemState = 'store' | 'equipped' | 'inventoried';
type Category = 'helmet' | 'chest' | 'leggins' | 'boots' | 'weapon' | 'pet' | 'skin';

// Создание динамического модального окна

// модальная форма просмотра описания карты

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
const modalInventory = new bootstrap.Modal(<HTMLFormElement>document.getElementById('shopModal'));

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
sendRequest('GET', server + '/items').then((data) => {
    for (let i = 0; i < data.length; i++) {
        const item: Item = data[i];
        itemsInventory.set(item.id, item);

    }
})

function emptyDescription(description: string): string {
     return description !== '' ? description : `Описание отсутствует`;
}

document.addEventListener('click', (e) => {
    const target = <HTMLElement>e.target;
    // получение из хранилища предмета, на карту которого кликнули
    const regexp = /idItem=\d+/;
    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-ignore
    const strId = findID(target, regexp);
    const id = parseInt(strId?.substring(7));
    // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
    const item = itemsInventory.get(id)!;
    console.log(itemsInventory);
    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-ignore
    if (hasParentClass(target, 'for__click')) {
        // просмотр описания карты

        const windowForm = <HTMLElement>document.getElementsByClassName('shop__form')[0];
        const buf = (<HTMLElement>document.getElementsByClassName(`idItem=${id}`)[0]).getBoundingClientRect();

        windowForm.style.top = `${buf.y}px`;
        windowForm.style.left = `calc(${buf.x}px - 1vw - 1.4rem)`;
        windowForm.style.width = `calc(${buf.right - buf.x}px + 2vw + 2.8rem)`;
        windowForm.style.borderColor = `${getRarityColor('common')}`;
        modalInventory.show();

        const titleModal = <HTMLInputElement>document.getElementById('ShopModalTitle');
        // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
        titleModal.innerHTML = item.name!

        const description = <HTMLInputElement>document.getElementById('shopModalBody');
        description.innerHTML = `
            <img src="https://wg.grechkogv.ru/assets/${item.imageSrc}"
            class="px-3 pb-3 img-fluid" alt="${item.imageSrc}">
            <div class="card-body" style="background: ${getRarityColor(item.rarity)}">
            <p class="card-text">${emptyDescription(item.description)}</p>
            `;

        const footer = <HTMLInputElement>document.getElementById('shopModalFooter');
        footer.innerHTML = `
            <button type="button" class="btn btn-primary btn-lg">
                Купить за ${item.price} todoкоинов?
            </button>
        `;

    }
});