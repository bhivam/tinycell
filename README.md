## Description
A non-gui csv calculator inspired by youtuber/programming [tsoding](https://github.com/tsoding/minicel). The goal is to take in csv files like this:

```
1, 2, 3
1, 2, 3
=A1+A2, =B1+B2, =C1+C2
```
and out put this:

```
1, 2, 3
1, 2, 3
2, 4, 6
```

## Usage
Edit example.csv to have whatever content you would like.

Then do `go run tinycell.go`

## WIP

These are a list of things which I would like to add to the project. The priority of each of these items is indicated by '!' at the beginning. The more '!' the higher the priority

- !!! Create ability to do multiple operations in one cell
- !! Add precedence operators like '()'
- !! Add unary operators '-'
- !! Add scope for other value types like text
- ! Add in basic, builtin functions
- ! Allow for user defined functions
- ! Better error tracking
- ! Testing
- ! Rewrite parser in cleaner, more exenstible way (Only after testing for each feature is added to ensure same behaviour)
