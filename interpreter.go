package main

import (
	"errors"
	"math"
	"strconv"
)

type Interpreter struct{
    cells []Cell
    num_cols int
}

func (it *Interpreter) interpret(cells []Cell, num_cols int) {
    it.cells = cells
    it.num_cols = num_cols
    
    for i := range it.cells {
        cell := &cells[i]

        if cell.calculating {
            panic(errors.New("cycle detected"))
        }
        cell.calculating = true
        cell.value = it.evaluate(cell.expr)
        cell.calculating = false  
    }
}

func (it *Interpreter) visit_binary(expr *Binary) any {
    op := expr.operator.kind
    left_val, left_ok := it.evaluate(expr.left).(float64)
    right_val, right_ok := it.evaluate(expr.right).(float64)
    
    if !(left_ok && right_ok) {
        panic(errors.New("passsing non-numerics to binary operation"))
    }

    switch op {
    case PLUS:
         return left_val + right_val
    case MINUS:
         return left_val - right_val
    case SLASH:
         return left_val / right_val
    case STAR:
         return left_val * right_val
    }
    panic(errors.New("Unkown operation type"))
}

func (it *Interpreter) visit_unary(expr *Unary) any {
    right_val, right_ok := it.evaluate(expr).(float64)

    if !right_ok {
        panic(errors.New("passsing non-numerics to binary operation"))
    }  

    if expr.operator.kind != MINUS {
        panic(errors.New("Unkown operation type"))
    }

    return -right_val
}

func (it *Interpreter) visit_grouping(expr *Grouping) any {
	return it.evaluate(expr.expr)
}

func (it *Interpreter) visit_literal(expr *Literal) any {
	return expr.value
}

func (it *Interpreter) visit_variable(expr *Variable) any {
    cell := it.cells[it.get_index(expr.name.lexeme)]
    println(expr.name.lexeme, cell.calculating)
    if cell.value == nil {
        if cell.calculating {
            panic(errors.New("cycle detected"))
        }
        cell.calculating = true
        cell.value = it.evaluate(cell.expr)
        cell.calculating = false
    }
	return cell.value
}

func (it *Interpreter) evaluate(expr Expr) any {
	return expr.accept(it)
}

// HELPER FUNCTIONS

func (it *Interpreter) get_index(str_index string) int {
	var num_start int
	str := ""

	for i := range str_index {
		c := str_index[i]
		if is_num(c) {
			num_start = i
			break
		}
		str += string(c)
	}

	if num_start == 0 {
		panic(errors.New("table index does not start with number"))
	}

	col_i := 0
	for i := range str {
		c := str[i]
		col_i += int(float64(c-64) * math.Pow(26.0, float64(len(str)-(i+1))))
	}

	str = ""
	for i := num_start; i < len(str_index); i++ {
		str += string(str_index[i])
	}

	row_i, err := strconv.ParseInt(str, 10, 64)
	check(err)

	return col_i - 1 + int(row_i-1)*it.num_cols
}

