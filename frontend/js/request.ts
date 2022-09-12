function sendRequest(method: string, url: string, body?: string) {
    const headers = {
        'Content-Type': 'application/json',
    }

    return fetch(url, {
        method: method,
        body: body,
        headers: headers,
    }).then((response) => {
        if (response.ok) {
            return response.json()
        }
        return response.json().then((error) => {
            const e = new Error('Пиво')
            e.message = error
            throw e
        })
    })
}
