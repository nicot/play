package guru

type FlatExpr struct {
	Operation Value
	Arguments []Value
}

func (f FlatExpr) String() string {
	s := f.Operation.String()
	for _, a := range f.Arguments {
		s += " " + a.String()
	}
	return s
}

func (f FlatExpr) Equals(f1 FlatExpr) bool {
	if f.Operation != f1.Operation {
		return false
	}

	if len(f.Arguments) != len(f1.Arguments) {
		return false
	}

	t := true
	for i := range f.Arguments {
		t = t && f.Arguments[i] == f1.Arguments[i]
	}

	return t
}

type ByteCode []FlatExpr

func (b ByteCode) String() string {
	s := "["
	for _, f := range b {
		s += "\n" + f.String()
	}
	return s + "\n]"
}

func (b ByteCode) Equals(b2 ByteCode) bool {
	if len(b) != len(b2) {
		return false
	}

	t := true
	for i := range b {
		t = t && b[i].Equals(b2[i])
	}

	return t
}

func Flatten(node Node) ByteCode {
	return nil
}
