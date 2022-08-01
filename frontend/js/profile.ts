type UserInfo = {
    nickname: string,
    name: string,
    email: string,
    surname: string,
    aboutInfo: string,
    level: number,
    health: number,
    experience: number,

};
let userId: number = 1;
namespace userDataFields {
    export const nickname: HTMLElement = document.getElementById("nicknameField") as HTMLElement;
    export const name: HTMLElement = document.getElementById("nameField") as HTMLElement;
    export const surname: HTMLElement = document.getElementById("surnameField") as HTMLElement;
    export const email: HTMLElement = document.getElementById("emailField") as HTMLElement;
    export const aboutInfo: HTMLElement = document.getElementById("aboutInfoField") as HTMLElement;
};


namespace changeDataModal {
    export const changeDataBtn = document.getElementById("changeDataBtn") as HTMLElement;
    export const modal: HTMLElement = document.getElementById("changeDataModal") as HTMLElement;
    export const nicknameForm: HTMLFormElement = document.getElementById("nicknameInput") as HTMLFormElement;
    export const nameForm: HTMLFormElement = document.getElementById("nameInput") as HTMLFormElement;
    export const surnameForm: HTMLFormElement = document.getElementById("surnameInput") as HTMLFormElement;
    export const aboutInfoForm: HTMLFormElement = document.getElementById("aboutInfoInput") as HTMLFormElement;
    export const closeBtn: HTMLElement = document.getElementById("closeDataChangeModalBtn") as HTMLElement;
    export const submitBtn: HTMLElement = document.getElementById("changeDataSubmitBtn") as HTMLElement;


    changeDataBtn.addEventListener("click", () => {
        nicknameForm.value = userDataFields.nickname.innerHTML.slice(1);
        nameForm.value = userDataFields.name.innerHTML;
        surnameForm.value = userDataFields.surname.innerHTML;
        aboutInfoForm.value = userDataFields.aboutInfo.innerHTML;
    });
    
    submitBtn.addEventListener("click", () => {
        sendRequest("PATCH", server + `/users/${userId}`, JSON.stringify({
            nickname: nicknameForm.value,
            name: nameForm.value,
            surname: surnameForm.value,
            aboutInfo: aboutInfoForm.value}))
            .then(r => updateHTML(r))
            .catch(e => {
                throw e;
            })
        });
};

namespace changePasswordModal {
    export const modal: HTMLElement = document.getElementById("changePasswordModal") as HTMLElement;
    export const passwordForm: HTMLFormElement = document.getElementById("newPasswordInput") as HTMLFormElement;
    export const passwordRepeatForm: HTMLFormElement = document.getElementById("newPasswordRepeatInput") as HTMLFormElement;
    export const submitBtn: HTMLFormElement = document.getElementById("changePasswordSubmitBtn") as HTMLFormElement;
    export const validationSpan: HTMLElement = document.getElementById("validationResult") as HTMLElement;
    export const closeBtn: HTMLElement = document.getElementById("closeChangePasswordModalBtn") as HTMLElement;

    function clearPasswordForm(): void {
        changePasswordModal.passwordForm.value = "";
        changePasswordModal.passwordRepeatForm.value = "";
        changePasswordModal.validationSpan.innerHTML = "";
    }

    passwordRepeatForm.addEventListener("input", () => {
        if (passwordRepeatForm.value == passwordForm.value) {
            validationSpan.innerHTML = "&#10003 Пароли совпадают";
            validationSpan.style.color = "green";
        } else {
            validationSpan.innerHTML = "&#10060 Пароли не совпадают";
            validationSpan.style.color = "red";
        }
    });

    closeBtn.addEventListener("click", clearPasswordForm);
    modal.addEventListener("click", (evt) => {
    if (evt.target == modal) {
        clearPasswordForm();
    }
});
};




function updateHTML(info: UserInfo): void {
    userDataFields.name.innerHTML = info.name ?? "Неизвестно";
    userDataFields.surname.innerHTML = info.surname ?? "Неизвестно";
    userDataFields.email.innerHTML = info.email ?? "Неизвестно";
    userDataFields.nickname.innerHTML = "@" + info.nickname ?? "@Неизвестно";
    userDataFields.aboutInfo.innerHTML = info.aboutInfo ?? "Неизвестно";
}

function renderProfile(id: number): void {
    sendRequest("GET", server + `/users/${id}`)
        .then(r => updateHTML(r))
        .catch(e => {
            throw e
        });
}

renderProfile(userId)
