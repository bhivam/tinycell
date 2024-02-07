package main

import "strconv"

func literal_to_string(literal any) string {
	switch literal.(type) {
	case []byte:
		literal_string := literal.([]byte)
		return string(literal_string)
	case float64:
		literal_float := literal.(float64)
		return  strconv.FormatFloat(literal_float, 'f', 3, 64)
	default:
		return "UNKNOWN_TYPE"
	}
}

type TokenType int

const (
	EQUAL       TokenType = iota
	PLUS        TokenType = iota
	MINUS       TokenType = iota
	STAR        TokenType = iota
	SLASH       TokenType = iota
	LEFT_PAREN  TokenType = iota
	RIGHT_PAREN TokenType = iota

	COMMA    TokenType = iota
	NEW_LINE TokenType = iota

	NUMBER     TokenType = iota
	STRING     TokenType = iota
	IDENTIFIER TokenType = iota

	EOF      TokenType = iota
	NO_TOKEN TokenType = iota
)

func (tk_type TokenType) to_string() string {
	switch tk_type {
	case EQUAL:
		return "equal"
	case PLUS:
		return "plus"
	case MINUS:
		return "minus"
	case STAR:
		return "star"
	case SLASH:
		return "slash"
	case LEFT_PAREN:
		return "left parenthesis"
	case RIGHT_PAREN:
		return "right parenthesis"
	case COMMA:
		return "comma"
	case NEW_LINE:
		return "new_line"
	case NUMBER:
		return "number"
	case STRING:
		return "string"
	case IDENTIFIER:
		return "identifier"
    case EOF:
        return "end of file"
	default:
		return "UNKNOWN_TOKEN_TYPE"
	}
}

type Token struct {
	kind    TokenType
	lexeme  string
	literal any
}

func (t Token) to_string() string {
    literal := ""
    if t.literal != nil {
        literal = "|" + literal_to_string(t.literal)
    }   
	return t.kind.to_string() + "|" + t.lexeme + literal
}
