import { HMRMSG } from './types'

const wsurl = "ws://localhost:6347"
const rootid = 'root-script'

const handleMessage = (main) =>  (ev: MessageEvent): void => {
	const data: HMRMSG = JSON.parse(ev.data)
	if (data.reload !== "hot") {
		throw new Error("Unexpected message.")
	}
	if (data.kind === "hard") {
		location.reload()
	}

	let script = document.createElement('script')
	script.text = data.code
	document.body.appendChild(script)
}


export default function(main: () => void): void {
	const global = self as any;
	if (global.hot) {
		main()
		return
	}
	global.hot = true
	let socket = new WebSocket(wsurl)
	socket.onmessage = handleMessage(main)
	main()
}
