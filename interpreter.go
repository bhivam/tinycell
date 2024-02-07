package main

import (
	"errors"
	"math"
	"strconv"
)

type Interpreter struct {
	cells    []Cell
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
	left_val_f, left_ok_f := it.evaluate(expr.left).(float64)
	right_val_f, right_ok_f := it.evaluate(expr.right).(float64)

	left_val_s, left_ok_s := it.evaluate(expr.left).([]byte)
	right_val_s, right_ok_s := it.evaluate(expr.right).([]byte)
    

    if left_ok_f && right_ok_f {
        switch op {
        case PLUS:
            return left_val_f + right_val_f
        case MINUS:
            return left_val_f - right_val_f
        case SLASH:
            return left_val_f / right_val_f
        case STAR:
            return left_val_f * right_val_f
        }
	    panic(errors.New("Unkown operation type for number"))
    } else if left_ok_s && right_ok_s {
        if op == PLUS {
            return append(left_val_s, right_val_s...)
        }
	    panic(errors.New("Unkown operation type for string"))
    } else {
        panic(errors.New("Invalid operands for binary expression"))
    }
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
