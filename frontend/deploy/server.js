'use strict'

import fs from 'fs'
import path from 'path'
import http from 'http'
import fontColors from './fontColors.js'

const cache = {}

let reroutePaths = new Map([
    ['/html/spa.html', '/app'],
    ['/html/login.html', '/login'],
    ['', ''],
])
let pathsReroute = new Map([
    ['/app', '/html/spa.html'],
    ['/login', '/html/login.html'],
    ['/', 'index.html'],
    ['', ''],
])
if (process.env['ENV_MODE'] === 'PROD') {
    console.log(fontColors.get('black'),'PROD MODE, index.html set to start_page.html.')
    pathsReroute.set('/', '/html/start_page.html')
} else {
    console.log(fontColors.get('black'),'DEV MODE, index.html not modified.')
}

/**
 * lookup content type
 * infer from the extension
 * no extension would resolve in "text/plain"
 * @param {string} fileName
 */
function lookupContentType(fileName) {
    const ext = fileName.toLowerCase().split('.').slice(1).pop()
    switch (ext) {
        case 'txt':
            return 'text/plain'
        case 'js':
            return 'text/javascript'
        case 'css':
            return 'text/css'
        case 'pdf':
            return 'application/pdf'
        case 'jpg':
        case 'jpeg':
            return 'image/jpeg'
        case 'mp4':
            return 'video/mp4'
        case 'svg':
            return 'image/svg+xml'
        case 'webp':
            return 'image/webp'
        default:
            return ''
    }
}

/**
 * plain 404 response
 * @param {http.ServerResponse} res
 */
function send404(res) {
    res.writeHead(404, { 'Content-Type': 'text/plain' })
    res.write('Error 404: resource not found.')
    res.end()
}

/**
 * sending file response
 * @param {http.ServerResponse} res
 * @param {string} filePath
 */
function sendFile(res, filePath, fileContents) {
    res.writeHead(200, { 'Content-Type': lookupContentType(path.basename(filePath)) })
    res.end(fileContents)
}

/**
 * plain 302 response
 * @param {http.ServerResponse} res
 * @param {string} url
 */
function send302(res, url) {
    res.writeHead(302, { location: url })
    res.end()
}

/**
 * serve static content
 * @param {http.ServerResponse} res
 * @param {{}} cache
 * @param {string} absPath
 * @returns {Promise}
 */
function serveStatic(res, cache, absPath) {
    return new Promise((resolve, reject) => {
        let reqCode
        // use cache if there is any
        if (cache[absPath]) {
            sendFile(res, absPath, cache[absPath])
            reqCode = '200'
            resolve(reqCode)
        } else {
            fs.exists(absPath, function (fileExists) {
                // attempt to read the resource only if it exist
                if (fileExists) {
                    fs.readFile(absPath, function (err, data) {
                        // not able to read the resource
                        if (err) {
                            reqCode = '403'
                            send404(res)
                            resolve(reqCode)
                        } else {
                            cache[absPath] = data
                            reqCode = '200'
                            sendFile(res, absPath, data)
                            resolve(reqCode)
                        }
                    })
                } else {
                    // resource does not exist
                    reqCode = '404'
                    send404(res)
                    resolve(reqCode)
                }
            })
        }
    })
}

/**
 *  Starts server.
 * @param {{path: string, port: number}} spec 
 * @returns 
 */
export default function startServer(spec) {
    let { path, port } = spec

    // create server object
    var server = http.createServer(function (req, res) {
        let filePath
        let handleDefault = () => {
            filePath = path + req.url.substring(1)
            serveStatic(res, cache, filePath).then((code) => {
                console.log(
                    fontColors.get('green'),
                    'GET:',
                    fontColors.get('blue'),
                    path,
                    ' ',
                    req.url,
                    ' ',
                    code === '404' ? fontColors.get('red') : fontColors.get('green'),
                    code,
                )
            })
        }
        if (reroutePaths.has(req.url)) {
            let reroute = reroutePaths.get(req.url)
            send302(res, reroute)
            console.log(fontColors.get('green'), 'GET:', fontColors.get('blue'), req.url, ' ', fontColors.get('yellow'), '302 -->', reroutePaths.get(req.url))
        } else if (pathsReroute.has(req.url)) {
            filePath = path + pathsReroute.get(req.url)
            serveStatic(res, cache, filePath).then((code) => {
                console.log(
                    fontColors.get('green'),
                    'GET:',
                    fontColors.get('blue'),
                    path,
                    ' ',
                    pathsReroute.get(req.url),
                    ' ',
                    code === '404' ? fontColors.get('red') : fontColors.get('green'),
                    code,
                )
            })
        } else {
            handleDefault()
        }
    })

    server.listen(port, function () {
        console.log('Server listening on port: ' + port)
    })
    return server
}
