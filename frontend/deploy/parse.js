// @ts-check

import fs from 'fs'

var workdir = '../dist/'
var regexp = new RegExp(/<div class="body.+">.+<\/div>/)
var files = ['todo.html', 'shop.html', 'profile_page.html', 'inventory.html']

if (!fs.existsSync(workdir + 'spa/views/')) {
    fs.mkdirSync(workdir + 'spa/views/', { recursive: true })
}

/**
 * Read from file and parse specific file
 * @param {string} filename 
 * @returns {Promise}
 */
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

/**
 * Parses all pages for SPA
 * @async
 * @returns {Promise<void>}
 */

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
