package guru

type FlatExprs [][]Value

func (f FlatExprs) String() string {
	s := "["
	for _, l := range f {
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

	for i, v := range node.Children {
		if v.Value != nil {

		}
	}

	return FlatExprs{}
}
