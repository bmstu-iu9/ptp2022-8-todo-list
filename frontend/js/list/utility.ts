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
