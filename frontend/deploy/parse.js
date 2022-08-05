// @ts-check

let fs = require('fs'),
    path = require('path')

var workdir = '../dist/'
var regexp = new RegExp(/<div class="body">.+<\/div>/)
var files = ['todo.html', 'shop.html', 'profile_page.html']

var readAndParse = (filename) => {
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
    })
}

files.forEach((file) => {
    readAndParse(file)
})

console.log('Parsed ' + files)