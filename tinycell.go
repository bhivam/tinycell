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

func main() {
	dat, err := os.ReadFile("./example.csv")
	check(err)

    sn := Scanner{}
    sn.scan_tokens(dat)
    if has_error {
        return
    }

    for _, token := range sn.tokens {
        fmt.Println(token.to_string()) 
    }
}
