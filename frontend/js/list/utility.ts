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
    color: string
    text: string
    id: number
}

function makeLabel(colorHex: string, textInput: string, classUl: string, n: number) {
    let lbl: Label = {
        id: lbl_id,
        color: colorHex.substring(1),
        text: textInput,
    }
    labels.set(lbl_id++, lbl)
    lbl.id++
    sendRequest('POST', server + '/labels', JSON.stringify(lbl))
    let ul = <HTMLUListElement>document.getElementsByClassName(classUl)[0]
    let lbl_text = `<li><a class="dropdown-item d-flex align-items-center gap-2 py-1" href="#" id="lbl-${lbl.id}">
                    <span style="background-color: #${lbl!.color}99;"
                    class="d-inline-block rounded-circle p-1"></span>
                    <label class="todo__category">${lbl.text}</label>
                </a></li>`
    ul.innerHTML = ul.innerHTML.concat(lbl_text)
    let container = document.getElementsByClassName('chosen__categories')[n]
    container.innerHTML =
        container.innerHTML +
        `<div class="category col-auto" 
                    style="background-color: #${lbl!.color}99;" id="lbl-${lbl?.id}">
                    <label>${lbl?.text}</label>
                    <button type="button" class="btn-close btn-close-lbl"></button>
                    </div>`
}

function openUl(form: string, show: string, searchId: string, list: string): boolean {
    let elem = <HTMLDivElement>document.getElementById(form)
    if (elem.classList.contains(show)) {
        elem.classList.remove(show)
        let ul = <HTMLUListElement>document.getElementsByClassName(list)[0]
        ul.innerHTML = ''
        let search = <HTMLInputElement>document.getElementById(searchId)
        search.value = ''
        return true
    } else {
        let ul = <HTMLUListElement>document.getElementsByClassName(list)[0]
        labels.forEach((lbl) => {
            let lbl_text = `<li><a class="dropdown-item d-flex align-items-center gap-2 py-1" href="#" id="lbl-${lbl.id}">
                                <span style="background-color: #${lbl.color};"
                                class="d-inline-block rounded-circle p-1"></span>
                                <label class="todo__category">${lbl.text}</label>
                            </a></li>`
            ul.innerHTML = ul.innerHTML.concat(lbl_text)
            if (sortLabels.has(lbl.id) && show === 'show') {
                ;(<HTMLElement>document.getElementById(`lbl-${lbl!.id}`))!.style.backgroundColor = `#${lbl.color}99`
            }
        })
        elem.classList.toggle(show)
        return false
    }
}

function closeUl(dropbtn: string, form: string, list: string, searchId: string, show: string) {
    ;(<HTMLUListElement>document.getElementsByClassName(dropbtn)[0]).setAttribute('placeholder', 'Выбрать категорию')
    let elem = <HTMLDivElement>document.getElementById(form)
    elem.classList.remove(show)
    let ul = <HTMLUListElement>document.getElementsByClassName(list)[0]
    ul.innerHTML = ''
    let search = <HTMLInputElement>document.getElementById(searchId)
    search.value = ''
}

function addLabel(n: number, id: string) {
    let lbl = labels.get(parseInt(id.substring(4)) - 1)
    let container = document.getElementsByClassName('chosen__categories')[n]
    container.innerHTML =
        container.innerHTML +
        `<div class="category col-auto" 
                        style="background-color: #${lbl!.color}99;" id="lbl-${lbl?.id}">
                        <label>${lbl?.text}</label>
                        <button type="button" class="btn-close btn-close-lbl"></button>
                        </div>`
    LabelToHtml(<Label>lbl)
}

function search(classUl: string, inpt: string) {
    if (inpt === '') {
        let ul = <HTMLUListElement>document.getElementsByClassName(classUl)[0]
        ul.innerHTML = ''
        labels.forEach((lbl) => {
            let lbl_text = `<li><a class="dropdown-item d-flex align-items-center gap-2 py-1" href="#" id="lbl-${lbl.id}">
                                    <span style="background-color: #${lbl.color};"
                                        class="d-inline-block rounded-circle p-1"></span>
                                    <label class="todo__category">${lbl.text}</label>
                                </a></li>`
            ul.innerHTML = ul.innerHTML.concat(lbl_text)

            if (sortLabels.has(lbl.id)) {
                ;(<HTMLElement>document.getElementById(`lbl-${lbl!.id}`))!.style.backgroundColor = `#${lbl.color}99`
            }
        })
    } else {
        let ul = <HTMLUListElement>document.getElementsByClassName(classUl)[0]
        ul.innerHTML = ''
        labels.forEach((lbl) => {
            if (lbl.text.startsWith(inpt)) {
                let lbl_text = `<li><a class="dropdown-item d-flex align-items-center gap-2 py-1" href="#" id="lbl-${lbl.id}">
                                        <span style="background-color: #${lbl.color};"
                                            class="d-inline-block rounded-circle p-1"></span>
                                        <label class="todo__category">${lbl.text}</label>
                                    </a></li>`
                ul.innerHTML = ul.innerHTML.concat(lbl_text)

                if (sortLabels.has(lbl.id)) {
                    ;(<HTMLElement>document.getElementById(`lbl-${lbl!.id}`))!.style.backgroundColor = `#${lbl.color}99`
                }
            }
        })
    }
}

function LabelToHtml(lbl: Label): string {
    return `<label class="todo__category"
            style="background-color: #${lbl.color}99">
            ${lbl.text}</label>`
}

type Status = 'active' | 'completed' | 'archived'
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
