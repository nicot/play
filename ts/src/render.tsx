import * as React from 'react'
import * as ReactDOM from 'react-dom'
import * as InfiniteList from 'infinite-list'

let container = {
	width: "100%",
	height: 200,
	borderRadius: 3,
	border: "gray",
	borderStyle: "solid",
}

class Foo extends React.Component<{}, {}> {
	foo = (el) => {
		new InfiniteList({
			initialPage: {
				itemsCount: 100000,
			},
			itemHeightGetter: (index) => {
				return 20
			},
			itemRenderer: (index, el) => {
				ReactDOM.render(<div>Hello there number: {index}</div>, el)
			}
		}).attach(el)
	}

	render() {
		return <div style={container}>
			hello there willimena!
			<div style={{height: 500}} ref={this.foo} />
		</div>
	}
}

export default function () {
	let el = document.getElementById('document-root')
	ReactDOM.render(<Foo />, el)
}
