import * as browserify from 'browserify'
import * as tsify from 'tsify'
import * as watchify from 'watchify'
import * as fs from 'fs'

function dirname(path: string): string {
	return `${__dirname}/../${path}`
}

let b = browserify({
	entries: [dirname('src/main.ts')],
	cache: {},
	packageCache: {},
	debug: true,
})

b.plugin(tsify, { project: "../." })
b.plugin(watchify)

function error(error) {
	console.error(error.toString())
}

function time(time: number) {
	console.log(`Bundle took ${time}ms to build.`)
}

function bundle() {
	let path = dirname('out/bundle.js');
	let outstream = fs.createWriteStream(path);
	b.bundle()
		.on('error', error)
		.pipe(outstream);
}

b.on('update', bundle)
b.on('time', time)

bundle()
