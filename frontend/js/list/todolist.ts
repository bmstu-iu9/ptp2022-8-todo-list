let glob_id = 0
let lbl_id = 0
declare var bootstrap: any

// Текущее хранилище задач
var tasks = new Map<number, Task>()
// Текущее хранилище лейблов
var labels = new Map<number, Label>()
// Хранилище айди лейблов для сортировки
var sortLabels = new Map<number, boolean>()
// модальная форма ввода задачи
var modal: any
// модальная форма редактирования задачи
var modalEditor: any

var btn: HTMLButtonElement

function onTodoLoad() {
    glob_id = 0
    lbl_id = 0
    modal = new bootstrap.Modal(<HTMLFormElement>document.getElementById('modal'))
    modalEditor = new bootstrap.Modal(<HTMLFormElement>document.getElementById('modal__editor'))
    btn = <HTMLButtonElement>document.querySelector('#btn')

    btn.addEventListener('click', function () {
        let colorInput = <HTMLInputElement>document.getElementById('color__category')
        colorInput.value = '#' + Math.random().toString(16).slice(-6)
    })

    // Получение лейблов с сервера
    sendRequest('GET', server + '/labels').then((data) => {
        for (let i = 0; i < data.length; i++) {
            labels.set(lbl_id++, data[i])
        }
    })

    // Получение задач с сервера
    sendRequest('GET', server + '/tasks').then((data) => {
        for (let i = 0; i < data.length; i++) {
            let buf = data[i]
            // Саше не понравилось, нужно переделать
            let newTask = new Task(buf.id, buf.name, buf.description, buf.dueDate, buf.labels, false)
            newTask.setStatus(buf.status, false)
            tasks.set(buf.id, newTask)
            if (newTask.getStatus() === 'active') {
                newTask.toHTMLBlock()
            }
            glob_id = buf.id + 1
        }
    })
}

try {
    onTodoLoad()
} catch (error) {}

