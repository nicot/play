import hotify from './hotify'

function sayHi(): void {
	let x = 1
	console.log('foobara!')
}


function foo(bar?: number) {

}

foo(null);


function foo2(bar: string | undefined) {

}

function main(): void {
	sayHi();
}

hotify(main);
