// каталог магазина
let catalogShop = []
// для примера возьмем первого
const user = 0
// баланс беру с сервера
let balance: number
// все предметы в приложении id -> предмет
const itemsShop = new Map<number, Item>()
// модальная форма просмотра описания предмета в магазине
let modalShop: any

function onShopLoad() {
    modalShop = new bootstrap.Modal(<HTMLFormElement>document.getElementById('shopModal'))

    // Загружаю баланс
    sendRequest('GET', server + '/users/' + user).then((data) => {
        balance = data.balance
        const balanceShop = <HTMLInputElement>document.getElementById('balance')
        balanceShop.innerText = 'Ваш баланс: ' + balance.toString() + ' коинов'
    })

    // получаю предметы с сервера
    sendRequest('GET', server + '/items').then((data) => {
        catalogShop = data
        renderShop()
    })

    // заполняю карту itemsShop предметами
    sendRequest('GET', server + '/items').then((data) => {
        for (let i = 0; i < data.length; i++) {
            const item: Item = data[i]
            itemsShop.set(item.id, item)
        }
    })
}

try {
    onShopLoad()
} catch (error) {
}
