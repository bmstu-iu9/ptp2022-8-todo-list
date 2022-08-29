import startServer from './server.js'
import preparePages from './parse.js'

console.log('\x1b[30m', 'Preparing to parse pages.')

try {
    preparePages().then(() => {
        console.log('\x1b[32m', 'Parsed SPA pages succesfully.')
    })
} catch (error) {
    console.error(error)
    throw new error()
}

startServer({ path: '../dist/', port: 3000 })
