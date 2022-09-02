;(function () {
    function init() {
        var router = new Router(
            new Map<string, Route>([
                ['todo', new Route('todo', 'todo.html', 'Список дел', onTodoLoad, true)],
                ['shop', new Route('shop', 'shop.html', 'Магазин', onShopLoad)],
                ['profile', new Route('profile', 'profile_page.html', 'Профиль', onProfileLoad)],
                ['inventory', new Route('inventory', 'inventory.html', 'Инвентарь', onInventoryLoad)],
            ]),
        )
    }
    init()
})()
