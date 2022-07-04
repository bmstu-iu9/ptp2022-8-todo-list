'use script'
var loginURL = 'localhost'
var logoutURL = '/WebServices/LogOff'
var userAgent = navigator.userAgent.toLowerCase()
var firstLogIn = true

function ping(host, port) {
    var started = new Date().getTime()

    var http = new XMLHttpRequest()

    http.open('GET', 'http://' + host + ':' + port, /*async*/ true)
    http.onreadystatechange = function () {
        if (http.readyState == 4) {
            var ended = new Date().getTime()

            var milliseconds = ended - started

            window.alert(host + ' ping : ' + milliseconds)
        }
    }
    try {
        http.send(null)
    } catch (exception) {
        // this is expected
    }
}

function ping1() {
    ping('localhost', 8080)
}

async function sha256(message) {
    // encode as UTF-8
    const msgBuffer = new TextEncoder('utf-8').encode(message)

    // hash the message
    const hashBuffer = await crypto.subtle.digest('SHA-256', msgBuffer)

    // convert ArrayBuffer to Array
    const hashArray = Array.from(new Uint8Array(hashBuffer))

    // convert bytes to hex string
    const hashHex = hashArray.map((b) => ('00' + b.toString(16)).slice(-2)).join('')
    console.log(hashHex)
    return hashHex
}

var login = function () {
    var form = document.forms[0]
    var username = form.username.value
    var pswd = form.password.value
    var submitBtn = document.getElementById('submit')
    var xhr = new XMLHttpRequest()

    var json = JSON.stringify({
        name: username,
        password: sha256(password),
    })

    xhr.open('POST', 'https://localhost:8080', true)
    xhr.setRequestHeader('Content-type', 'application/json; charset=utf-8')

    xhr.onreadystatechange = {}

    // Отсылаем объект в формате JSON и с Content-Type application/json
    // Сервер должен уметь такой Content-Type принимать и раскодировать
    xhr.send(json)
}
