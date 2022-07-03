'use script'
var loginURL = "localhost";
var logoutURL = "/WebServices/LogOff";
var userAgent = navigator.userAgent.toLowerCase();
var firstLogIn = true;

function ping(host, port) {

    var started = new Date().getTime();

    var http = new XMLHttpRequest();

    http.open("GET", "http://" + host + ":" + port, /*async*/ true);
    http.onreadystatechange = function () {
        if (http.readyState == 4) {
            var ended = new Date().getTime();

            var milliseconds = ended - started;

            window.alert(host + " ping : " + milliseconds);
        }
    };
    try {
        http.send(null);
    } catch (exception) {
        // this is expected
    }

}

function ping1() {
    ping("localhost", 8080);
}

function cyrb53(str, seed = 0) {
    let h1 = 0xdeadbeef ^ seed, h2 = 0x41c6ce57 ^ seed;
    for (let i = 0, ch; i < str.length; i++) {
        ch = str.charCodeAt(i);
        h1 = Math.imul(h1 ^ ch, 2654435761);
        h2 = Math.imul(h2 ^ ch, 1597334677);
    }
    h1 = Math.imul(h1 ^ (h1>>>16), 2246822507) ^ Math.imul(h2 ^ (h2>>>13), 3266489909);
    h2 = Math.imul(h2 ^ (h2>>>16), 2246822507) ^ Math.imul(h1 ^ (h1>>>13), 3266489909);
    return 4294967296 * (2097151 & h2) + (h1>>>0);
};

var login = function () {
    var form = document.forms[0];
    var username = form.username.value;
    var pswd = form.password.value;
    var submitBtn = document.getElementById('submit');
    var xhr = new XMLHttpRequest();

    var json = JSON.stringify({
        name: username,
        password: cyrb53(password)
    });

    xhr.open('POST', 'https://localhost:8080', true)
    xhr.setRequestHeader('Content-type', 'application/json; charset=utf-8');

    xhr.onreadystatechange = {

    };

    // Отсылаем объект в формате JSON и с Content-Type application/json
    // Сервер должен уметь такой Content-Type принимать и раскодировать
    xhr.send(json);
}