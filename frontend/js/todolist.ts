let glob_id = 0
declare var bootstrap: any

type Color = {
    red: number
    green: number
    blue: number
}

function createColor(r: number, g: number, b: number): Color {
    let c: Color = {
        red: r,
        green: g,
        blue: b,
    }
    return c
}

type Label = {
    color: Color
    text: string
    id: number
}

function LabelToHtml(lbl: Label): string {
    return `<label class="todo__category"
            style="background-color: rgba(${lbl.color.red},${lbl.color.green},${lbl.color.blue},0.5)">
            ${lbl.text}</label>`
}

type Status = 'active' | 'completed' | 'deleted'
type displayStatus = 'visible' | 'hidden'

function formatDate(d: Date): string {
    if (!d) return ''
    let WeekMap = new Map<number, string>([
        [0, 'Вс'],
        [1, 'Пн'],
        [2, 'Вт'],
        [3, 'Ср'],
        [4, 'Чт'],
        [5, 'Пт'],
        [6, 'Сб'],
    ])
    let MonthMap = new Map<number, string>([
        [0, 'Января'],
        [1, 'Февраля'],
        [2, 'Марта'],
        [3, 'Апреля'],
        [4, 'Мая'],
        [5, 'Июня'],
        [6, 'Июля'],
        [7, 'Августа'],
        [8, 'Сентября'],
        [9, 'Октября'],
        [10, 'Ноября'],
        [11, 'Декабря'],
    ])
    let weekday = WeekMap.get(d.getDay())
    let day = d.getDate()
    let month = MonthMap.get(d.getMonth())
    let hour = d.getHours()
    let minutes = d.getMinutes()
    let time = ''
    if (hour !== 0 && minutes !== 0) {
        time = hour + ':' + (minutes < 10 ? '0' + minutes : minutes)
    }

    return weekday + ', ' + day + ' ' + month + ' ' + time
}

function hasParentClass(elem: HTMLElement, className: string) {
    if (elem.classList.contains(className)) return true
    return elem.parentElement && hasParentClass(elem.parentElement, className)
}

function findID(elem: HTMLElement, regexp: RegExp): string | undefined {
    if (regexp.test(elem.classList.value)) return elem.classList.value.match(regexp)![0].toString()
    if (elem.parentElement) return findID(elem.parentElement, regexp)
    return 'failure'
}

class Task {
    private name: string
    private desc: string
    private createdOn: Date
    private dueDate: Date
    private shtirlitchHumoresque: string
    private labels: Label[]
    private id: number
    private status: Status

    constructor(inName: string, inDesc?: string, inDueDate?: string, inLabels?: Label[]) {
        this.name = inName
        this.createdOn = new Date()
        if (inDesc) this.desc = inDesc
        if (inDueDate) this.dueDate = new Date(inDueDate)
        if (inLabels) this.labels = Object.assign([], inLabels)
        else this.labels = []
        this.id = glob_id
        glob_id++
        this.status = 'active'
    }

    public setStatus(s: Status) {
        this.status = s
        this.updateHTML()
        this.saveToLocalStorage()
    }

