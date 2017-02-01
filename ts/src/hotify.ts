import * as browserify from 'browserify'

interface Options {

}

export default function(b: browserify.BrowserifyObject, opts: Options) {
	console.log(opts)
	console.log(b)
	return b
}
