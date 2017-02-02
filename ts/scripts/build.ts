import * as browserify from 'browserify'
import * as tsify from 'tsify'
import * as watchify from 'watchify'
import listen, { dirname, bundle } from './server'

let b = browserify({
	entries: [dirname('src/main.ts')],
	cache: {},
	packageCache: {},
	debug: true,
})

b.plugin(tsify, { project: "../." })
b.plugin(watchify)

function time(time: number) {
	console.log(`Bundle took ${time}ms to build.`)
}

b.on('time', time)

listen(b)
bundle(b);
