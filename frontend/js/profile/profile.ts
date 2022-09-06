type User = {
    nickname: string
    name: string
    email: string
    surname: string
    aboutInfo: string
    level: number
    balance: number
    health: number
    experience: number
}
let userId: number = 3

function changeDataModal() {
    let nickname = document.getElementById('nicknameField')
    let name = document.getElementById('nameField')
    let surname = document.getElementById('surnameField')
    let email = document.getElementById('emailField')
    let aboutInfo = document.getElementById('aboutInfoField')
    let balance = document.getElementById('balance')

    const changeDataBtn = document.getElementById('changeDataBtn') as HTMLElement
    const modal: HTMLElement = document.getElementById('changeDataModal') as HTMLElement
    const nicknameForm: HTMLFormElement = document.getElementById('nicknameInput') as HTMLFormElement
    const nameForm: HTMLFormElement = document.getElementById('nameInput') as HTMLFormElement
    const surnameForm: HTMLFormElement = document.getElementById('surnameInput') as HTMLFormElement
    const aboutInfoForm: HTMLFormElement = document.getElementById('aboutInfoInput') as HTMLFormElement
    const closeBtn: HTMLElement = document.getElementById('closeDataChangeModalBtn') as HTMLElement
    const submitBtn: HTMLElement = document.getElementById('changeDataSubmitBtn') as HTMLElement
    const switchThemeBtn: HTMLInputElement = document.getElementById('theme-switch-checkbox') as HTMLInputElement
    if (localStorage.getItem('theme') == 'light') {
        switchThemeBtn.checked = true
    } else {
        switchThemeBtn.checked = false
    }
    changeDataBtn.addEventListener('click', () => {
        nicknameForm.value = nickname?.innerHTML.slice(1)
        nameForm.value = name?.innerHTML
        surnameForm.value = surname?.innerHTML
        aboutInfoForm.value = aboutInfo?.innerHTML
    })

    function getDataChanges(): {
        [key: string]: string
    } {
        let changes: {
            [key: string]: string
        } = {}
        if (nicknameForm.value != nickname?.innerHTML.slice(1)) {
            changes.nickname = nicknameForm.value
        }
        if (surnameForm.value != surname?.innerHTML) {
            changes.surname = surnameForm.value
        }
        if (nameForm.value != name?.innerHTML) {
            changes.name = nameForm.value
        }
        if (aboutInfoForm.value != aboutInfo?.innerHTML) {
            changes.aboutInfo = aboutInfoForm.value
        }
        return changes
    }
    submitBtn.addEventListener('click', () => {
        let changes = getDataChanges()
        if (Object.keys(changes).length > 0) {
            sendRequest('PATCH', server + `/users/${userId}`, JSON.stringify(changes))
                .then((r) => {
                    updateHTML(r)
                })
                .catch((e) => {
                    throw e
                })
        }
        if (switchThemeBtn.checked == true && localStorage.getItem('theme') == 'dark') {
            localStorage.setItem('theme', 'light')
            switchProfileTheme()
        } else if (switchThemeBtn.checked == false && localStorage.getItem('theme') == 'light') {
            localStorage.setItem('theme', 'dark')
            switchProfileTheme()
        }
    })
}

function changePasswordModal() {
    const modal: HTMLElement = document.getElementById('changePasswordModal') as HTMLElement
    const passwordForm: HTMLFormElement = document.getElementById('newPasswordInput') as HTMLFormElement
    const passwordRepeatForm: HTMLFormElement = document.getElementById('newPasswordRepeatInput') as HTMLFormElement
    const submitBtn: HTMLFormElement = document.getElementById('changePasswordSubmitBtn') as HTMLFormElement
    const validationSpan: HTMLElement = document.getElementById('validationResult') as HTMLElement
    const closeBtn: HTMLElement = document.getElementById('closeChangePasswordModalBtn') as HTMLElement
    const oldPasswordForm: HTMLFormElement = document.getElementById('oldPasswordInput') as HTMLFormElement

    function clearPasswordForm(): void {
        oldPasswordForm.value = ''
        passwordForm.value = ''
        passwordRepeatForm.value = ''
        validationSpan.innerHTML = ''
    }

    function setSuccesValidationStyles(msg: string): void {
        validationSpan.innerHTML = `&#10003 ${msg}`
        validationSpan.style.color = 'green'
    }

    function setFailValidationStyles(msg: string): void {
        validationSpan.innerHTML = `&#10060 ${msg}`
        validationSpan.style.color = 'red'
    }

    passwordForm.addEventListener('input', () => {
        if (passwordRepeatForm.value != '') {
            if (passwordRepeatForm.value === passwordForm.value) {
                setSuccesValidationStyles('Пароли совпадают')
            } else {
                setFailValidationStyles('Пароли не совпадают')
            }
        }
    })

    passwordRepeatForm.addEventListener('input', () => {
        if (passwordRepeatForm.value == passwordForm.value) {
            setSuccesValidationStyles('Пароли совпадают')
        } else {
            setFailValidationStyles('Пароли не совпадают')
        }
    })

    submitBtn.addEventListener('click', () => {
        if (passwordForm.value === passwordRepeatForm.value) {
            if (passwordForm.value.length < 8) {
                setFailValidationStyles('Длина нового пароля меньше 8 символов')
            } else {
                /* тут должен отправляться запрос на сервер*/
            }
        }
    })

    closeBtn.addEventListener('click', clearPasswordForm)
    modal.addEventListener('click', (evt) => {
        if (evt.target == modal) {
            clearPasswordForm()
        }
    })
}

function updateHTML(user: User): void {
    let nickname = document.getElementById('nicknameField')
    let name = document.getElementById('nameField')
    let surname = document.getElementById('surnameField')
    let email = document.getElementById('emailField')
    let aboutInfo = document.getElementById('aboutInfoField')
    let balance = document.getElementById('balance')

    name!.innerHTML = user.name ?? 'Не указано'
    balance!.innerHTML = String(user.balance) ?? '0'
    surname!.innerHTML = user.surname ?? 'Не укзана'
    email!.innerHTML = user.email ?? 'Неизвестно'
    nickname!.innerHTML = '@' + user.nickname ?? '@Неизвестно'
    aboutInfo!.innerHTML = user.aboutInfo ?? 'Не указана'
}

function switchElementTheme(el: HTMLElement | Element): void {
    el.classList.toggle('text-white')
    el.classList.toggle('bg-dark')
    el.classList.toggle('border-light')
}

function switchProfileTheme() {
    document.body.classList.toggle('bg-dark')
    document.querySelectorAll('section').forEach((section) => switchElementTheme(section))
    document.querySelectorAll('.container').forEach((container) => switchElementTheme(container))
    document.querySelectorAll('.card').forEach((card) => switchElementTheme(card))
    document.querySelectorAll('.list-group-item').forEach((li) => switchElementTheme(li))
}

function onProfileLoad(): void {
    /* if (localStorage.getItem('theme') == 'dark') {
        switchProfileTheme()
    } */
    changeDataModal()
    changePasswordModal()
    sendRequest('GET', server + `/users/${userId}`)
        .then((r) => {
            updateHTML(r)
            let e: Equipment = getEquipment(r.Items)
            renderHero('user-img-card', e)
            let petCard: HTMLElement = document.getElementById('pet-container')
            if (e.pet) {
                petCard.appendChild(createEquipHtml(e.pet))
            }
        })
        .catch((e) => {
            throw e
        })
}

try {
    onProfileLoad()
} catch (error) {}
