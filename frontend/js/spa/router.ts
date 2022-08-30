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
            let route = this.routes.get(name)!
            this.goToRoute(route)
            route?.evalFn()
            let href = document.querySelector(`[href="#${this.curPage}"]`)
            href?.parentElement!.classList.remove('menu__open')
            href = document.querySelector(`[href="#${route?.getName()}"]`)
            href?.parentElement!.classList.add('menu__open')
            this.curPage = route?.getName()!
        } else {
            let route = this.routes.get('todo')!
            this.goToRoute(route)
            route?.evalFn()
            let href = document.querySelector(`[href="#${route?.getName()}"]`)
            href?.parentElement!.classList.add('menu__open')
            this.curPage = route?.getName()!
        }
    }

    public goToRoute(route: Route) {
        ;(function (scope: Router) {
            var url = route ? '../spa/views/' + route.getHtmlName() : '../spa/view/undefined',
                xhttp = new XMLHttpRequest()
            xhttp.onreadystatechange = function () {
                if (this.readyState === 4 && this.status === 200) {
                    scope.rootElem.innerHTML = this.responseText
                    document.title = route.getName()
                } else {
                    scope.rootElem.innerHTML = `<div class="body text-center fs-1 fw-bold" style="margin: 0; position: absolute; top: 35%;">
                                                        Ошибка 404<br>
                                                        <p class="fs-3 fw-normal">Страница не найдена</p>
                                                </div>`
                }
            }
            xhttp.open('GET', url, true)
            xhttp.send()
        })(this)
    }
}
