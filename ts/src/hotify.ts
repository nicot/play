const wsurl = "ws://localhost:6347"

export default function(main: () => void): void {
	let socket = new WebSocket(wsurl);
	main()
}