// Общая обработка кликов по странице
document.addEventListener('click', (e) => {
    const target = <HTMLElement>e.target
    //console.log(target.classList)
    if (target.classList.contains('form-check-input')) {
        // нажатие чекбокса у задачи
        let chckBox = <HTMLInputElement>target
        let id = chckBox.id
        let li = <HTMLUListElement>document.getElementsByClassName('id='.concat(id))[0]
        if (chckBox.checked) {
            chckBox!.setAttribute('checked', '')
            tasks.get(parseInt(id))!.setStatus('completed')
            tasks.get(parseInt(id))!.clearHTML()
        } else {
            chckBox!.removeAttribute('checked')
            tasks.get(parseInt(id))!.setStatus('active')
            tasks.get(parseInt(id))!.clearHTML()
        }
    } else if (target.classList.contains('add-btn')) {
        // Быстрое создание задачи
        let inptBox = <HTMLInputElement>document.getElementsByClassName('todo__text')[0]
        let name = inptBox!.value
        if (name !== '') {
            inptBox!.value = ''
            let newTask = new Task(glob_id + 1, name)
            tasks.set(glob_id++, newTask)
            newTask.toHTMLBlock()
        }
    } else if (target.classList.contains('add-extended')) {
        // Расширенное создание задачи
        let inpt = <HTMLInputElement>document.getElementById('name')
        if (inpt.classList.contains('is-invalid')) inpt.classList.remove('is-invalid')
        else if (inpt.classList.contains('is-valid')) inpt.classList.remove('is-valid')
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
        for (let i = 0; i < lbls.length; i++) {
            taskLabels.push(labels.get(parseInt(lbls[i].id.substring(4)) - 1)!)
        }
        if (inptBox!.value !== '') {
            inptBox!.value = ''
            dateBox!.value = ''
            timeBox!.value = ''
            commentText.value = ''
            let container = document.getElementsByClassName('chosen__categories')[0]
            container.innerHTML = ''
            let newTask: Task
            if (date !== ' ') {
                newTask = new Task(glob_id + 1, name, comment, date, taskLabels)
            } else {
                newTask = new Task(glob_id + 1, name, comment, undefined, taskLabels)
            }
            tasks.set(glob_id++, newTask)
            newTask.toHTMLBlock()
        }
    } else if (hasParentClass(target, 'todo__archive')) {
        // архивирование задачи
        let regexp = /id=\d+/
        let strId = findID(target, regexp)
        let id = parseInt(strId!.substring(3))
        if (tasks.get(id)!.getStatus() === 'active') {
            tasks.get(id)!.setStatus('archived')
            tasks.get(id)!.clearHTML()
        } else {
            tasks.get(id)!.setStatus('active')
            tasks.get(id)!.clearHTML()
        }
    } else if (hasParentClass(target, 'todo__delete')) {
        // удаление задачи (TODO: понять, что с ней делать)
        let regexp = /id=\d+/
        let strId = findID(target, regexp)
        let id = parseInt(strId!.substring(3))
        tasks.get(id)!.clearHTML()
        sendRequest('DELETE', server + `/tasks/${id}`)
        tasks.delete(id)
    } else if (hasParentClass(target, 'name__openModal')) {
        // Просмотр информации о задаче
        let regexp = /id=\d+/
        let strId = findID(target, regexp)
        let id = parseInt(strId!.substring(3))
        document.getElementById('modal__editor')!.setAttribute('opened-task-id', `${id}`)
        let buf = tasks.get(id)
        let titleModal = <HTMLInputElement>document.getElementById('modal-title-info') //меняем заголовок
        titleModal.innerHTML = 'Информация о задаче'
        let btnEditSave = <HTMLInputElement>document.getElementsByClassName('btn-edit-save')[0] //меняем кнопку внизу
        btnEditSave.style.display = 'none'
        if (tasks.get(id)!.getStatus() !== 'archived' || tasks.get(id)!.getStatus() !== 'completed') {
            let btnEditSave = <HTMLInputElement>document.getElementsByClassName('btn-edit')[0] //меняем кнопку внизу
            btnEditSave.style.display = 'block'
        } else {
            let btnEditSave = <HTMLInputElement>document.getElementsByClassName('btn-edit')[0] //меняем кнопку внизу
            btnEditSave.style.display = 'none'
        }
        let formEditCategories = <HTMLInputElement>document.getElementById('form__edit__categories') //убираем список категорий
        formEditCategories.innerHTML = ''
        modalEditor.show()
        let inptBoxInfo = <HTMLInputElement>document.getElementById('nameE') //заполняем форму данными задачи
        let dateBoxInfo = <HTMLInputElement>document.getElementById('dateE')
        let timeBoxInfo = <HTMLInputElement>document.getElementById('timeE')
        let commentTextInfo = <HTMLTextAreaElement>document.getElementById('commentE')
        let fileInfo = <HTMLTextAreaElement>document.getElementById('input__fileE')
        inptBoxInfo.setAttribute('readonly', '') //делаем поля недоступными
        dateBoxInfo.setAttribute('readonly', '')
        timeBoxInfo.setAttribute('readonly', '')
        commentTextInfo.setAttribute('readonly', '')
        fileInfo.setAttribute('readonly', '')
        inptBoxInfo.value = buf?.getName()!
        if (buf?.getDueDate()!) {
            let year = buf?.getDueDate()!.getFullYear()
            let month = buf?.getDueDate()!.getMonth()! + 1
            let day = buf?.getDueDate()!.getDate()!
            dateBoxInfo.value = `${year}-${month > 9 ? month : `0${month}`}-${day > 9 ? day : `0${day}`}`
            let hours = buf?.getDueDate().getHours()
            let minutes = buf?.getDueDate().getMinutes()
            if (hours || minutes) {
                timeBoxInfo.value = buf?.getDueDate().toTimeString()!.substring(0, 5)!
            }
        } else {
            dateBoxInfo.value = ''
            timeBoxInfo.value = ''
        }
        commentTextInfo.value = `${buf?.getDesc()}`
        fileInfo.value = ''
        let taskLabels = <Array<Label>>buf!.getLabels()
        let container = document.getElementsByClassName('chosen__categories')[1]
        taskLabels.length === 0 ? (container.innerHTML = 'Категории отсутствуют') : (container.innerHTML = '')
        for (let i = 0; i < taskLabels.length; i++) {
            let lbl = taskLabels[i]
            container.innerHTML =
                container.innerHTML +
                `<div class="category category__edit col-auto" 
                        style="background-color: #${lbl!.color}99;" id="lbl-${lbl.id}">
                        <label>${lbl.text}</label>
                        <button type="button" style="display: none;" class="btn-close btn-close-lbl btn-close-lbl-edit"></button>
                        </div>`
        }
    } else if (target.classList.contains('btn-edit')) {
        // редактирование задачи
        let titleModal = <HTMLInputElement>document.getElementById('modal-title-info') //меняем заголовок
        titleModal.innerHTML = 'Редактирование задачи'
        let btnEditSave = <HTMLInputElement>document.getElementsByClassName('btn-edit-save')[0] //меняем кнопку внизу
        btnEditSave.style.display = 'block'
        let btnEdit = <HTMLInputElement>document.getElementsByClassName('btn-edit')[0]
        btnEdit.style.display = 'none'
        modalEditor.show()
        let inptBoxEdit = <HTMLInputElement>document.getElementById('nameE')
        let dateBoxEdit = <HTMLInputElement>document.getElementById('dateE')
        let timeBoxEdit = <HTMLInputElement>document.getElementById('timeE')
        let commentTextEdit = <HTMLTextAreaElement>document.getElementById('commentE')
        let fileEdit = <HTMLTextAreaElement>document.getElementById('input__fileE')
        inptBoxEdit.removeAttribute('readonly') // делаем поля изменяемыми
        dateBoxEdit.removeAttribute('readonly')
        timeBoxEdit.removeAttribute('readonly')
        commentTextEdit.removeAttribute('readonly')
        fileEdit.removeAttribute('readonly')
        let container = document.getElementsByClassName('chosen__categories')[1]
        if (container.innerHTML === 'Категории отсутствуют') container.innerHTML = ''
        let taskLabels = container.children
        let btnCloseLbl = document.getElementsByClassName('btn-close-lbl-edit')
        for (let i = 0; i < btnCloseLbl.length; i++) {
            ;(<HTMLElement>btnCloseLbl[i]).style.display = 'block'
        }
        for (let i = 0; i < taskLabels.length; i++) {
            let lbl = taskLabels[i]
            lbl.classList.remove('category__edit')
        }
        let formEditCategories = <HTMLInputElement>document.getElementById('form__edit__categories')
        formEditCategories.innerHTML = `<form class="bg-light p-0 m-0">
                    <input id="edit__search" autocomplete="off" type="search" class="dropbtn2 form-control" autocomplete="false"
                      placeholder="Выбрать категорию">
                  </form>
                  <div id="form-edit__category" class="list-category-contener">
                    <ul class="edit__list-category list-category mb-0 px-0">

                    </ul>

                    <div id="Dropdown2" class="add__category">
                      <div class="row">
                        <div class="m-0 p-0 name__category">
                          <input type="text" class="form-control" id="edit__name__category" placeholder="Новая категория">
                        </div>
                        <div class="m-0 color__category">
                          <input class="h-100 color__category" type="color" title="Задать цвет"
                            id="edit__color__category">
                        </div>
                        <div class="add p-0">
                          <a class="bg-primary text-white active d-inline-block edit__add-lbl" title="Добавить" href="#">
                            <svg xmlns="http://www.w3.org/2000/svg" width="37" height="37" fill="currentColor"
                              class="bi bi-plus" viewBox="0 0 16 16">
                              <path
                                d="M8 4a.5.5 0 0 1 .5.5v3h3a.5.5 0 0 1 0 1h-3v3a.5.5 0 0 1-1 0v-3h-3a.5.5 0 0 1 0-1h3v-3A.5.5 0 0 1 8 4z" />
                            </svg>
                          </a>
                        </div>
                      </div>
                    </div>
                  </div>`
        let colorInput = <HTMLInputElement>document.getElementById('edit__color__category')
        colorInput.value = '#' + Math.random().toString(16).slice(-6)
    } else if (target.classList.contains('btn-edit-save')) {
        // сохранение изменений при редактировании
        let id = document.getElementById('modal__editor')!.getAttribute('opened-task-id')!
        let nameEdit = (<HTMLInputElement>document.getElementById('nameE')).value
        let dateEdit = (<HTMLInputElement>document.getElementById('dateE')).value
        let timeEdit = (<HTMLInputElement>document.getElementById('timeE')).value
        let date = dateEdit + ' ' + timeEdit
        let descEdit = `${(<HTMLTextAreaElement>document.getElementById('commentE')).value}`
        let fileEdit = (<HTMLTextAreaElement>document.getElementById('input__fileE')).value
        let lbls = document.getElementsByClassName('chosen__categories')[1].children
        let taskLabels: Label[] = []
        for (let i = 0; i < lbls.length; i++) {
            taskLabels.push(labels.get(parseInt(lbls[i].id.substring(4)) - 1)!)
        }

        if (date !== ' ') {
            tasks.get(parseInt(id))!.editTask(nameEdit, date, taskLabels, descEdit)
        } else {
            tasks.get(parseInt(id))!.editTask(nameEdit, undefined, taskLabels, descEdit)
        }
    } else if (target.classList.contains('btn-close-lbl')) {
        // Удаление лейбла в модалке
        target.parentElement!.outerHTML = ''
    } else if (hasParentClass(target, 'dropdown-item') && hasParentClass(target, 'edit__list-category')) {
        // добавление лейбла в контейнер при редактировании
        let link = <HTMLElement>target.closest('.dropdown-item')
        addLabel(1, link.id)
    } else if (hasParentClass(target, 'dropdown-item') && hasParentClass(target, 'list-category')) {
        // добавление лейбла в контейнер при создании задачи
        let link = <HTMLElement>target.closest('.dropdown-item')
        addLabel(0, link.id)
    } else if (hasParentClass(target, 'add-lbl')) {
        // создание лейбла в форме расширенного создания задачи
        let textInput = <HTMLInputElement>document.getElementById('name__category')
        let colorInput = <HTMLInputElement>document.getElementById('color__category')
        makeLabel(colorInput!.value, textInput.value, 'list-category', 0)
        textInput.value = ''
        colorInput.value = '#' + Math.random().toString(16).slice(-6)
    } else if (hasParentClass(target, 'edit__add-lbl')) {
        // создание лейбла в форме редактирования задачи
        let textInput = <HTMLInputElement>document.getElementById('edit__name__category')
        let colorInput = <HTMLInputElement>document.getElementById('edit__color__category')
        makeLabel(colorInput.value, textInput.value, 'edit__list-category', 1)
        textInput.value = ''
        colorInput.value = '#' + Math.random().toString(16).slice(-6)
    } else if (hasParentClass(target, 'dropdown-item') && hasParentClass(target, 'sort__list-category')) {
        //сортировка отображаемых задач по выбранному лейблу
        let link = <HTMLElement>target.closest('.dropdown-item')
        let id = parseInt(link.id.substring(4)) - 1
        let lbl = labels.get(id)!
        let flt = <HTMLSelectElement>document.getElementsByClassName('stage')[0]
        let status = flt.selectedOptions[0].value
        if (sortLabels.get(lbl.id)) {
            sortLabels.delete(lbl.id)
            link.style.backgroundColor = `#0000`
        } else {
            link.style.backgroundColor = `#${lbl.color}99`
            sortLabels.set(lbl.id, true)
        }
        let list = <HTMLUListElement>document.getElementById('todolist')
        list.innerHTML = ''
        tasks.forEach((task) => {
            if (task.getStatus() === status || status === 'all') {
                if (sortLabels.size === 0) {
                    task.toHTMLBlock()
                } else {
                    let lbls = task.getLabels()
                    for (let i = 0; i < lbls.length; i++) {
                        if (sortLabels.has(lbls[i].id)) {
                            task.toHTMLBlock()
                            break
                        }
                    }
                }
            }
        })
    } else if (target.classList.contains('sort__btn-all')) {
        // сбросить сортировку по лейблам
        sortLabels.clear()
        let list = <HTMLUListElement>document.getElementById('todolist')
        list.innerHTML = ''
        let flt = <HTMLSelectElement>document.getElementsByClassName('stage')[0]
        let status = flt.selectedOptions[0].value
        tasks.forEach((task) => {
            if (task.getStatus() === status || status === 'all') task.toHTMLBlock()
        })
        let elem = <HTMLDivElement>document.getElementById('sort__form__category')
        elem.classList.remove('show')
        let ul = <HTMLUListElement>document.getElementsByClassName('sort__list-category')[0]
        ul.innerHTML = ''
    } else if (target.classList.contains('dropbtn')) {
        //выпадающий список лейблов в сортировке
        target.setAttribute('placeholder', 'Поиск')
        if (openUl('sort__form__category', 'show', 'sort__search', 'sort__list-category'))
            target.setAttribute('placeholder', 'Категории')
    } else if (
        !(target.matches('.dropbtn') || document.getElementsByClassName('sort__list-category')[0]?.contains(target)) &&
        (<HTMLDivElement>document.getElementById('sort__form__category'))?.classList.contains('show')
    ) {
        //сворачивание списка лейблов в сортировке задач по лейблам
        closeUl('dropbtn', 'sort__form__category', 'sort__list-category', 'sort__search', 'show')
    } else if (target.classList.contains('dropbtn1')) {
        // выпадающий список лейблов в создании задачи
        target.setAttribute('placeholder', 'Поиск')
        if (openUl('form__category', 'show1', 'add-task__search', 'list-category'))
            target.setAttribute('placeholder', 'Выбрать категорию')
    } else if (
        !(target.matches('.dropbtn1') || document.getElementById('form__category')?.contains(target)) &&
        (<HTMLDivElement>document.getElementById('form__category'))?.classList.contains('show1')
    ) {
        // сворачивание списка лейблов в создании задачи
        closeUl('dropbtn1', 'form__category', 'list-category', 'add-task__search', 'show1')
    } else if (target.classList.contains('dropbtn2')) {
        // выпадающий список лейблов в редактировании задачи
        target.setAttribute('placeholder', 'Поиск')
        if (openUl('form-edit__category', 'show1', 'edit__search', 'edit__list-category'))
            target.setAttribute('placeholder', 'Выбрать категорию')
    } else if (
        !target.matches('.dropbtn2') &&
        document.getElementsByClassName('edit__list-category').length != 0 &&
        !document.getElementById('form-edit__category')!.contains(target) &&
        (<HTMLDivElement>document.getElementById('form-edit__category'))?.classList.contains('show1')
    ) {
        // сворачивание списка лейблов в редактировании задачи
        closeUl('dropbtn2', 'form-edit__category', 'edit__list-category', 'edit__search', 'show1')
    }
})

