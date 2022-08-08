class Router {
    private routes: Route[]
    private rootElem: HTMLDivElement

    constructor(routes: Route[]){
        try {
            if (!routes) {
                throw 'error: routes param is mandatory';
            }
            this.routes = routes
            this.rootElem = <HTMLDivElement>document.getElementById('body')
            this.init();
        } catch (e) {
            console.error(e);   
        }
    }

    private init() {
        var routes = this.routes;
        (function(scope, routes) { 
            window.addEventListener('hashchange', function (e) {
                scope.hasChanged();
            });
        })(this, routes);
        this.hasChanged();
    }


    public hasChanged (){
        if (window.location.hash.length > 0) {
            for (var i = 0, length = this.routes.length; i < length; i++) {
                var route = this.routes[i];
                if(route.isActiveRoute(window.location.hash.substr(1))) {
                    this.goToRoute(route.getHtmlName());
                }
            }
        } else {
            for (var i = 0, length = this.routes.length; i < length; i++) {
                var route = this.routes[i];
                if(route.getIsDefault()) {
                    this.goToRoute(route.getHtmlName());
                }
            }
        }
    }

    public goToRoute(getHtmlName: string) {
        (function(scope) { 
            var url = '../spa/views/' + getHtmlName,
                xhttp = new XMLHttpRequest();
            xhttp.onreadystatechange = function () {
                if (this.readyState === 4 && this.status === 200) {
                    scope.rootElem.innerHTML = this.responseText;
                }
            };
            xhttp.open('GET', url, true);
            xhttp.send();
        })(this);
    }

}
