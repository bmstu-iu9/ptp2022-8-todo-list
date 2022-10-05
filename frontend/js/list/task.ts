class Task {
    private name: string
    private desc: string
    private createdOn: Date
    private dueDate: Date
    private shtirlitchHumoresque: string
    private labels: Label[]
    private id: number
    private status: Status

    constructor(
        inId: number,
        inName: string,
        inDesc?: string,
        inDueDate?: string,
        inLabels?: Label[],
        created: boolean = true
    ) {
        this.name = inName
        this.createdOn = new Date()
        this.desc = inDesc ? inDesc : ''
        if (inDueDate) this.dueDate = new Date(inDueDate)
        this.labels = inLabels ? Object.assign([], inLabels) : []
        this.id = inId
        this.status = 'active'
        if (created) {
            const tsk = {
                name: this.name,
                id: this.id + 1,
                description: this.desc,
                createdOn: this.createdOn,
                dueDate: this.dueDate,
                labels: this.labels,
                status: this.status,
            }
            sendRequest('POST', server + '/tasks', JSON.stringify(tsk))
        }
    }

    private createHTMLBlock() {
        let labelStr = ''
        if (this.labels.length > 0) {
            this.labels.forEach((element) => {
                labelStr = labelStr.concat(LabelToHtml(element))
            })
        }
        let str: string
        str = `<li class="${this.labels.length !== 0 ? 'pb-0 ' : ''}list-group-item id=${this.id}" ${
            this.status === 'archived' ? 'style="background-color: #e9ecef"' : ''
        }> 
                    <div>
                    <input class="form-check-input me-2" type="checkbox" value="1" ${
                        this.status === 'completed' ? 'checked' : ''
                    } ${this.status === 'archived' ? 'disabled' : ''} id="${this.id}">
                    <a class = "name__openModal">
                    <strong ${this.status === 'completed' ? 'style="text-decoration: line-through"' : ''}>${
            this.name
        }</strong></a>

                    <a title="У задачи есть прикрепленный файл">
                        <svg xmlns="http://www.w3.org/2000/svg" height="18" width="18" viewBox="0 0 24 24" class="bi bi-paperclip">
                        <g xmlns="http://www.w3.org/2000/svg" fill="none" fill-rule="evenodd">
                            <path d="m0 0h24v24h-24z"></path>
                            <path d="m20.0291094 15.0279907-5.384726 5.2303888c-2.5877049 2.513536-6.71408829 2.4838066-9.26530792-.0667538-2.6116233-2.6109485-2.61217034-6.8446794-.00122186-9.4563027.00760974-.0076117.01523784-.015205.02288425-.0227799l8.06657363-7.99110563c1.7601202-1.7436532 4.6004898-1.73030402 6.344143.02981623.0091252.00921136.0182104.01846224.0272554.02775238 1.7500823 1.79751906 1.7306631 4.66777042-.0435807 6.44144506l-8.1308667 8.12825806c-.8479169.8476448-2.20023168.9147308-3.12787932.1551687l-.1337127-.1094846c-.8947528-.7326277-1.02618115-2.0518803-.29355343-2.9466331.03855837-.047091.0791516-.0924786.12166404-.1360332l5.46733261-5.60136864" stroke="#828a99" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.8"></path>
                        </g>
                        </svg>
                    </a>
                    ${
                        this.desc !== ''
                            ? `
                    <a title="У задачи есть описание">
                        <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" fill="currentColor" class="bi bi-chat" viewBox="0 0 16 16">
                        <path d="M2.678 11.894a1 1 0 0 1 .287.801 10.97 10.97 0 0 1-.398 2c1.395-.323 2.247-.697 2.634-.893a1 1 0 0 1 .71-.074A8.06 8.06 0 0 0 8 14c3.996 0 7-2.807 7-6 0-3.192-3.004-6-7-6S1 4.808 1 8c0 1.468.617 2.83 1.678 3.894zm-.493 3.905a21.682 21.682 0 0 1-.713.129c-.2.032-.352-.176-.273-.362a9.68 9.68 0 0 0 .244-.637l.003-.01c.248-.72.45-1.548.524-2.319C.743 11.37 0 9.76 0 8c0-3.866 3.582-7 8-7s8 3.134 8 7-3.582 7-8 7a9.06 9.06 0 0 1-2.347-.306c-.52.263-1.639.742-3.468 1.105z"></path>
                        </svg>
                    </a>
                    `
                            : ''
                    }
                    

                    ${
                        this.dueDate
                            ? `
                    <div class="todo__time d-inline">
                        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="${
                            this.dueDate && this.dueDate.getTime() >= new Date().getTime()
                                ? 'todo__non-urgent'
                                : 'todo__important'
                        } bi bi-clock-fill" viewBox="0 0 16 16">
                        <path d="M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0zM8 3.5a.5.5 0 0 0-1 0V9a.5.5 0 0 0 .252.434l3.5 2a.5.5 0 0 0 .496-.868L8 8.71V3.5z"></path>
                        </svg>
                        <span>${formatDate(this.dueDate)}</span>
                    </div>`
                            : ''
                    }
                    ${labelStr}

                    <div class="d-flex flex-row justify-content-end">
                        <a href="#!" class="todo__archive" title="Архивировать">
                        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-archive-fill" viewBox="0 0 16 16">
                        <path d="M12.643 15C13.979 15 15 13.845 15 12.5V5H1v7.5C1 13.845 2.021 15 3.357 15h9.286zM5.5 7h5a.5.5 0 0 1 0 1h-5a.5.5 0 0 1 0-1zM.8 1a.8.8 0 0 0-.8.8V3a.8.8 0 0 0 .8.8h14.4A.8.8 0 0 0 16 3V1.8a.8.8 0 0 0-.8-.8H.8z"/>
                        </svg>
                        </a>
                        <a href="#" class="todo__delete text-danger" title="Удалить">
                        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-trash3-fill" viewBox="0 0 16 16">
                            <path d="M11 1.5v1h3.5a.5.5 0 0 1 0 1h-.538l-.853 10.66A2 2 0 0 1 11.115 16h-6.23a2 2 0 0 1-1.994-1.84L2.038 3.5H1.5a.5.5 0 0 1 0-1H5v-1A1.5 1.5 0 0 1 6.5 0h3A1.5 1.5 0 0 1 11 1.5Zm-5 0v1h4v-1a.5.5 0 0 0-.5-.5h-3a.5.5 0 0 0-.5.5ZM4.5 5.029l.5 8.5a.5.5 0 1 0 .998-.06l-.5-8.5a.5.5 0 1 0-.998.06Zm6.53-.528a.5.5 0 0 0-.528.47l-.5 8.5a.5.5 0 0 0 .998.058l.5-8.5a.5.5 0 0 0-.47-.528ZM8 4.5a.5.5 0 0 0-.5.5v8.5a.5.5 0 0 0 1 0V5a.5.5 0 0 0-.5-.5Z"></path>
                        </svg>
                        </a>
                    </div>

                    </div>
                </li>`
        return str
    }

    public toHTMLBlock() {
        let str = this.createHTMLBlock()
        let buf = document.querySelector('.todolist')!.innerHTML
        document.querySelector('.todolist')!.innerHTML = str.concat(buf)
    }

    public updateHTML() {
        let str = this.createHTMLBlock()
        let old_li = document.getElementsByClassName('id='.concat(this.id.toString()))[0]
        old_li.outerHTML = str
        this.saveToLocalStorage()
    }

    public clearHTML() {
        let old_li = document.getElementsByClassName('id='.concat(this.id.toString()))[0]
        old_li.outerHTML = ''
    }

    public addLabel(lbl: Label) {
        this.labels.push(lbl)
        this.updateHTML()
        this.saveToLocalStorage()
    }

    public saveToLocalStorage() {
        localStorage.setItem(this.id.toString(), JSON.stringify(this))
    }

    public log() {
        console.log(JSON.stringify(this))
    }

    public getId(): number {
        return this.id
    }

    public getName(): string {
        return this.name
    }

    public getLabels(): Label[] {
        return this.labels
    }

    public getDueDate(): Date {
        return this.dueDate
    }

    public getDesc(): string {
        return this.desc
    }

    public getStatus(): Status {
        return this.status
    }

    public setStatus(s: Status, edited: boolean = true) {
        this.status = s
        if (edited) {
            this.updateHTML()
            sendRequest('PATCH', server + `/tasks/${this.id}`, JSON.stringify({ status: this.status }))
        }
    }

    public editTask(n: string, date: string | undefined, l: Label[], d: string) {
        this.name = n
        if (date) this.dueDate = new Date(date)
        else delete this.dueDate
        this.labels = l
        this.desc = d
        this.updateHTML()
        const tsk = {
            name: this.name,
            id: this.id,
            description: this.desc,
            createdOn: this.createdOn,
            dueDate: this.dueDate,
            labels: this.labels,
            status: this.status,
        }
        sendRequest('PUT', server + `/tasks/${this.id}`, JSON.stringify(tsk))
    }
}
