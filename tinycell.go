package main

import (
    "fmt"
    "os"
    "strconv"
)

type op int
type cell_type int

const (
    ADD op = iota
    SUB op = iota
    MUL op = iota
    DIV op = iota
)

const (
    equation cell_type = iota
    literal cell_type = iota
)

type eq struct {
    operation op
    value1 int
    value2 int
    col1 int
    row1 int
    col2 int
    row2 int
}

type cell struct {
    col int
    row int
    value float64
    equation eq
    ctype cell_type
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
    var err error
    
    new_cell := cell{row: cur_row, col: cur_col, ctype: literal}

    for i := 0; i < len(file); i++ {
        switch file[i] {
        case '=':
            if curr_word != "" {
                curr_word += string(file[i])
            } else {
                new_cell.ctype = equation
            }
        case '+': // add in other signs later
            if new_cell.ctype != equation {
                curr_word += string(file[i])
            } else {
                // curr_word, here, is an int literal or a cell
                // make a function to fgure it out
            }
        case ',':
            if new_cell.ctype == equation {
                // curr_word, here, is an int literal or a cell
                // make a function to fgure it out
            } else {
                fmt.Printf("%s", curr_word)
                new_cell.value, err = strconv.ParseFloat(curr_word, 64)
                check(err)
            }
            grid[cur_row] = append(grid[cur_row], new_cell)
            fmt.Println("comma")
            curr_word = ""
            cur_col += 1
            new_cell = cell{row: cur_row, col: cur_col}
        case '\n':
            // We want to detect when the type is different from this
            if new_cell.ctype == equation {
                // curr_word, here, is an int literal or a cell
                // make a function to fgure it out
            } else {
                fmt.Printf("%s", curr_word)
                new_cell.value, err = strconv.ParseFloat(curr_word, 64)
                check(err)
            }
            grid[cur_row] = append(grid[cur_row], new_cell)
            fmt.Println("newline")
            curr_word = ""
            cur_row += 1
            cur_col = 0;
            grid = append(grid, []cell{})
        case ' ':
            continue
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
            fmt.Printf("%f,", grid[i][j].value)
        }
        fmt.Println()
    }
}
