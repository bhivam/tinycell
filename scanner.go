package main

import (
	"strconv"
	"strings"
)

// TODO Make this scanner compatible with '.' in numbers
type Scanner struct {
	source        []byte
	tokens        []Token
	start         int
	current       int
	line          int
	equation_mode bool
}

func (sn *Scanner) scan_tokens(source []byte)  {
	sn.source = source
	sn.tokens = []Token{}
	sn.start = 0
	sn.current = 0
	sn.line = 1
    
	for !sn.is_at_end() {
		sn.scan_token()
	}

	sn.tokens = append(sn.tokens, Token{kind: EOF, lexeme: "", literal: nil})
}

func (sn *Scanner) scan_token() {
	c := sn.advance()
	switch c {
	// for new line and comma add string or number if last token was
	// also comma
	case ',':
		sn.add_token_nl(COMMA)
	case '\n':
		sn.line++
		sn.add_token_nl(NEW_LINE)
	case '=':
		lt_kind := sn.last_token().kind
		if lt_kind == NEW_LINE || lt_kind == COMMA {
			sn.add_token_nl(EQUAL)
			sn.equation()
		} else {
			sn.advance()
		}
	default:
		for sn.peek() != ',' && sn.peek() != '\n' {
			sn.advance()
		}
		literal := sn.source[sn.start:sn.current]
		for i := range literal {
			b := literal[i]
			if !is_num(b) {
				sn.add_token(STRING, literal)
				return
			}
		}
		x, err := strconv.ParseFloat(string(literal), 64)
		check(err)

		sn.add_token(NUMBER, x)
	}
}

func (sn *Scanner) equation() {
	alpha_over := false
	has_alpha := false
	for sn.peek() != ',' && sn.peek() != '\n' {
		c := sn.peek()

		if (c == '+' || c == '-' || c == '*' ||
			c == '/' || c == '(' || c == ')') &&
			sn.start < sn.current {
			if has_alpha {
				sn.add_token_nl(IDENTIFIER)
			} else {
				x, err := strconv.
					ParseFloat(string(sn.source[sn.start:sn.current]), 64)
				check(err)
				sn.add_token(NUMBER, x)
			}
			alpha_over = false
			has_alpha = false
		}

		c = sn.advance()

		switch c {
		case '+':
			sn.add_token_nl(PLUS)
		case '-':
			sn.add_token_nl(MINUS)
		case '*':
			sn.add_token_nl(STAR)
		case '/':
			sn.add_token_nl(SLASH)
		case '(':
			sn.add_token_nl(RIGHT_PAREN)
		case ')':
			sn.add_token_nl(LEFT_PAREN)
		default:
			if is_num(c) {
				alpha_over = true
			} else if is_alpha(c) {
				if alpha_over {
					report_error("Invalid literal.", sn.line)
				}
				has_alpha = true
			} else {
				report_error("Unexpected character.", sn.line)
			}
		}
	}
}

// HELPER FUNCTIONS
func is_alpha(c byte) bool {
	return (c < 91 && c > 64) || (c < 123 && c > 96)
}

func is_num(c byte) bool {
	return (c < 58 && c > 47) || c == '.'
}

func (sn *Scanner) is_at_end() bool {
	return sn.current >= len(sn.source)
}

func (sn *Scanner) advance() byte {
    sn.current += 1
	return sn.source[sn.current-1]
}

func (sn *Scanner) add_token(kind TokenType, literal any) {
	sn.tokens = append(sn.tokens, Token{
		kind:    kind,
		lexeme:  strings.Clone(string(sn.source[sn.start:sn.current])),
		literal: literal,
	})
	sn.start = sn.current
}

func (sn *Scanner) add_token_nl(kind TokenType) {
	sn.add_token(kind, nil)
}

func (sn *Scanner) peek() byte {
	if sn.is_at_end() {
		return 0
	}
	return sn.source[sn.current]
}

func (sn *Scanner) last_token() Token {
	n := len(sn.tokens)
	if n == 0 {
		return Token{kind: NO_TOKEN}
	}
	return sn.tokens[n-1]
}
