import startServer from './server.js'
import preparePages from './parse.js'
import fontColors from './fontColors.js'

console.log(fontColors.get('black'), 'Preparing to parse pages.')

preparePages().then(() => {
    console.log(fontColors.get('green'), 'Parsed SPA pages succesfully.')
})

console.log(fontColors.get('black'), 'Starting server...')

startServer({ path: '../dist/', port: 3000 })
