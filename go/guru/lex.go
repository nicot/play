package guru

import (
	"strconv"
	"strings"
)

type TokenType int

const (
	Error TokenType = iota
	Delimiter
	TString
	TNumber
	TSymbol
)

type Token struct {
	Type  TokenType
	Value interface{}
}

const delims = " \t\n\r()" // TODO: unicode whitespace

func Lex(input string) []Token {
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
				tokens = append(tokens, Token{TString, s})
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
					tokens = append(tokens, Token{TSymbol, s})
				} else {
					tokens = append(tokens, Token{TNumber, n})
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
	case TString:
		return "\"" + t.Value.(string) + "\""
	case TNumber:
		return strconv.Itoa(t.Value.(int))
	case TSymbol:
		return t.Value.(string)
	case Error:
		return "Error"
	}
	panic("unreachable")
}
