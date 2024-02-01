package main

import (
    "fmt"
    "os"
)

type op int

const (
    ADD op = iota
    SUB op = iota
    MUL op = iota
    DIV op = iota
)

type eq struct {
    operation op
    operand1 *cell
    operand2 *cell
}

type cell struct {
    col int
    row int
    lexeme string
    value float32
    equation eq
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func parse(file []byte) ([][]cell) {
    
    grid := [][]cell{{}}
    cur_row := 0
    cur_col := 0

    curr_word := ""

    for i := 0; i < len(file); i++ {
        switch file[i] {
        case ',':
            new_cell := cell{
                col: cur_col,
                row: cur_row,
                lexeme: curr_word,
            }
            grid[cur_row] = append(grid[cur_row], new_cell)
            fmt.Println("comma")
            curr_word = ""
            cur_col += 1
        case '\n':
            new_cell := cell{
                col: cur_col,
                row: cur_row,
                lexeme: curr_word,
            }
            grid[cur_row] = append(grid[cur_row], new_cell)
            fmt.Println("newline")
            curr_word = ""
            cur_row += 1
            grid = append(grid, []cell{})
        default:
            curr_word += string(file[i])
        }
    }

    return grid 
}

func main() {
    dat, err := os.ReadFile("./example.csv")
    check(err)

    grid := parse(dat)

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            fmt.Printf("%s,", grid[i][j].lexeme)
        }
        fmt.Println()
    }
}
