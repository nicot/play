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
	ns := readIn()
	r := solve(ns)
	fmt.Println(r)
}

// 1 2 3 4 5 6 7
// 3 4 5
// q*2 + p simple version, many nodes
// less than 100 nodes and less than 100 edges
// 2
