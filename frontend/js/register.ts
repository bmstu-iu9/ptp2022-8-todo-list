window.onload = function () {
    if (localStorage.getItem('theme') === 'light') changeModeLogin()
    const re = new RegExp(/^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/)


    const rgBtn = <HTMLButtonElement>document.getElementsByClassName('btn-regis')[0]
    const lgBtn = <HTMLButtonElement>document.getElementsByClassName('btn-login')[0]
    const lgForm = <HTMLFormElement>document.getElementsByClassName('login')[0]
    const rgForm = <HTMLFormElement>document.getElementsByClassName('register')[0]
    const eye = <HTMLElement>document.getElementsByClassName('password-control')[0]
    const register = document.getElementById('register')
    const login = document.getElementById('login')
    let footer = <HTMLElement>document.querySelector('footer')
    let check = <HTMLElement>document.getElementsByClassName('password-check')[0]
    let check1 = <HTMLElement>document.getElementsByClassName('login-check')[0]

    rgBtn.addEventListener('click', () => {
        rgForm.classList.add('hidden')
        lgForm.classList.remove('hidden')
        document.title = 'Sign in'
        footer.style.margin = '1rem 0 0 0'
        check.style.visibility = 'hidden'
    })
    lgBtn.addEventListener('click', () => {
        lgForm.classList.add('hidden')
        rgForm.classList.remove('hidden')
        document.title = 'Sign up'
        footer.style.margin = '110px 0 0 0'
        check1.style.visibility = 'hidden'
    })

    login!.addEventListener('click', () => {
        let email = (<HTMLInputElement>document.getElementById('email_lg')).value
        let password = (<HTMLInputElement>document.getElementById('rgPassword')).value

        if (password.length < 8 || !re.test(email.toLowerCase())) {
            check1.style.visibility = 'visible'
        }
    })

    register!.addEventListener('click', () => {
        let password1 = (<HTMLInputElement>document.getElementById('rgPassword')).value
        let password2 = (<HTMLInputElement>document.getElementById('confirmPassword')).value
        let email = (<HTMLInputElement>document.getElementById('email_rg')).value
        let nickname = (<HTMLInputElement>document.getElementById('nickname')).value
        if (nickname.length == 0 || email.length == 0 || password1.length == 0 || password2.length == 0) {
            check.innerHTML = 'Заполнены не все поля'
            check.style.visibility = 'visible'
        } else if (password1.length < 8) {
            check.innerHTML = 'Длина пароля меньше 8 символов'
            check.style.visibility = 'visible'
        } else if (password1 !== password2) {
            check.innerHTML = 'Пароли не совпадают'
            check.style.visibility = 'visible'
        } else if (email === '' || !re.test(email.toLowerCase())) {
            check.innerHTML = `Проверьте ваш email-адрес`
            check.style.visibility = 'visible'
        } else {
            check.style.visibility = 'hidden'
        }
    })

    eye.addEventListener('click', () => {
        let password = document.getElementById('lgPassword')
        if (password!.getAttribute('type') === 'password') {
            eye.innerHTML = `<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="currentColor" class="bi bi-eye-slash-fill" viewBox="0 0 16 16">
            <path d="m10.79 12.912-1.614-1.615a3.5 3.5 0 0 1-4.474-4.474l-2.06-2.06C.938 6.278 0 8 0 8s3 5.5 8 5.5a7.029 7.029 0 0 0 2.79-.588zM5.21 3.088A7.028 7.028 0 0 1 8 2.5c5 0 8 5.5 8 5.5s-.939 1.721-2.641 3.238l-2.062-2.062a3.5 3.5 0 0 0-4.474-4.474L5.21 3.089z"/>
            <path d="M5.525 7.646a2.5 2.5 0 0 0 2.829 2.829l-2.83-2.829zm4.95.708-2.829-2.83a2.5 2.5 0 0 1 2.829 2.829zm3.171 6-12-12 .708-.708 12 12-.708.708z"/>
            </svg>`
            password!.setAttribute('type', 'text')
        } else {
            eye.innerHTML = `<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="currentColor"
            class="bi bi-eye-fill" viewBox="0 0 16 16">
            <path d="M10.5 8a2.5 2.5 0 1 1-5 0 2.5 2.5 0 0 1 5 0z" />
            <path d="M0 8s3-5.5 8-5.5S16 8 16 8s-3 5.5-8 5.5S0 8 0 8zm8 3.5a3.5 3.5 0 1 0 0-7 3.5 3.5 0 0 0 0 7z" />
            </svg>`
            password!.setAttribute('type', 'password')
        }
    })
}


function changeModeLogin() {
    document.body.classList.toggle('bg-dark')
    document.body.classList.toggle('bg-light')
    document.body.classList.toggle('text-white')
}
