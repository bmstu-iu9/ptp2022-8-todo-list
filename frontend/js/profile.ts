const changePasswordForm = document.getElementById("newPasswordInput") as HTMLFormElement;
const passwordRepeatForm = document.getElementById("newPasswordRepeatInput") as HTMLFormElement;
const changePasswordSubmitBtn = document.getElementById("changePasswordSubmitBtn") as HTMLFormElement;
const validationResSpan = document.getElementById("validationResult") as HTMLElement;
const closeChangePswrdModaldBtn = document.getElementById("closeChangePasswordModalBtn") as HTMLElement;
const changePasswordModal = document.getElementById("changePasswordModal") as HTMLElement;
 
passwordRepeatForm.addEventListener("input", (evt) => {
    if (passwordRepeatForm.value == changePasswordForm.value) {
        validationResSpan.innerHTML = "&#10003 Пароли совпадают";
        validationResSpan.style.color = "green";
    } else {
        validationResSpan.innerHTML = "&#10060 Пароли не совпадают";
        validationResSpan.style.color = "red";
    }
});

function clearChangePswrdForm(): void {
    changePasswordForm.value = "";
    passwordRepeatForm.value = "";
    validationResSpan.innerHTML = "";
}

closeChangePswrdModaldBtn.addEventListener("click", clearChangePswrdForm);
changePasswordModal.addEventListener("click", (evt) => {
    if (evt.target == changePasswordModal) {
        clearChangePswrdForm();
    }
});


