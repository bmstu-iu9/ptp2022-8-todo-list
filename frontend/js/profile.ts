const changePasswordForm: HTMLFormElement = document.getElementById("newPasswordInput") as HTMLFormElement;
const passwordRepeatForm: HTMLFormElement = document.getElementById("newPasswordRepeatInput") as HTMLFormElement;
const changePasswordSubmitBtn: HTMLFormElement = document.getElementById("changePasswordSubmitBtn") as HTMLFormElement;
const validationResSpan: HTMLElement = document.getElementById("validationResult") as HTMLElement;
const closeChangePswrdModaldBtn: HTMLElement = document.getElementById("closeChangePasswordModalBtn") as HTMLElement;
const usernameField: HTMLElement = document.getElementById("usernameField") as HTMLElement;
const nameField: HTMLElement = document.getElementById("nameField") as HTMLElement;
const surnameField: HTMLElement = document.getElementById("surnameField") as HTMLElement;
const emailField: HTMLElement = document.getElementById("emailField") as HTMLElement;
const aboutInfoField: HTMLElement = document.getElementById("aboutInfoField") as HTMLElement;
const changeDataBtn: HTMLElement = document.getElementById("changeDataBtn") as HTMLElement;
const changePasswordModal: HTMLElement = document.getElementById("changePasswordModal") as HTMLElement;
const changeUsernameForm: HTMLFormElement = document.getElementById("usernameInput") as HTMLFormElement;
const changeNameForm: HTMLFormElement = document.getElementById("nameInput") as HTMLFormElement;
const changeSurnameForm: HTMLFormElement = document.getElementById("surnameInput") as HTMLFormElement;
const changeAboutInfoForm: HTMLFormElement = document.getElementById("aboutInfoInput") as HTMLFormElement;
const closeDataChangeModalBtn: HTMLElement = document.getElementById("closeDataChangeModalBtn") as HTMLElement;
const changeDataSubmitBtn: HTMLElement = document.getElementById("changeDataSubmitBtn") as HTMLElement;


function clearChangePswrdForm(): void {
    changePasswordForm.value = "";
    passwordRepeatForm.value = "";
    validationResSpan.innerHTML = "";
}

changeDataBtn.addEventListener("click", () => {
    changeUsernameForm.value = usernameField.innerHTML.slice(1);
    changeNameForm.value = nameField.innerHTML;
    changeSurnameForm.value = surnameField.innerHTML;
    changeAboutInfoForm.value = aboutInfoField.innerHTML;
});

passwordRepeatForm.addEventListener("input", (evt) => {
    if (passwordRepeatForm.value == changePasswordForm.value) {
        validationResSpan.innerHTML = "&#10003 Пароли совпадают";
        validationResSpan.style.color = "green";
    } else {
        validationResSpan.innerHTML = "&#10060 Пароли не совпадают";
        validationResSpan.style.color = "red";
    }
});

closeChangePswrdModaldBtn.addEventListener("click", clearChangePswrdForm);
changePasswordModal.addEventListener("click", (evt) => {
    if (evt.target == changePasswordModal) {
        clearChangePswrdForm();
    }
});


