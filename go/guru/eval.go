package guru

import (
	"strconv"
)

type Type int

const (
	TNumber Type = iota
	TString

	// TSymbol can be user defined or a builtin. It can be eager or lazy.
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

func (v Value) String() string {
	switch v.Type {
	case TNumber:
		return strconv.Itoa(v.Value.(int))
	case TString:
		return "\"" + v.Value.(string) + "\""
	case TSymbol:
		return v.Value.(string)
	}
	panic("unreachable")
}

func EvalRec(node Node) Value {
	if node.Token != nil {
		return ConvertValue(*node.Token)
	}
	nodes := node.Children
	if len(nodes) < 1 {
		panic("Not enough nodes")
	}
	f := EvalRec(nodes[0])
	if f.Type != TSymbol {
		panic("Expected symbol in call position")
	}
	sym := f.Value.(string)

	// bad

	if sym == "+" {
		return Value{TNumber, EvalRec(nodes[1]).Value.(int) + EvalRec(nodes[2]).Value.(int)}
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
