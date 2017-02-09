package guru

import (
	"strconv"
	"strings"
)

type TokenType int

const (
	Delimiter TokenType = iota
	String
	Number
	Symbol
	Error
)

type Token struct {
	Type  TokenType
	Value interface{}
}

const delims = " \t\n\r()" // TODO: unicode whitespace

func Read(input string) []Token {
	var pos int
	var tokens []Token
	for pos < len(input) {
		char := input[pos]
		if char == '(' {
			tokens = append(tokens, Token{Delimiter, "open"})
			pos++
		} else if char == ')' {
			tokens = append(tokens, Token{Delimiter, "close"})
			pos++
		} else if char == '"' {
			pos++
			stop := strings.IndexByte(input[pos:], '"')
			if stop == -1 {
				tokens = append(tokens, Token{Error, nil})
				pos++
			} else {
				stop = pos + stop
				s := input[pos:stop]
				tokens = append(tokens, Token{String, s})
				pos = stop + 1
			}
		} else if strings.IndexRune(delims, rune(char)) == -1 {
			stop := strings.IndexAny(input[pos:], delims)
			if stop == -1 {
				tokens = append(tokens, Token{Error, nil})
				pos++
			} else {
				stop = pos + stop
				s := input[pos:stop]
				n, err := strconv.Atoi(s)
				if err != nil {
					tokens = append(tokens, Token{Symbol, s})
				} else {
					tokens = append(tokens, Token{Number, n})
				}
				pos = stop
			}
		} else {
			pos++
		}
	}
	return tokens
}

func (t Token) String() string {
	switch t.Type {
	case Delimiter:
		if t.Value == "open" {
			return "("
		} else if t.Value == "close" {
			return ")"
		}
	case String:
		return "\"" + t.Value.(string) + "\""
	case Number:
		return strconv.Itoa(t.Value.(int))
	case Symbol:
		return t.Value.(string)
	case Error:
		return "Error"
	}
	panic("unreachable")
}
