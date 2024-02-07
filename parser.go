package main

import (
	"errors"
)

type Parser struct {
	tokens  []Token
	cells   []Cell
	current int
}

/*

cell       -> expression ("\n" | ",")
expression -> (("=" term) | literal)
term       -> factor ( ( "-" | "+" ) factor )*
factor     -> unary ( ( "/" | "*" ) unary )*
unary      -> "-" unary | primary
primary    -> literal | "(" term ")"
literal    -> NUMBER | STRING

*/

func (ps *Parser) parse(tokens []Token) {
	ps.tokens = tokens
    ps.cells = []Cell{}
	ps.current = 0

	for !ps.is_at_end() {
        ps.cells = append(ps.cells, Cell{expr: ps.cell()})
	}
}

func (ps *Parser) cell() Expr {
	expr := ps.expression()
	if ps.match(NEW_LINE, COMMA, EOF) {
		return expr
	}
	panic(errors.New("expected EOF, comma, or new line at end of cell"))
}

func (ps *Parser) expression() Expr {
	if ps.match(EQUAL) {
		return ps.term()
	}
	literal := ps.literal()
	return literal
}

func (ps *Parser) term() Expr {
	expr := ps.factor()

	for ps.match(MINUS, PLUS) {
		operator := ps.previous()
		right := ps.factor()
		expr = &Binary{left: expr, operator: operator, right: right}
	}

	return expr
}

func (ps *Parser) factor() Expr {
	expr := ps.unary()

	for ps.match(STAR, SLASH) {
		operator := ps.previous()
		right := ps.unary()
		expr = &Binary{left: expr, operator: operator, right: right}
	}

	return expr
}

func (ps *Parser) unary() Expr {
	if ps.match(MINUS) {
		operator := ps.previous()
		right := ps.primary()
		return &Unary{operator: operator, right: right}
	}
	return ps.primary()
}

func (ps *Parser) primary() Expr {
	if ps.match(LEFT_PAREN) {
		expr := ps.term()
		ps.consume(RIGHT_PAREN, "expected ')'")
		return &Grouping{expr: expr}
	}

	if ps.match(IDENTIFIER) {
		return &Variable{name: ps.previous()}
	}

	return ps.literal()
}

func (ps *Parser) literal() Expr {
	if ps.match(NUMBER, STRING) {
		return &Literal{value: ps.previous().literal}
	}
	panic(errors.New("invalid literal"))
}

// HELPER FUNCTIONS

func (ps *Parser) check(kind TokenType) bool {
	if ps.is_at_end() {
		return false
	}
	return ps.peek().kind == kind
}

func (ps *Parser) match(kinds ...TokenType) bool {
	for _, kind := range kinds {
		if ps.check(kind) {
			ps.advance()
			return true
		}
	}
	return false
}

func (ps *Parser) peek() Token {
	return ps.tokens[ps.current]
}

func (ps *Parser) previous() Token {
	return ps.tokens[ps.current-1]
}

func (ps *Parser) is_at_end() bool {
	return ps.peek().kind == EOF
}

func (ps *Parser) advance() Token {
	if !ps.is_at_end() {
		ps.current++
		return ps.previous()
	}
	return ps.previous()
}

func (ps *Parser) consume(kind TokenType, message string) Token {
	if ps.check(kind) {
		return ps.advance()
	}
	panic(errors.New(message))
}
