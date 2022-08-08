class Route {
    private name: string
    private htmlName: string
    private isDefault: boolean

    constructor(name: string, htmlName: string, isDefault = false ) {
        try {
            if(!name || !htmlName) {
                throw 'error: name and htmlName params are mandatories';
            }
            this.name = name;
            this.htmlName = htmlName;
            this.isDefault = isDefault;
        } catch (e) {
            console.error(e);
        }
    }

    public isActiveRoute(hashedPath: string) : boolean {
        return hashedPath.replace('#', '') === this.name; 
    }
    
    public getHtmlName() : string {
        return this.htmlName
    }

    public getIsDefault(): boolean {
        return this.isDefault
    }
    
}
