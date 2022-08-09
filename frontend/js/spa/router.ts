class Router {
    private routes: Map<string, Route>
    private rootElem: HTMLDivElement
    private curPage: string

    constructor(routes: Map<string, Route>) {
        try {
            if (!routes) {
                throw 'error: routes param is mandatory'
            }
            this.routes = routes
            this.rootElem = <HTMLDivElement>document.getElementById('body')
            this.init()
        } catch (e) {
            console.error(e)
        }
    }

    private init() {
        var routes = this.routes
        ;(function (scope: Router) {
            window.addEventListener('hashchange', function (e) {
                scope.hasChanged()
            })
        })(this)
        this.hasChanged()
    }

    public hasChanged() {
        if (window.location.hash.length > 0) {
            let name = window.location.hash.substr(1).replace('#', '')
            let route = this.routes.get(name)
            this.goToRoute(route?.getHtmlName()!)
            route?.evalFn()
            let href = document.querySelector(`[href="#${this.curPage}"]`)
            href?.parentElement!.classList.remove('menu__open')
            href = document.querySelector(`[href="#${route?.getName()}"]`)
            href?.parentElement!.classList.add('menu__open')
            this.curPage = route?.getName()!
        } else {
            let route = this.routes.get('todo')
            this.goToRoute(route?.getHtmlName()!)
            route?.evalFn()
            let href = document.querySelector(`[href="#${route?.getName()}"]`)
            href?.parentElement!.classList.add('menu__open')
            this.curPage = route?.getName()!
        }
    }

    public goToRoute(getHtmlName: string) {
        ;(function (scope: Router) {
            var url = '../spa/views/' + getHtmlName,
                xhttp = new XMLHttpRequest()
            xhttp.onreadystatechange = function () {
                if (this.readyState === 4 && this.status === 200) {
                    scope.rootElem.innerHTML = this.responseText
                }
            }
            xhttp.open('GET', url, true)
            xhttp.send()
        })(this)
    }
}
