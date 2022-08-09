class Route {
    private name: string
    private htmlName: string
    private isDefault: boolean
    private fn: Function

    constructor(name: string, htmlName: string, fn: Function, isDefault = false) {
        try {
            if (!name || !htmlName) {
                throw 'error: name and htmlName params are mandatories'
            }
            this.name = name
            this.htmlName = htmlName
            this.fn = fn
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
        let timer = setInterval(() => {
            this.fn()
            count++
            if (document.getElementById('modal') || document.getElementById('inventoryModal') || count > 10) {
                clearInterval(timer)
            }
        }, 100)
    }
}
