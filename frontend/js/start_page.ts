function changeTheme() {
    const changeDark = document.querySelectorAll(".bg-dark");
    const changeLight = document.querySelectorAll(".bg-light");
    var navLink = document.querySelectorAll(".nav-link");
    var navLinkw = document.querySelectorAll(".nav-linkw");
    var svg = document.querySelector(".sun-moon");
    var active = document.querySelector(".active");
    var activew = document.querySelector(".activew");
    var body = document.querySelector("body");

    changeDark.forEach(object => {
        object.classList.toggle('bg-light');
        object.classList.remove('bg-dark');
    });
    (<HTMLElement>svg).style.color = "rgba(0, 0, 0, 0.5)";
    navLink.forEach(link => {
        link.classList.toggle('nav-linkw');
        link.classList.remove('nav-link');
    });
    changeLight.forEach(object => {
        object.classList.toggle('bg-dark');
        object.classList.remove('bg-light');
    });
    (<HTMLElement>svg).style.color = "rgba(255, 255, 255, 0.5)";
    navLinkw.forEach(link => {
        link.classList.toggle('nav-link');
        link.classList.remove('nav-linkw');
    });

    if (body!.classList.contains('text-white')) {
        body!.classList.remove('text-white');
        (<HTMLElement>svg).style.color = "rgba(0, 0, 0, 0.5)";
        active!.classList.toggle('activew');
        active!.classList.remove('active');
    } else {
        body!.classList.toggle('text-white');
        (<HTMLElement>svg).style.color = "rgba(255, 255, 255, 0.5)";
        activew!.classList.toggle('active');
        activew!.classList.remove('activew');
    }

}