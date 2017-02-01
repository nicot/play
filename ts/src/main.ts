// import hotify from 'hotify'

function sayHi(): void {
	let x = 1
	console.log('foobar!')
}

sayHi()

function foo(bar?: number) {

}

foo(null);


function foo2(bar: string | undefined) {

}

foo2();
