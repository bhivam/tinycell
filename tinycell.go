package main

import (
	"fmt"
	"os"
)

var has_error bool = false

func report_error(message string, line int) {
    has_error = true
    fmt.Println(message)
    fmt.Printf("Error occured on line %d.\n", line)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
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

func main() {
	dat, err := os.ReadFile("./example.csv")
    num_cols := get_num_cols(dat)
	check(err)

    sn := Scanner{}
    sn.scan_tokens(dat)
    if has_error {
        return
    }
   
    ps := Parser{}
    ps.parse(sn.tokens)

    printer := &ASTprinter{}
    for _, cell := range ps.cells {
        // TODO print out the cell each one is from
        fmt.Println(cell.expr.accept(printer))
    }

    it := &Interpreter{}
    it.interpret(ps.cells, num_cols) 

    for i := 0; i < len(it.cells)/num_cols; i++ {
		for j := 0; j < num_cols; j++ {
			cell := it.cells[i*num_cols+j]
            fmt.Print(literal_to_string(cell.value))
			fmt.Print(", ")
		}
		fmt.Println()
	}
}