// Обработка событий изменений

document.addEventListener('change', (e) => {
    let target = <HTMLElement>e.target
    if (target.classList.contains('stage')) {
        // код фильтра
        let flt = <HTMLSelectElement>target
        let status = flt.selectedOptions[0].value
        let ul = <HTMLUListElement>document.getElementById('todolist')
        ul.innerHTML = ''
        tasks.forEach((task) => {
            if (task.getStatus() === status || status === 'all') {
                if (sortLabels.size === 0) {
                    task.toHTMLBlock()
                } else {
                    let lbls = task.getLabels()
                    for (let i = 0; i < lbls.length; i++) {
                        if (sortLabels.has(lbls[i].id)) {
                            task.toHTMLBlock()
                            break
                        }
                    }
                }
            }
        })
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

document.addEventListener('input', (e) => {
    let target = <HTMLElement>e.target
    if (target.id === 'sort__search') {
        let inpt = (<HTMLInputElement>target).value
        search('sort__list-category', inpt)
    } else if (target.id === 'add-task__search') {
        let inpt = (<HTMLInputElement>target).value
        search('list-category', inpt)
    } else if (target.id === 'edit__search') {
        let inpt = (<HTMLInputElement>target).value
        search('edit__list-category', inpt)
    }
})

// Создает список с двумя рабочими элементами
function generateExample(): void {
    let foo = new Task(glob_id + 1, 'first task', 'моя задача очень важна', '07-10-2022 22:22')
    foo.toHTMLBlock()
    tasks.set(glob_id++, foo)
    let foo1 = new Task(glob_id + 1, 'second task', 'моя задача очень важна', '07-30-2022 21:21')
    foo1.toHTMLBlock()
    tasks.set(glob_id++, foo1)
}

// Отдельные обработчики событий

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
            tasks.set(glob_id++, newTask)
            newTask.toHTMLBlock()
        }
    }
}
