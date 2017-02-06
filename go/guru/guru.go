package guru

import (
	"fmt"
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
				tokens = append(tokens, Token{Symbol, s})
				pos = stop
			}
		} else {
			pos++
		}
		fmt.Println(string(char), tokens)
	}
	return tokens
}
