package main

type ExprVisitor interface {
	visit_binary(*Binary) any
	visit_unary(*Unary) any
	visit_grouping(*Grouping) any
	visit_literal(*Literal) any
    visit_variable(*Variable) any
}

type Expr interface {
    accept(ExprVisitor) any
}

/* BINARY DEFINITION*/
type Binary struct {
	left     Expr
	right    Expr
	operator Token
}

func (expr *Binary) accept(visitor ExprVisitor) any {
    return visitor.visit_binary(expr)
}

/* GROUPING DEFINITION */
type Grouping struct {
	expr Expr
}

func (expr *Grouping) accept(visitor ExprVisitor) any {
    return visitor.visit_grouping(expr)
}

/* UNARY DEFINITION */
type Unary struct {
	right     Expr
	operator Token
}

func (expr *Unary) accept(visitor ExprVisitor) any {
    return visitor.visit_unary(expr)
}

/* LITERAL DEFINITION */
type Literal struct {
	value any
}

func (expr *Literal) accept(visitor ExprVisitor) any {
    return visitor.visit_literal(expr)
}

/* VARIABLE DEFINITION */
type Variable struct {
    name Token
}

func (expr *Variable) accept(visitor ExprVisitor) any {
    return visitor.visit_variable(expr)
}
