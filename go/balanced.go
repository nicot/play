package main

import "fmt"

func main() {
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		panic(err)
	}

	for i := 0; i < n; i ++ {
		var s string
		fmt.Scanln(&s)
		if (isBalanced(s)) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

var opening map[rune]struct{} = map[rune]struct{}{
	'(': struct{}{},
	'{': struct{}{},
	'[': struct{}{},
}

var closing map[rune]rune = map[rune]rune{
	')': '(',
	'}': '{',
	']': '[',
}

func isBalanced(s string) bool {
	stack := make([]rune, 0)

	// I'm not sure if c is a run or a byte. I think it is a rune.
	for _, c := range s {
		// if it's an opening bracket, add it to the stack
		if _, isOpening := opening[c]; isOpening {
			stack = append(stack, c)
			continue
		}
		// if it's a closing bracket, pop the top of the stack
		expected, isClosing := closing[c]
		if !isClosing {
			panic("found non-bracket")
		}
		if len(stack) == 0 {
			return false
		}
		top := stack[len(stack) - 1]
		if expected != top {
			// if the top of the stack is not equivalent, return false
			return false
		}
		stack = stack[0:len(stack) - 1]
	}

	return len(stack) == 0
}