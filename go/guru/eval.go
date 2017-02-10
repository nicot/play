package guru

import (
	"strconv"
)

type Type int

const (
	Number Type = iota
	String

	// TSymbol can be user defined or a builtin. It can be eager or lazy.
	Symbol
)

type Value struct {
	Type  Type
	Value interface{}
}

func TokenToValue(token Token) Value {
	switch token.Type {
	case TNumber:
		return Value{Number, token.Value.(int)}
	case TString:
		return Value{String, token.Value.(string)}
	case TSymbol:
		return Value{Symbol, token.Value.(string)}
	}
	panic("unreachable")
}

func (v Value) String() string {
	switch v.Type {
	case Number:
		return strconv.Itoa(v.Value.(int))
	case String:
		return "\"" + v.Value.(string) + "\""
	case Symbol:
		return v.Value.(string)
	}
	panic("unreachable")
}

func EvalRec(node Node) Value {
	if node.Value != nil {
		return *node.Value
	}
	nodes := node.Children
	if len(nodes) < 1 {
		panic("Not enough nodes")
	}
	f := EvalRec(nodes[0])
	if f.Type != Symbol {
		panic("Expected symbol in call position")
	}
	sym := f.Value.(string)

	// bad

	if sym == "+" {
		return Value{Number, EvalRec(nodes[1]).Value.(int) + EvalRec(nodes[2]).Value.(int)}
	}
	return Value{}
}

func Eval(node Node) Value {
	byteCode := Flatten(node)
	return EvalFlat(byteCode)
}

func EvalFlat(f []FlatExpr) Value {
	return Value{}
}
