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
        userDataFields.nickname.innerHTML = "@"+ nicknameForm.value;
        userDataFields.name.innerHTML = nameForm.value;
        userDataFields.surname.innerHTML = surnameForm.value;
        userDataFields.aboutInfo.innerHTML = aboutInfoForm.value;
        localStorage.setItem("userInfo", JSON.stringify({
            nickname: nicknameForm.value,
            name: nameForm.value,
            surname: surnameForm.value,
            aboutInfo: aboutInfoForm.value,
    
        }));
        /*
            Тут, наверное, должен отправлятьcя запрос на сервер
        */   
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
}


function updateHTML(user: User): void {
    userDataFields.name.innerHTML = user.name ?? "Неизвестно";
    userDataFields.surname.innerHTML = user.surname ?? "Неизвестно";
    userDataFields.email.innerHTML = user.email ?? "Неизвестно";
    userDataFields.nickname.innerHTML = "@" + user.nickname ?? "@Неизвестно";
    userDataFields.aboutInfo.innerHTML = user.aboutInfo ?? "Неизвестно";
}



function renderProfile(id: number): void {
    sendRequest("GET", server + `/users/${id}`)
        .then(r => {
            updateHTML(r);
            renderHero("mainProfileCard");
        })
        .catch(e => {
            throw e
        });
}

renderProfile(userId);
sendRequest("PATCH", server + `/users/3`, JSON.stringify(
    {
    Items: [
        {
          ItemId: 1,
          ItemName: "Кожаный шлем",
          Description: "",
          ImageSrc: "armorHelmet1.webp",
          ImageForHero: "https://psv4.userapi.com/c235031/u179231152/docs/d4/f2cf1861a0ce/Kozhany_shlem_JE4_BE3.png?extra=A4xAyRgIKqcX-owdpzRfegyMNT08Qg8CyRcEQrB3j5uHhwyY9DzUtcR8J-z5TAbY_oJUlekCLMIE-IkBcFdhJUEYISY6F1qR4mLEgTEUXjfnXwEV6Htcvtn8LClV7QgWTRUThVgj5q6Zlger8BgeQJIdJQ",
          Price: 66,
          Category: "helmet",
          Rarity: "common",
          ItemState: "equipped"
        },
        {
            ItemId: 2,
            ItemName: "Железный нагрудник",
            Description: "",
            ImageSrc: "armorHelmet1.webp",
            ImageForHero: "https://psv4.userapi.com/c536236/u179231152/docs/d16/08a915a12893/Nagrudnik.png?extra=aWuQQXKwwdxQ05XW5JQNIw7xkETXrY5fYzrXgck5rqpuwy8NCtlEYhAjvi5zB0KzB6n2y4Mlchfpin0rGtOC-fFLpYHxtWQEQaerBoXOnkEHGqxqcA9XfpLBfu9ApsBOrhDj97Tp8-308jhjsARQ6z5G-A",
            Price: 66,
            Category: "chestplate",
            Rarity: "common",
            ItemState: "equipped"
          },
          {
            ItemId: 3,
            ItemName: "Алмазные поножи",
            Description: "",
            ImageSrc: "armorHelmet1.webp",
            ImageForHero: "https://psv4.userapi.com/c236331/u179231152/docs/d3/633d544801f6/ponozhi.png?extra=SQigaDT9BeZDP71GZvVts51jGBl6lySOt8rDyhTSyeHeH9ID11qUMO2wQoRVPkQsxMMaoTDbG__bE_vr2dMGf3iJnfgPEVxQwYMsDi4BedBZ6rAvkNDnmlgMsPlauCgUnWsz4stuJkUrnRo-TBDnW_7i_g",
            Price: 66,
            Category: "leggings",
            Rarity: "common",
            ItemState: "equipped"
          },
          {
            ItemId: 1,
            ItemName: "Золотые ботинки",
            Description: "",
            ImageSrc: "armorHelmet1.webp",
            ImageForHero: "https://psv4.userapi.com/c237131/u179231152/docs/d25/3908b0d82e16/botinki.png?extra=V2fkpsFSxl7aWlmQpxSuDv5HQWOMOP2ACueqPZ2_d6k0wQpD5yGmj3W2t-I2UuULH3XRoFbjuBxFDw56KITeutYG24962m6h-WL7UJRTFWqmQr6tWM9IgKCmeLSi1ejvEbsF6d59kqWQYy9Zg1AIhfiYMA",
            Price: 66,
            Category: "boots",
            Rarity: "common",
            ItemState: "inventoried"
          },
    ]}
));
