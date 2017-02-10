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

// Danger: recursive
func Flatten(node Node) (FlatExprs, *Value) {
	if node.Value != nil {
		return nil, node.Value
	}

	var prog FlatExprs
	var stmt []Value
	var ret *Value
	for _, v := range node.Children {
		f, v := Flatten(v)
		if f != nil {
			if len(f) > 1 {
				fmt.Println(f)
				prog = append(prog, f[0:]...)
			}
			binding := Value{SSA, 0}
			asn := append([]Value{{Symbol, "assign"}, binding}, f[0]...)
			prog = append(prog, asn)
			stmt = append(stmt, binding)
		} else {
			stmt = append(stmt, *v)
		}
	}

	prog = append(prog, stmt)
	return prog, ret
}
