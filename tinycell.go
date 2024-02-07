package main

import (
	"flag"
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
    show_exprs := flag.Bool("e", false,
        `Will print the expressions parsed in each cell of the input`)

    
    flag.Parse()
    vargs := flag.Args() 

    if (len(vargs) < 2) {
        fmt.Println("usage: tinycell [-e] <input csv> <output csv>")
        return 
    }

	dat, err := os.ReadFile(vargs[0])
    num_cols := get_num_cols(dat)
	check(err)

    sn := Scanner{}
    sn.scan_tokens(dat)
    if has_error {
        return
    }
   
    ps := Parser{}
    ps.parse(sn.tokens)
    
    if *show_exprs {
        printer := &ASTprinter{}
        for _, cell := range ps.cells {
            // TODO print out the cell each one is from
            fmt.Println(cell.expr.accept(printer))
        }
    }
    

    it := &Interpreter{}
    it.interpret(ps.cells, num_cols) 

    f, err := os.Create(vargs[1])
    check(err)

    defer f.Close()

    for i := 0; i < len(it.cells)/num_cols; i++ {
		for j := 0; j < num_cols; j++ {
			cell := it.cells[i*num_cols+j]
            f.Write([]byte(literal_to_string(cell.value)))
			f.Write([]byte(", "))
		}
		f.Write([]byte{'\n'})
	}
}
