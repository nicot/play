package guru

type Type int

const (
	TNumber Type = iota
	TString
	TSymbol
)

type Value struct {
	Type  Type
	Value interface{}
}

func ConvertValue(token Token) Value {
	switch token.Type {
	case Number:
		return Value{TNumber, token.Value.(int)}
	case String:
		return Value{TString, token.Value.(string)}
	case Symbol:
		return Value{TSymbol, token.Value.(string)}
	}
	panic("unreachable")
}

func Eval(node Node) Value {
	if node.Token != nil {
		return ConvertValue(*node.Token)
	}
	nodes := node.Children
	if len(nodes) < 1 {
		panic("Not enough nodes")
	}
	f := Eval(nodes[0])
	if f.Type != TSymbol {
		panic("Expected symbol in call position")
	}
	sym := f.Value.(string)

	// bad

	if sym == "+" {
		return Value{TNumber, Eval(nodes[1]).Value.(int) + Eval(nodes[2]).Value.(int)}
	}
	return Value{}
}
