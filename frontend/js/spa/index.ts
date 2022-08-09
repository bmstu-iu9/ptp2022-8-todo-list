
(function () {
    function init() {
        var router = new Router([
            new Route('todo', 'todo.html',onTodoLoad, true),            
            new Route('shop', 'shop.html', () => {}),
            new Route('profile', 'profile_page.html', () => {}),
            new Route('inventory', 'inventory.html', onInventoryLoad)
        ]);
    }
    init();
}());