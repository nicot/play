package main

import (
	"fmt"
	"strconv"
)

func readInt() int {
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		panic(err)
	}
	return n
}


func main() {
	n := readInt()
	for i := 0; i < n; i++ {
		shield := readInt()
		var instructions []byte
		fmt.Scanln(&instructions)
		s := solve(shield, instructions)
		fmt.Printf("Case #%d: %s\n", i + 1, s)
	}
}

func solve(shield int, instructions []byte) string {
	shotCount := 0
	for _, c := range instructions {
		if c == 'S' {
			shotCount++
		}
	}
	if shotCount > shield {
		return "IMPOSSIBLE"
	}

	i := 0
	for {
		d := calcDamage(instructions)
		if d <= shield {
			return strconv.Itoa(i)
		}
		instructions = shuffle(instructions)
		i++
	}
}

func calcDamage(instructions []byte) int {
	charge := 1
	damage := 0
	for _, c := range instructions {
		if c == 'C' {
			charge *= 2
		} else if c == 'S' {
			damage += charge
		}
	}
	return damage
}

func shuffle(instructions []byte) []byte {
	index := 0
	for i := 1; i < len(instructions); i++ {
		c := instructions[i]
		if c == 'S'  && instructions[i-1] == 'C' {
			index = i
		}
	}
	tmp := instructions[index]
	instructions[index] = instructions[index - 1]
	instructions[index - 1] = tmp
	return instructions
}