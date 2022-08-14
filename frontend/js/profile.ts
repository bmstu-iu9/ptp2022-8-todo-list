type User = {
    nickname: string,
    name: string,
    email: string,
    surname: string,
    aboutInfo: string,
    level: number,
    health: number,
    experience: number,
};
let userId: number = 3;
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
    export const switchThemeBtn: HTMLElement = document.getElementById("theme-switch-checkbox") as HTMLInputElement;
    if (localStorage.getItem("theme") == "light") {
        switchThemeBtn.checked = true;
    } else {
        switchThemeBtn.cheked = false;
    }
    switchThemeBtn.addEventListener("click", () => console.log(switchThemeBtn.checked));
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
            aboutInfo: aboutInfoForm.value,}))
            .then(r => {
                updateHTML(r);
                if ((switchThemeBtn.checked == true) && (localStorage.getItem("theme") == "dark")) {
                    localStorage.setItem("theme", "light");
                    switchProfileTheme();
                } else if ((switchThemeBtn.checked == false) && (localStorage.getItem("theme") == "light")) {
                    localStorage.setItem("theme", "dark");
                    switchProfileTheme();
                }
            })
            .catch(e => {
                throw e;
            })
        });
};

namespace changePasswordModal {
    const modal: HTMLElement = document.getElementById("changePasswordModal") as HTMLElement;
    const passwordForm: HTMLFormElement = document.getElementById("newPasswordInput") as HTMLFormElement;
    const passwordRepeatForm: HTMLFormElement = document.getElementById("newPasswordRepeatInput") as HTMLFormElement;
    const submitBtn: HTMLFormElement = document.getElementById("changePasswordSubmitBtn") as HTMLFormElement;
    const validationSpan: HTMLElement = document.getElementById("validationResult") as HTMLElement;
    const closeBtn: HTMLElement = document.getElementById("closeChangePasswordModalBtn") as HTMLElement;
    const oldPasswordForm: HTMLElement = document.getElementById("oldPasswordInput") as HTMLElement;

    function clearPasswordForm(): void {
        oldPasswordForm.value = "";
        passwordForm.value = "";
        passwordRepeatForm.value = "";
        validationSpan.innerHTML = "";
    } 

    function setSuccesValidationStyles(msg: string): void {
            validationSpan.innerHTML = `&#10003 ${msg}`;
            validationSpan.style.color = "green";
    }

    function setFailValidationStyles(msg: string): void {
            validationSpan.innerHTML = `&#10060 ${msg}`;
            validationSpan.style.color = "red";
    }

    passwordForm.addEventListener("input", () => {
        if (passwordRepeatForm.value != "") {
            if (passwordRepeatForm.value === passwordForm.value) {
                setSuccesValidationStyles("Пароли совпадают");
            } else {
                setFailValidationStyles("Пароли не совпадают");
            }
        }
    })

    passwordRepeatForm.addEventListener("input", () => {
        if (passwordRepeatForm.value == passwordForm.value) {
            setSuccesValidationStyles("Пароли совпадают");
        } else {
            setFailValidationStyles("Пароли не совпадают");
        }
    });

    submitBtn.addEventListener("click", () => {
        if (passwordForm.value === passwordRepeatForm.value) {
            if (passwordForm.value.length < 8) {
                setFailValidationStyles("Длина нового пароля меньше 8 символов");
            } else {
                /* тут должен отправляться запрос на сервер*/
            }
        }
    })

    closeBtn.addEventListener("click", clearPasswordForm);
    modal.addEventListener("click", (evt) => {
    if (evt.target == modal) {
        clearPasswordForm();
    }
    });
}


function updateHTML(user: User): void {
    userDataFields.name.innerHTML = user.name ?? "Неизвестно";
    userDataFields.surname.innerHTML = user.surname ?? "Неизвестно";
    userDataFields.email.innerHTML = user.email ?? "Неизвестно";
    userDataFields.nickname.innerHTML = "@" + user.nickname ?? "@Неизвестно";
    userDataFields.aboutInfo.innerHTML = user.aboutInfo ?? "Неизвестно";
}


function switchElementTheme(el: HTMLElement): void {
    el.classList.toggle("text-white");
    el.classList.toggle("bg-dark");
    el.classList.toggle("border-light")
}

function switchProfileTheme() {
    document.body.classList.toggle("bg-dark");
    document.querySelectorAll(".container").forEach(container => switchElementTheme(container));
    document.querySelectorAll(".card").forEach(card => switchElementTheme(card));
    document.querySelectorAll(".list-group-item").forEach(li => switchElementTheme(li));
}

window.addEventListener("load", () => {
    if (localStorage.getItem("theme") == "dark") {
        switchProfileTheme();
    }
})

sendRequest("GET", server + `/users/${userId}`)
    .then(r => {
        updateHTML(r);
        let e: Equipment = getEquipment(r.Items);
        renderHero("user-img-card", e);
        let petCard: HTMLElement = document.getElementById("pet-container");
            if (e.pet) {
                petCard.appendChild(createEquipHtml(e.pet));
            }
        

    })
    .catch(e => {
    throw e
    });
