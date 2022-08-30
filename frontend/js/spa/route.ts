class Route {
    private name: string
    private htmlName: string
    private isDefault: boolean
    private onLoadFn: Function
    private title: string

    constructor(name: string, htmlName: string, title: string, fn: Function, isDefault = false) {
        try {
            if (!name || !htmlName) {
                throw 'error: name and htmlName params are mandatories'
            }
            this.name = name
            this.htmlName = htmlName
            this.title = title
            this.onLoadFn = fn
            this.isDefault = isDefault
        } catch (e) {
            console.error(e)
        }
    }

    public getHtmlName(): string {
        return this.htmlName
    }

    public getName(): string {
        return this.name
    }

    public getIsDefault(): boolean {
        return this.isDefault
    }

    public evalFn() {
        let count = 0
        let isSuccessfull = true
        let timer = setInterval(() => {
            try {
                this.onLoadFn()
            } catch (error) {
                console.error(error)
                isSuccessfull = false
            }
            count++
            if (isSuccessfull || count > 10) {
                clearInterval(timer)
            }
            isSuccessfull = true
        }, 50)
    }

    public getTitle(): string {
        return this.title
    }
}
