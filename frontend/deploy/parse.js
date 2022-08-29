// @ts-check

import { rejects } from 'assert'
import fs from 'fs'
import { resolve } from 'path'

var workdir = '../dist/'
var regexp = new RegExp(/<div class="body.+">.+<\/div>/)
var files = ['todo.html', 'shop.html', 'profile_page.html', 'inventory.html']

if (!fs.existsSync(workdir + 'spa/views/')) {
    fs.mkdirSync(workdir + 'spa/views/', { recursive: true })
}

var readAndParse = (filename) => {
    return Promise.resolve(
        fs.readFile(workdir + 'html/' + filename, function (err, data) {
            if (err) {
                console.log(err.message)
            }
            let block = data.toString().match(regexp)[0]
            fs.writeFile(workdir + 'spa/views/' + filename, block, function (err) {
                if (err) {
                    throw err
                }
            })
        }),
    )
}

export default async function preparePages() {
    return /** @type {Promise<void>} */(new Promise((resolve, reject) => {
        files.forEach(async (file) => {
            readAndParse(file).then(() => {
                console.log('\x1b[34m', `Parsed ${file}`)
            })
        })
        resolve()
    }))
}
