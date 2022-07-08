// @ts-check

window.onload = function () {

    const pingBtn = <HTMLButtonElement>document.getElementsByClassName('ping-btn')[0]

    pingBtn.addEventListener('click', () => {
        const started = new Date().getTime()
        const http = new XMLHttpRequest()
        const host = 'localhost' // Тут меняешь :)
        const port = 8080

        http.open('GET', 'http://' + host + ':' + port, /*async*/ true)
        http.onreadystatechange = function () {
            if (http.readyState == 4) {
                const ended = new Date().getTime()

                const milliseconds = ended - started

                window.alert(host + ' ping : ' + milliseconds)
            }
        }
        try {
            http.send(null)
        } catch (exception) {
            // this is expected
        }
    })
}