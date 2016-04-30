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
	diff, _ := strconv.Atoi(l)
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
	r := trip(ns, diff)
	fmt.Println(r)
}
