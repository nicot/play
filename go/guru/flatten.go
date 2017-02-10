package guru

import (
	"fmt"
)

type FlatExprs [][]Value

func (f FlatExprs) String() string {
	s := "["
	for _, l := range f {
		s += "\n"
		for _, v := range l {
			s += v.String() + " "
		}
	}
	return s + "\n]"
}

func (f FlatExprs) Equals(f1 FlatExprs) bool {
	if len(f) != len(f1) {
		return false
	}

	t := true
	for i := range f {
		if len(f[i]) != len(f1[i]) {
			return false
		}
		for j := range f[i] {
			if f[i][j] != f1[i][j] {
				return false
			}
		}
	}

	return t
}

const qsize int = 2 << 20

// Danger: recursive
func Flatten(node Node) FlatExprs {
	if node.Value != nil {
		return FlatExprs{{*node.Value}}
	}

	q := make(chan Node, qsize)
	for _, n := range node.Children {
		q <- n
	}

	var prog FlatExprs
	var ssn int
	var stmt []Value
	for n := range q {
		fmt.Println(n)
		if n.Value != nil {
			stmt = append(stmt, *n.Value)
		} else {
			binding := Value{SSA, ssn}
			for _, n := range node.Children {
				q <- n
			}
			assn := []Value{{Symbol, "bind"}, binding}
			prog = append(prog, assn)
			stmt = append(stmt, binding)
		}
	}

	return prog
}
