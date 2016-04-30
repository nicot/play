package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	buf, _ := ioutil.ReadAll(os.Stdin)
	s := string(buf)
	l := strings.Split(s, "\n")
	s = l[1]
	l = strings.Split(s, " ")
	ns := make([]int, len(l))
	for i := 0; i < len(l); i++ {
		ns[i], _ = strconv.Atoi(l[i])
	}

	r := clouds(ns)
	fmt.Println(r)
}

func clouds(ns []int) int {
	count := 0
	for i := 0; i < len(ns)-1; {
		if i == len(ns)-2 {
			count += 1
			break
		}
		if ns[i+2] == 1 {
			i += 1
		} else {
			i += 2
		}
		count += 1
	}
	return count
}
