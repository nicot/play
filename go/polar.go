package main

import (
	"fmt"
	"math"
)

type pol struct {
	a     float64
	b     float64
	angle float64
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func angle(a float64, b float64) float64 {
	ang := math.Atan2(b, a)
	if ang < 0 {
		ang += 2 * math.Pi
	}
	return ang
}

func mag(p pol) float64 {
	ret := math.Pow(p.a, 2) + math.Pow(p.b, 2)
	ret = math.Sqrt(ret)
	return ret
}

var zero pol

func insert(slice []pol, x pol) []pol {
	zero = pol{0, 0, 0}
	for i := range slice {
		//fmt.Println(x, slice[i], slice)
		if slice[i] == zero {
			slice[i] = x
			return slice
		} else if x.angle < slice[i].angle {
			copy(slice[i+1:], slice[i:])
			slice[i] = x
			return slice
		} else if x == slice[i] && mag(x) < mag(slice[i]) {
			copy(slice[i+1:], slice[i:])
			slice[i] = x
			return slice
		}
	}
	return slice
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	check(err)
	slice := make([]pol, n)
	for i := 0; i < n; i++ {
		var a, b float64
		var p pol

		_, err := fmt.Scan(&a, &b)
		check(err)
		p.a, p.b, p.angle = a, b, angle(a, b)
		slice = insert(slice, p)
	}

	for _, e := range slice {
		a, b := e.a, e.b
		fmt.Println(int(a), int(b))
	}
}
