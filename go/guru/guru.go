package guru

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

func Read(input string) []Token {
	var tokens []Token
	var context string
	var num string
	var inString bool
	for _, v := range input {
		var t Token

		if inString {
			if v == '"' {
				t = Token{String, context}
				inString = false
			} else {
				context += string(v)
			}
		} else if v == '(' {
			t = Token{Delimiter, "open"}
		} else if v == ')' {
			t = Token{Delimiter, "close"}
		} else if v == ' ' {

		} else if v > '0' && v <= '9' {
			num += string(v)
			continue
		} else if v == '"' {
			inString = true
		} else {
			t = Symbol{Error, nil}
		}

		if num != "" {
			c := Token{Number, num}
			tokens = append(tokens, c)
			num = ""
		}
		if t != (Token{}) {
			tokens = append(tokens, t)
		}
	}
	return tokens
}
