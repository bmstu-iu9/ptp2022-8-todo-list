'use strict';

(function () {
    function init() {
        var router = new Router([
            new Route('todo', 'todo.html', true),            
            new Route('shop', 'shop.html'),
            new Route('profile', 'profile_page.html')
        ]);
    }
    init();
}());