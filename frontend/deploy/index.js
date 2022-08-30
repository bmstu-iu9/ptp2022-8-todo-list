import startServer from './server.js'
import preparePages from './parse.js'

console.log('\x1b[30m', 'Preparing to parse pages.')

preparePages().then(() => {
    console.log('\x1b[32m', 'Parsed SPA pages succesfully.')
})

console.log('\x1b[30m', 'Starting server...')

startServer({ path: '../dist/', port: 3000 })
