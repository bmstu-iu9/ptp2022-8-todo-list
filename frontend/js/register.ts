// @ts-check

window.onload = function () {

    const rgBtn = <HTMLButtonElement>document.getElementsByClassName('btn-regis')[0]
    const lgBtn = <HTMLButtonElement>document.getElementsByClassName('btn-login')[0]
    const lgForm = <HTMLFormElement>document.getElementsByClassName('login')[0]
    const rgForm = <HTMLFormElement>document.getElementsByClassName('register')[0]

    rgBtn.addEventListener('click', () => {
        rgForm.classList.add('hidden')
        lgForm.classList.remove('hidden')
        document.title = "Sign in"
    })
    lgBtn.addEventListener('click', () => {
        lgForm.classList.add('hidden')
        rgForm.classList.remove('hidden')
        document.title = "Sign up"
    })
}

