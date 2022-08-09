;(function () {
    function init() {
        var router = new Router(
            new Map<string, Route>([
                ['todo', new Route('todo', 'todo.html', onTodoLoad, true)],
                ['shop', new Route('shop', 'shop.html', () => {})],
                ['profile', new Route('profile', 'profile_page.html', () => {})],
                ['inventory', new Route('inventory', 'inventory.html', onInventoryLoad)],
            ])
        )
    }
    init()
})()
