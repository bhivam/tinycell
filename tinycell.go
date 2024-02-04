package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

/*
 * Need to start enforcing alphabet instead of just not numberic
 */

type OpType int

const (
	ADD OpType = iota
	SUB OpType = iota
	MUL OpType = iota
	DIV OpType = iota
)

type ExprType int

const (
	NULL    ExprType = iota
	BINARY  ExprType = iota
	LITERAL ExprType = iota
)

// everything will be an expression, they will link to eachother
type Expr struct {
	expression_type ExprType
	// literal expression
	value float64
	// binary expression
	lhs            int // table index
	rhs            int
	lf_literal     float64
	rf_literal     float64
	operation_type OpType
}

type Token struct {
	lexeme string
	col    int
	row    int
}

type Cell struct {
	col        int
	row        int
	expression Expr
}

func get_op_type(op_sign byte) OpType {
	switch op_sign {
	case '+':
		return ADD
	case '-':
		return SUB
	case '*':
		return MUL
	case '/':
		return DIV
	default:
		panic(errors.New("no such op_sign"))
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func is_num(c byte) bool {
	return c > 47 && c < 58
}

func get_table_index(str_index string, num_cols int) int {
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

	return col_i - 1 + int(row_i-1)*num_cols
}

func parse_lexeme(lexeme string, num_cols int) Expr {
	new_expression := Expr{lhs: -1, rhs: -1}

	if len(lexeme) == 0 {
		new_expression.expression_type = NULL
		return new_expression
	}

	if lexeme[0] == '=' {
		new_expression.expression_type = BINARY
	} else {
		// should be albe to convert to float directly, if not we have problems in format
		new_expression.expression_type = LITERAL
		value, err := strconv.ParseFloat(lexeme, 64)
		check(err)
		new_expression.value = value

		return new_expression
	}

	operand := ""
	for i := 1; i < len(lexeme); i++ {
		c := lexeme[i]
		switch c {
		case '+':
			fallthrough
		case '-':
			fallthrough
		case '/':
			fallthrough
		case '*':
			new_expression.operation_type = get_op_type(c)
			if !is_num(operand[0]) { // first thing is not num, so table index
				new_expression.lhs = get_table_index(operand, num_cols)
			} else { // this is a literal
				var err error
				new_expression.lf_literal, err = strconv.ParseFloat(operand, 64)
				check(err)
			}
			operand = ""
			// detect if we have table index or literal
		default:
			operand += string(c)
		}
	}

	if !is_num(operand[0]) { // first thing is not num, so table index
		new_expression.rhs = get_table_index(operand, num_cols)
	} else { // this is a literal
		var err error
		new_expression.rf_literal, err = strconv.ParseFloat(operand, 64)
		check(err)
	}

	return new_expression
}

func get_num_cols(file []byte) int {
	i := 0
	num_commas := 0

	for ; file[i] != '\n'; i++ {
		if file[i] == ',' {
			num_commas += 1
		}
	}

	return num_commas + 1
}

/*
ignoring spaces completely rn, should only be ignoring comma, newline adj whitespace
*/
func parse_file(file []byte) []Expr {
	tokens := []Token{}
	cur_row := 0
	cur_col := 0

	current_word := ""

	for i := 0; i < len(file); i++ {
		switch file[i] {
		case ' ':
			continue
		case ',':
			token := Token{
				lexeme: current_word,
				col:    cur_col,
				row:    cur_row,
			}
			tokens = append(tokens, token)
			current_word = ""
			cur_col += 1
		case '\n':
			token := Token{
				lexeme: current_word,
				col:    cur_col,
				row:    cur_row,
			}
			tokens = append(tokens, token)
			current_word = ""
			cur_row += 1
			cur_col = 0
		default:
			current_word += string(file[i])
		}
	}

	num_cols := get_num_cols(file)

	expressions := []Expr{}
	for i := range tokens {
		expressions = append(expressions, parse_lexeme(tokens[i].lexeme, num_cols))
	}

	return expressions
}

func main() {
	dat, err := os.ReadFile("./example.csv")
	check(err)

	expressions := parse_file(dat)

	for _, expression := range expressions {
		fmt.Printf("%d %d\n", expression.lhs, expression.rhs)
	}
}
