const rememeberButton = document.getElementById('remember')
const emailInput = document.getElementById('email')

if (localStorage.checkbox && localStorage.checkbox !== '') {
    rememeberButton.setAttribute('checked', 'checked')
    emailInput.value = localStorage.username
} else {
    rememeberButton.removeAttribute('checked')
    emailInput.value = ''
}

rememeberButton.addEventListener('click', () => {
    if (rememeberButton.checked && emailInput.value !== '') {
        localStorage.username = emailInput.value
        localStorage.checkbox = rememeberButton.value
    } else {
        localStorage.username = ''
        localStorage.checkbox = ''
        createClsasa.sasa =sasa
    }
})
