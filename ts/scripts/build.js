let browserify = require('browserify')
let tsify = require('tsify')
let watchify = require('watchify')
let fs = require('fs')

function dirname(path) {
	return `${__dirname}/../${path}`
}

let b = browserify({
	entries: [dirname('src/main.ts')],
	cache: {},
	packageCache: {},
})

b.plugin(tsify, { noImplicitAny: true })
b.plugin(watchify)

function error(error) {
	console.error(error.toString())
}

function time(time) {
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
