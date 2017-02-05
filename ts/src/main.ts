import hotify from './hotify'
import start from './render'

function sayHi(): void {
	let x = 1
	console.log('the freaking coolest!')
}


function foo(bar?: number) {

}

foo(null);


function foo2(bar: string | undefined) {

}

function main(): void {
	sayHi()
	start()
}

hotify(main)
