package guru

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

func Flatten(node Node) FlatExprs {
	if node.Value != nil {
		panic("idk")
	}

	var stmts FlatExprs
	stmts = append(stmts, []Value{})
	i := 0
	for {
		v := node.Children[i]
		if v.Value != nil {
			stmts[0] = append(stmts[0], *v.Value)
		}
		i++
		if len(node.Children) == i {
			return stmts
		}
	}
}
