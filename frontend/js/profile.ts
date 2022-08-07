type User = {
    nickname: string,
    name: string,
    email: string,
    surname: string,
    aboutInfo: string,
    level: number,
    health: number,
    experience: number,
    equipment: {
        helmut: number | null,
        chestplate: number | null,
        leggings: number | null,
        boots: number | null,
    }

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
            .then(r => updateHTML(r))
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



sendRequest("GET", server + "/users/3")
        .then(r => {
            renderHero("mainProfileCard", r.Items)
        })
        .catch(e => {
            throw e;
        })

sendRequest("GET", server + `/users/${userId}`)
    .then(r => {
        updateHTML(r);
    })
    .catch(e => {
    throw e
    });


// sendRequest("PATCH", server + "/users/3", JSON.stringify(
//     {Items: [
//     {
//       "id": 1,
//       "name": "Кожаный шлем",
//       "description": "https://i.ibb.co/mHYz2J6/0-1.png",
//       "imageSrc": "armorHelmet1.webp",
//       "imageForHero": "",
//       "price": 66,
//       "category": "helmet",
//       "rarity": "common",
//       "state": "inventoried"
//     },
//     {
//       "id": 2,
//       "name": "Кольчуга",
//       "description": "",
//       "imageForHero": "https://i.ibb.co/mHYz2J6/0-1.png",
//       "imageSrc": "armorHelmet2.webp",
//       "price": 72,
//       "category": "helmet",
//       "rarity": "rare",
//       "state": "inventoried"
//     },
//     {
//       "id": 3,
//       "name": "Золотой шлем",
//       "description": "",
//       "imageForHero": "https://i.ibb.co/mHYz2J6/0-1.png",
//       "imageSrc": "armorHelmet3.webp",
//       "price": 78,
//       "category": "helmet",
//       "rarity": "epic",
//       "state": "inventoried"
//     },
//     {
//       "id": 4,
//       "name": "Незерский шлем",
//       "description": "",
//       "imageForHero": "https://i.ibb.co/tJPjpzx/0-1.png",
//       "imageSrc": "armorHelmet4.webp",
//       "price": 84,
//       "category": "helmet",
//       "rarity": "legendary",
//       "state": "equipped"
//     },
//     {
//       "id": 5,
//       "name": "Кожаный нагрудник",
//       "description": "",
//       "imageForHero": "https://i.ibb.co/hRf5YWk/1.png", 
//       "imageSrc": "armorChest1.webp",
//       "price": 90,
//       "category": "chest",
//       "rarity": "common",
//       "state": "inventoried"
//     },
//     {
//       "id": 6,
//       "name": "Кольчужный нагрудник",
//       "description": "",
//       "imageForHero": "https://i.ibb.co/hRf5YWk/1.png",
//       "imageSrc": "armorChest2.webp",
//       "price": 96,
//       "category": "chest",
//       "rarity": "rare",
//       "state": "inventoried"
//     },
//     {
//       "id": 7,
//       "name": "Золотые доспехи",
//       "description": "",
//       "imageForHero": "https://i.ibb.co/5nCNyWx/1.png",
//       "imageSrc": "armorChest3.webp",
//       "price": 102,
//       "category": "chest",
//       "rarity": "epic",
//       "state": "equipped"
//     },
//     {
//       "id": 8,
//       "name": "Незерская броня",
//       "description": "",
//       "imageForHero": "https://i.ibb.co/hRf5YWk/1.png",
//       "imageSrc": "armorChest4.webp",
//       "price": 108,
//       "category": "chest",
//       "rarity": "legendary",
//       "state": "inventoried"
//     },
//     {
//       "id": 9,
//       "name": "Кожаные штанишки",
//       "description": "",
//       "imageForHero": "https://i.ibb.co/bvmyfN5/2.png",
//       "imageSrc": "armorBeing1.webp",
//       "price": 114,
//       "category": "leggins",
//       "rarity": "common",
//       "state": "equipped"
//     },
//     {
//       "id": 10,
//       "name": "Кованые штаны",
//       "description": "",
//       "imageForHero": "https://i.ibb.co/hRf5YWk/1.png",
//       "imageSrc": "armorBeing2.webp",
//       "price": 120,
//       "category": "leggins",
//       "rarity": "rare",
//       "state": "inventoried"
//     },
//     {
//       "id": 11,
//       "name": "Золотые штаны",
//       "description": "",
//       "imageForHero": "https://i.ibb.co/hRf5YWk/1.png",
//       "imageSrc": "armorBeing3.webp",
//       "price": 126,
//       "category": "leggins",
//       "rarity": "epic",
//       "state": "inventoried"
//     },
//     {
//       "id": 12,
//       "name": "Алмазные штаны",
//       "description": "",
//       "imageForHero": "https://i.ibb.co/hRf5YWk/1.png",
//       "imageSrc": "armorBeing4.webp",
//       "price": 132,
//       "category": "leggins",
//       "rarity": "legendary",
//       "state": "inventoried"
//     }
//   ]}));
