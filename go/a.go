package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func readIn() []int {
	buf, _ := ioutil.ReadAll(os.Stdin)
	s := string(buf)
	l := strings.Split(s, "\n")
	s = l[1]
	l = strings.Split(s, " ")
	ns := make([]int, len(l))
	for i := 0; i < len(l); i++ {
		ns[i], _ = strconv.Atoi(l[i])
	}
	return ns
}

func main() {
	// ns := readIn()
	ns := []int{1, 3, 2}
	l := 1
	r := 3
	result := solve(ns[l-1 : r])
	fmt.Println(result)
}

func solve(ns []int) int {
	fmt.Println(ns)
	   1
   	  1 1
	 1 2 1
	1 3 3 1
	return 0
}
