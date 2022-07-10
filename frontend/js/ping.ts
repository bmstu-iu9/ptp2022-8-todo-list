// @ts-check

window.onload = function () {

    const pingBtn = <HTMLButtonElement>document.getElementsByClassName('ping-btn')[0]

    pingBtn.addEventListener('click', () => {
        const http = new XMLHttpRequest()
        

        http.open('GET', 'http://' + host + ':' + port + '/hello', /*async*/ true)
        try {
            http.send(null)
        } catch (exception) {
            // this is expected
        }
        http.onreadystatechange = function () {
            if (http.readyState == 4) {
                if (http.status != 200) {
                    alert(http.status + ": " + http.statusText)
                } else {
                    alert(http.responseText)
                }
            }
        }
    })
}
