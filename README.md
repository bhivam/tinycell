## Description
A non-gui csv calculator inspired by [tsoding](https://github.com/tsoding/minicel). The goal is to take in csv files like this:

```
1, 2, 3
1, 2, 3
=A1+A2, =B1+B2, =C1+C2
```
and output this:

```
1, 2, 3
1, 2, 3
2, 4, 6
```

## Usage
Edit example.csv to have whatever content you would like.

Then do `go run tinycell.go -e <input.csv> <output.csv>`

The -e command line flag will enbale printing of the expression stored in each cell of the input csv.

## Features
### Equations Support
- arithmetic e.g. A1+B3
- negation e.g. -A1
- grouping e.g. (10+A1)*B2
- string concatenation e.g. "hello" + "world"

### Non-Equation Cell Data Type
- numbers
- strings
### Others
- Cycle Detection e.g. if we place "A1" in the cell A1, then the value cannot be computed

## WIP
These are a list of things which I would like to add to the project. The priority of each of these items is indicated by '!' at the beginning. The more '!' the higher the priority

- ! Add in basic, builtin functions
- ! Allow for user defined functions
- ! Better error tracking
- ! Testing