    private createHTMLBlock() {
        let labelStr = ''
        if (this.labels.length > 0) {
            this.labels.forEach((element) => {
                labelStr = labelStr.concat(LabelToHtml(element))
            })
        }
        let str: string
        str = `<li class="list-group-item id=${this.id}">
                    <div>
                    <input class="form-check-input me-2" type="checkbox" value="1" ${
                        this.status === 'completed' ? 'checked' : ''
                    } id="${this.id}">
                    <strong ${this.status === 'completed' ? 'style="text-decoration: line-through"' : ''}>${
            this.name
        }</strong>

                    <a href="#" title="">
                        <svg xmlns="http://www.w3.org/2000/svg" height="18" width="18" viewBox="0 0 24 24" class="bi bi-paperclip">
                        <g xmlns="http://www.w3.org/2000/svg" fill="none" fill-rule="evenodd">
                            <path d="m0 0h24v24h-24z"></path>
                            <path d="m20.0291094 15.0279907-5.384726 5.2303888c-2.5877049 2.513536-6.71408829 2.4838066-9.26530792-.0667538-2.6116233-2.6109485-2.61217034-6.8446794-.00122186-9.4563027.00760974-.0076117.01523784-.015205.02288425-.0227799l8.06657363-7.99110563c1.7601202-1.7436532 4.6004898-1.73030402 6.344143.02981623.0091252.00921136.0182104.01846224.0272554.02775238 1.7500823 1.79751906 1.7306631 4.66777042-.0435807 6.44144506l-8.1308667 8.12825806c-.8479169.8476448-2.20023168.9147308-3.12787932.1551687l-.1337127-.1094846c-.8947528-.7326277-1.02618115-2.0518803-.29355343-2.9466331.03855837-.047091.0791516-.0924786.12166404-.1360332l5.46733261-5.60136864" stroke="#828a99" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.8"></path>
                        </g>
                        </svg>
                    </a>

                    <a href="#" title="">
                        <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" fill="currentColor" class="bi bi-chat" viewBox="0 0 16 16">
                        <path d="M2.678 11.894a1 1 0 0 1 .287.801 10.97 10.97 0 0 1-.398 2c1.395-.323 2.247-.697 2.634-.893a1 1 0 0 1 .71-.074A8.06 8.06 0 0 0 8 14c3.996 0 7-2.807 7-6 0-3.192-3.004-6-7-6S1 4.808 1 8c0 1.468.617 2.83 1.678 3.894zm-.493 3.905a21.682 21.682 0 0 1-.713.129c-.2.032-.352-.176-.273-.362a9.68 9.68 0 0 0 .244-.637l.003-.01c.248-.72.45-1.548.524-2.319C.743 11.37 0 9.76 0 8c0-3.866 3.582-7 8-7s8 3.134 8 7-3.582 7-8 7a9.06 9.06 0 0 1-2.347-.306c-.52.263-1.639.742-3.468 1.105z"></path>
                        </svg>
                    </a>

                    ${this.dueDate ? `
                    <div class="todo__time d-inline">
                        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="${
                            this.dueDate && this.dueDate.getTime() >= new Date().getTime()
                                ? 'todo__non_urgent'
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
                        <a href="#!" class="todo__edit text-info" title="Редактировать">
                        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-pencil-fill" viewBox="0 0 16 16">
                            <path d="M12.854.146a.5.5 0 0 0-.707 0L10.5 1.793 14.207 5.5l1.647-1.646a.5.5 0 0 0 0-.708l-3-3zm.646 6.061L9.793 2.5 3.293 9H3.5a.5.5 0 0 1 .5.5v.5h.5a.5.5 0 0 1 .5.5v.5h.5a.5.5 0 0 1 .5.5v.5h.5a.5.5 0 0 1 .5.5v.207l6.5-6.5zm-7.468 7.468A.5.5 0 0 1 6 13.5V13h-.5a.5.5 0 0 1-.5-.5V12h-.5a.5.5 0 0 1-.5-.5V11h-.5a.5.5 0 0 1-.5-.5V10h-.5a.499.499 0 0 1-.175-.032l-.179.178a.5.5 0 0 0-.11.168l-2 5a.5.5 0 0 0 .65.65l5-2a.5.5 0 0 0 .168-.11l.178-.178z"></path>
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
}

// Текущее хранилище задач
var tasks: Task[] = []
// Текущее хранилище лейблов
var labels: Label[] = []
// модальная форма ввода задачи
const modal = new bootstrap.Modal(<HTMLFormElement>document.getElementById('modal'))

// Лейблы по умолчанию
labels.push({color: createColor(0,116,15), text: "семья", id: 1})
labels.push({color: createColor(0, 73, 209), text: "работа", id: 2})
labels.push({color: createColor(101, 1, 168), text: "кот", id: 3})
labels.push({color: createColor(168, 54, 1), text: "позвонить гоше", id: 4})


// Общая обработка кликов по странице
document.addEventListener('click', (e) => {
    const target = <HTMLElement>e.target
    //console.log(target.classList)
    if (target.classList.contains('form-check-input')) { // нажатие чекбокса у задачи
        let chckBox = <HTMLInputElement>target
        let id = chckBox.id
        let li = <HTMLUListElement>document.getElementsByClassName('id='.concat(id))[0]
        if (chckBox.checked) {
            chckBox!.setAttribute('checked', '')
            tasks[id].setStatus('completed')
        } else {
            chckBox!.removeAttribute('checked')
            tasks[id].setStatus('active')
        }
    } else 
    if (target.classList.contains('add-btn')) { // Быстрое создание задачи
        let inptBox = <HTMLInputElement>document.getElementsByClassName('todo_text')[0]
        let name = inptBox!.value
        inptBox!.value = ''
        let newTask = new Task(name)
        tasks.push(newTask)
        newTask.toHTMLBlock()
    } else 
    if (target.classList.contains('add-extended')) { // Расширенное создание задачи 
        let inptBox = <HTMLInputElement>document.getElementById('name')
        let dateBox = <HTMLInputElement>document.getElementById('date')
        let timeBox = <HTMLInputElement>document.getElementById('time')
        let commentText = <HTMLTextAreaElement>document.getElementById('comment')
        let name = inptBox!.value
        let date = dateBox!.value + ' ' + timeBox!.value
        let comment = commentText.value
        let container = <HTMLDivElement>document.getElementsByClassName('chosen__categories')[0]
        let taskLabels: Label[] = []
        let lbls = container.getElementsByClassName('category')
        for (let i = 0; i < lbls.length; i++){
            taskLabels.push(labels[parseInt(lbls[i].id.substring(4)) - 1])
        }
        console.log(lbls)
        if (inptBox!.value !== '') {
            inptBox!.value = ''
            dateBox!.value = ''
            timeBox!.value = ''
            commentText.value = ''
            let newTask: Task
            if (date !== ' ') {
                newTask = new Task(name, comment, date, taskLabels)
            } else {
                newTask = new Task(name, comment, undefined, taskLabels)
            }
            tasks.push(newTask)
            newTask.toHTMLBlock()
        }
    } else 
    if (hasParentClass(target, 'todo__delete')) { // удаление задачи (TODO: понять, что с ней делать)
        let regexp = /id=\d+/
        let strId = findID(target, regexp)
        let id = parseInt(strId!.substring(3))
        tasks[id].setStatus('deleted')
        tasks[id].clearHTML()
    } else 
    if (target.classList.contains('btn-close')) { // Удаление лейбла в модалке
        target.parentElement!.outerHTML = ''
    } else 
    if (hasParentClass(target, 'dropdown-item') && hasParentClass(target, 'list-category')) { // добавление лейбла в контейнер
        let regexp = /id=lbl-\d+/
        let link = <HTMLElement>target.closest('.dropdown-item')
        let lbl = labels[parseInt(link.id.substring(4)) - 1]
        console.log(lbl)
        let container = document.getElementsByClassName('chosen__categories')[0]
        container.innerHTML = container.innerHTML + `<div class="category col-auto" 
                        style="background-color: rgba(${lbl.color.red + ', ' + lbl.color.green + ', ' + lbl.color.blue}, 0.5);" id="lbl-${lbl.id}">
                        <label>${lbl.text}</label>
                        <button type="button" class="btn-close"></button>
                        </div>`
         LabelToHtml(lbl)
    } else 
    if (hasParentClass(target, 'add-lbl')) {
        let textInput = <HTMLInputElement>document.getElementById('name__category')
        let colorInput = <HTMLInputElement>document.getElementById('color__category')
        let colorHex = colorInput!.value
        let lbl: Label = {
            color: createColor(parseInt(colorHex.substring(1, 3), 16), parseInt(colorHex.substring(3, 5), 16), parseInt(colorHex.substring(5, 7), 16)),
            text: textInput.value,
            id: labels.length + 1
        }
        labels.push(lbl)
        let ul = <HTMLUListElement>document.getElementsByClassName('list-category')[0]
        let lbl_text = `<li><a class="dropdown-item d-flex align-items-center gap-2 py-1" href="#" id="lbl-${lbl.id}">
                        <span style="background-color: rgba(${lbl.color.red + ', ' + lbl.color.green + ', ' + lbl.color.blue}, 0.5);"
                        class="d-inline-block rounded-circle p-1"></span>
                        <label class="todo__category">${lbl.text}</label>
                    </a></li>`
        ul.innerHTML = ul.innerHTML.concat(lbl_text)
    } else 
    if (target.classList.contains('dropbtn2')) { // отрисовка выбранных лейблов в списке
        let elem = <HTMLDivElement>document.getElementById('form__category')
        if (elem.classList.contains("show1")) {
            elem.classList.remove("show1")
            let ul = <HTMLUListElement>document.getElementsByClassName('list-category')[0]
            ul.innerHTML = ''
          } else {
            let ul = <HTMLUListElement>document.getElementsByClassName('list-category')[0]
            labels.forEach(lbl => {
                let lbl_text = `<li><a class="dropdown-item d-flex align-items-center gap-2 py-1" href="#" id="lbl-${lbl.id}">
                                <span style="background-color: rgba(${lbl.color.red + ', ' + lbl.color.green + ', ' + lbl.color.blue}, 0.5);"
                                class="d-inline-block rounded-circle p-1"></span>
                                <label class="todo__category">${lbl.text}</label>
                            </a></li>`
                ul.innerHTML = ul.innerHTML.concat(lbl_text)
            });
            elem.classList.toggle("show1")
          }
    }
})

// Обработка событий изменений

document.addEventListener('change', (e) => {
    let target = <HTMLElement>e.target
    if (target.classList.contains('stage')) {
        // код фильтра
    }
    if (target.id === 'name') {
        let inpt = <HTMLInputElement>target
        if (inpt.value === '') {
            inpt.classList.add('is-invalid')
            inpt.classList.remove('is-valid')
        } else {
            inpt.classList.add('is-valid')
            inpt.classList.remove('is-invalid')
        }
    }
})

// Создает список с двумя рабочими элементами
function generateExample(): void {
    let foo = new Task('first task', 'моя задача очень важна', '07-10-2022 22:22')
    foo.toHTMLBlock()
    tasks.push(foo)
    let foo1 = new Task('second task', 'моя задача очень важна', '07-30-2022 21:21')
    foo1.toHTMLBlock()
    tasks.push(foo1)
}

// Отдельные обработчики событий

const btn = <HTMLButtonElement>document.querySelector('#btn')

btn.addEventListener('click', function () {
    modal.show()
})

// Работа с локальным хранилищем

function saveAll() {
    tasks.forEach((task) => {
        localStorage.setItem(task.getId().toString(), JSON.stringify(task))
    })
}

function getFromStorage() {
    for (const task in localStorage) {
        if (Object.prototype.hasOwnProperty.call(localStorage, task)) {
            const element = localStorage[task]
            let buf = JSON.parse(element)
            let newTask = new Task(buf.name, buf.desc, buf.dueDate, buf.labels)
            tasks.push(newTask)
            newTask.toHTMLBlock()
        }
    }
}

