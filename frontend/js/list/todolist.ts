let glob_id = 0
let lbl_id = 0
declare var bootstrap: any

// Текущее хранилище задач
var tasks = new Map<number, Task>()
// Текущее хранилище лейблов
var labels = new Map<number, Label>()
// модальная форма ввода задачи
const modal = new bootstrap.Modal(<HTMLFormElement>document.getElementById('modal'))
// модальная форма редактирования задачи
const modalEditor = new bootstrap.Modal(<HTMLFormElement>document.getElementById('modal__editor'))

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
        let newTask = new Task(buf.id, buf.name, buf.description, buf.dueDate, buf.labels, false)
        newTask.setStatus(buf.status, false)
        tasks.set(buf.id, newTask)
        if (newTask.getStatus() === 'active' || newTask.getStatus() === 'completed') {
            newTask.toHTMLBlock()
        }
        glob_id = buf.id + 1
    }
})

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
        } else {
            chckBox!.removeAttribute('checked')
            tasks.get(parseInt(id))!.setStatus('active')
        }
    } else if (target.classList.contains('add-btn')) {
        // Быстрое создание задачи
        let inptBox = <HTMLInputElement>document.getElementsByClassName('todo_text')[0]
        if (inptBox!.value !== '') {
            let name = inptBox!.value
            inptBox!.value = ''
            let newTask = new Task(glob_id + 1, name)
            tasks.set(glob_id++, newTask)
            newTask.toHTMLBlock()
        }
    } else if (target.classList.contains('add-extended')) {
        // Расширенное создание задачи
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
    } else if (hasParentClass(target, 'todo__delete')) {
        // удаление задачи (TODO: понять, что с ней делать)
        let regexp = /id=\d+/
        let strId = findID(target, regexp)
        let id = parseInt(strId!.substring(3))
        tasks.get(id)!.setStatus('archived')
        tasks.get(id)!.clearHTML()
        sendRequest('DELETE', server + `/tasks/${id}`)
        tasks.delete(id)
    } else if (hasParentClass(target, 'name__openModal')) {
        // Просмотр информации о задаче
        let regexp = /id=\d+/
        let strId = findID(target, regexp)
        let id = parseInt(strId!.substring(3))
        let modal = document.getElementById('modal__editor')
        modal?.setAttribute('opened-task-id', `${id}`)
        let buf = tasks.get(id)
        let titleModal = <HTMLInputElement>document.getElementById('modal-title-info') //меняем заголовок
        titleModal.innerHTML = 'Информация о задаче'
        let btnEditSave = <HTMLInputElement>document.getElementsByClassName('btn-edit-save')[0] //меняем кнопку внизу
        btnEditSave.style.display = 'none'
        let btnEdit = <HTMLInputElement>document.getElementsByClassName('btn-edit')[0]
        btnEdit.style.display = 'block'
        let formEditCategories = <HTMLInputElement>document.getElementById('form__edit__categories') //убираем список категорий
        formEditCategories.innerHTML = ''
        modalEditor.show()
        let inptBoxInfo = <HTMLInputElement>document.getElementById('nameE') //заполняем форму данными задачи
        let dateBoxInfo = <HTMLInputElement>document.getElementById('dateE')
        let timeBoxInfo = <HTMLInputElement>document.getElementById('timeE')
        let commentTextInfo = <HTMLTextAreaElement>document.getElementById('commentE')
        let fileEdit = <HTMLTextAreaElement>document.getElementById('input__fileE')
        inptBoxInfo.setAttribute('readonly', '') //делаем поля недоступными
        dateBoxInfo.setAttribute('readonly', '')
        timeBoxInfo.setAttribute('readonly', '')
        commentTextInfo.setAttribute('readonly', '')
        fileEdit.setAttribute('readonly', '')
        inptBoxInfo.value = buf?.getName()!
        dateBoxInfo.value = ''
        commentTextInfo.value = ''
        let taskLabels = <Array<Label>>buf?.getLabels()
        let container = document.getElementsByClassName('chosen__categories')[1]
        container.innerHTML = ''
        for (let i = 0; i < taskLabels.length; i++) {
            let lbl = taskLabels[i]
            container.innerHTML =
                container.innerHTML +
                `<div class="category category__edit col-auto" 
                        style="background-color: ${lbl.color};" id="lbl-${lbl.id}">
                        <label>${lbl.text}</label>
                        <button type="button" style="display: none;" class="btn-close btn-close-lbl btn-close-lbl-edit"></button>
                        </div>`
        }
    } else if (target.classList.contains('btn-edit')) {
        // редактирование задачи
        let modal = document.getElementById('modal__editor')
        let id = modal?.getAttribute('opened-task-id')
        let buf = JSON.parse(localStorage[parseInt(id!)])
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
        let taskLabels = container.children
        let btnCloseLbl = document.getElementsByClassName('btn-close-lbl-edit')
        for (let i = 0; i < btnCloseLbl.length; i++) {
            ;(<HTMLElement>btnCloseLbl[i]).style.display = 'block'
        }
        for (let i = 0; i < taskLabels.length; i++) {
            let lbl = taskLabels[i]
            lbl.classList.remove('category__edit')
        }
        let formEditCategories = <HTMLInputElement>document.getElementById('form__edit__categories') //добавляем список категорий  (надо переделать открывание списка...для нескольких форм)
        formEditCategories.innerHTML = `<form class="bg-light p-0 m-0">
                    <input type="search" class="dropbtn2 form-control" autocomplete="false"
                      placeholder="Выбрать категорию">
                  </form>
                  <div id="form__category" class="list-category-contener">
                    <ul class="list-category mb-0 px-0">
                      <!-- <input class="form-control" id="myInput" type="text" placeholder="Поиск"> -->
                    </ul>

                    <div id="Dropdown2" class="add__category">
                      <div class="row">
                        <div class="m-0 p-0 name__category">
                          <input type="text" class="form-control" id="name__category" placeholder="Новая категория">
                        </div>
                        <div class="m-0 color__category">
                          <input class="h-100 color__category" type="color" value="#563d7c" title="Задать цвет"
                            id="color__category">
                        </div>
                        <div class="add p-0">
                          <a class="bg-primary text-white active d-inline-block add-lbl" title="Добавить" href="#">
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
        //нужно сделать сохранение введеных данных по клику на кнопку класса btn-edit-save
    } else if (target.classList.contains('btn-close-lbl')) {
        // Удаление лейбла в модалке
        target.parentElement!.outerHTML = ''
    } else if (hasParentClass(target, 'dropdown-item') && hasParentClass(target, 'list-category')) {
        // добавление лейбла в контейнер
        let regexp = /id=lbl-\d+/
        let link = <HTMLElement>target.closest('.dropdown-item')
        let lbl = labels.get(parseInt(link.id.substring(4)) - 1)
        let container = document.getElementsByClassName('chosen__categories')[0]
        container.innerHTML =
            container.innerHTML +
            `<div class="category col-auto" 
                        style="background-color: ${lbl?.color};" id="lbl-${lbl?.id}">
                        <label>${lbl?.text}</label>
                        <button type="button" class="btn-close btn-close-lbl"></button>
                        </div>`
        LabelToHtml(<Label>lbl)
    } else if (hasParentClass(target, 'add-lbl')) {
        let textInput = <HTMLInputElement>document.getElementById('name__category')
        let colorInput = <HTMLInputElement>document.getElementById('color__category')
        let colorHex = colorInput!.value
        let lbl: Label = {
            color: colorHex,
            text: textInput.value,
            id: lbl_id,
        }
        labels.set(lbl_id++, lbl)
        let ul = <HTMLUListElement>document.getElementsByClassName('list-category')[0]
        let lbl_text = `<li><a class="dropdown-item d-flex align-items-center gap-2 py-1" href="#" id="lbl-${lbl.id}">
                        <span style="background-color: ${lbl.color};"
                        class="d-inline-block rounded-circle p-1"></span>
                        <label class="todo__category">${lbl.text}</label>
                    </a></li>`
        ul.innerHTML = ul.innerHTML.concat(lbl_text)
    } else if (target.classList.contains('dropbtn2')) {
        // отрисовка выбранных лейблов в списке
        let elem = <HTMLDivElement>document.getElementById('form__category')
        if (elem.classList.contains('show1')) {
            elem.classList.remove('show1')
            let ul = <HTMLUListElement>document.getElementsByClassName('list-category')[0]
            ul.innerHTML = ''
        } else {
            let ul = <HTMLUListElement>document.getElementsByClassName('list-category')[0]
            labels.forEach((lbl) => {
                let lbl_text = `<li><a class="dropdown-item d-flex align-items-center gap-2 py-1" href="#" id="lbl-${lbl.id}">
                                <span style="background-color: ${lbl.color};"
                                class="d-inline-block rounded-circle p-1"></span>
                                <label class="todo__category">${lbl.text}</label>
                            </a></li>`
                ul.innerHTML = ul.innerHTML.concat(lbl_text)
            })
            elem.classList.toggle('show1')
        }
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
            if (task.getStatus() === status) {
                task.toHTMLBlock()
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
            tasks.set(glob_id++, newTask)
            newTask.toHTMLBlock()
        }
    }
}
