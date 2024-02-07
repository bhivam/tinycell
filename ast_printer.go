package main

type ASTprinter struct{}

func (printer *ASTprinter) visit_binary(expr *Binary) any {
	return printer.parenthesize(expr.operator.lexeme, expr.left, expr.right)
}

func (printer *ASTprinter) visit_unary(expr *Unary) any {
	return printer.parenthesize(expr.operator.lexeme, expr.right)
}

func (printer *ASTprinter) visit_grouping(expr *Grouping) any {
	return printer.parenthesize("group", expr.expr)
}

func (printer *ASTprinter) visit_literal(expr *Literal) any {
	if expr.value == nil {
		return "nil"
	}
	return literal_to_string(expr.value)
}

func (printer *ASTprinter) visit_variable(expr *Variable) any {
    return "(var " + expr.name.lexeme + ")" 
}

func (printer *ASTprinter) parenthesize(name string, exprs ...Expr) string {
	ret := "(" + name
	for _, expr := range exprs {
		expr_string := expr.accept(printer).(string)
		ret += " " + expr_string
	}
	ret += ")"

	return ret
}
