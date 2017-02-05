import * as WS from 'websocket'
import * as http from 'http'
import { BrowserifyObject } from 'browserify'
import { HMRMSG } from '../src/types'
import { createWriteStream } from 'fs'

let port = 6347

export function dirname(path: string): string {
	return `${__dirname}/../${path}`
}

function handleRequest(request, response) {
	console.log((new Date()) + ' Received request for ' + request.url)
	response.writeHead(404)
	response.end()
}

function serve() {
	console.log((new Date()) + ' Server is listening on port ' + port)
}

function error(error) {
	console.error(error.toString())
}


export function bundle(b: BrowserifyObject) {
	let path = dirname('out/bundle.js');
	let outstream = createWriteStream(path);
	b.bundle()
		.on('error', error)
		.pipe(outstream);
}

class Emitter implements NodeJS.EventEmitter {
	addListener(event, ls) {
		return this
	}

	on(event, ls) {
		return this
	}

	once(event, ls) {
		return this
	}

	removeListener(event, ls) {
		return this
	}

	removeAllListeners(event) {
		return this
	}

	setMaxListeners(n) {
		return this
	}

	getMaxListeners() {
		return 0
	}

	listeners() {
		return []
	}

	emit(event, args) {
		return false
	}

	listenerCount() {
		return 0
	}

	prependListener(event, ls) {
		return this
	}

	prependOnceListener(event, ls) {
		return this
	}

	eventNames() {
		return []
	}
}

class WSPiper extends Emitter implements NodeJS.WritableStream {
	constructor(
		private conn: WS.connection,
	) {
		super()
	}

	writable = true
	code = ""

	write(buffer) {
		this.code += buffer
		return true
	}

	end(buffer?) {
		if (buffer) {
			this.write(buffer)
		}
		let value: HMRMSG = {
			kind: "soft",
			reload: "hot",
			code: this.code,
		}
		let message = JSON.stringify(value)
		this.conn.send(message)
	}
}

let reload = (conn: WS.connection, b: BrowserifyObject) => () => {
	let acc = new WSPiper(conn)
	let path = dirname('out/bundle.js');
	let outstream = createWriteStream(path);
	let stream = b.bundle()
	stream.on('error', error)
		.pipe(acc)
	stream.pipe(outstream)
}

let handleWS = (b: BrowserifyObject) => (request: WS.request) => {
	let connection = request.accept();
	console.log((new Date()) + ' Connection accepted.');
	b.on('update', reload(connection, b))
}

export default function (b: BrowserifyObject) {
	let httpServer = http.createServer(handleRequest)
	httpServer.listen(port, serve)
	let ws = new WS.server({ httpServer })
	ws.on('request', handleWS(b))
}
