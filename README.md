# defunct
an esoteric language based around Functs

## basic spec
### Funct
everything is one of these. in Go, this is the following interface:
```go
type Funct interface {
	Call(args []Funct) Funct
	GetProp(name string) Funct
	SetProp(name string, val Funct) Funct
	String() string
}
```

### String
a Funct containing a series of characters, preferably UTF-8.

declared with any text that isn't `_`, `[`, or `]` (you can escape these with `~`).

### Under
a sort of predefined, global, magic Funct. it behaves as such:
```
_[] -> Zilch
_[X] -> String representation of X
_[X Y] -> X's property named by Y
_[X Y Z] -> X, after setting its property named by Y to the value Z
```

also does special things when used on itself:
```
_[_ X] -> returns the stdlib Funct named by X
_[_ .X] -> loads and returns the Funct in the file ./X.df
```

## other builtins
there are other builtin types used by the Go implementation:

### Error
basically a String, but indicates that an error has occurred (implements a Go `error`).

### Native
a basic function (a Go `func(args []Funct) Funct`).

### Thing
a basic key-value map (a Go `map[string]Funct`).

### Zilch
no value, does nothing but return itself.

## multi-file system
each Defunct (`.df`) file contains a series of lines, each of which are made up of chained or nested Funct calls.

across all of these lines, the Under instance referred to by `_` is the same. ideally, all instances of Zilch will be the same, but that's not required.

when you use `_[_ .X]`, a Funct referencing the file `./X.df` is loaded. calling this Funct executes the whole file, with the arguments passed as `_[_ in]`. the return value is set with `_[_ out]`.

similarly, when a Defunct file is run with the interpreter, leftover command-line arguments are passed as `_[_ in]`, and setting `_[_ out]` allows a value to be printed after executing the file.

## roadmap
 - [x] define basic Funct types
 - [ ] implement parser
 - [ ] implement interpreter
 - [ ] possibly generate go code?

## build instructions
### generated code
requires [antlr4](https://github.com/antlr/antlr4) and [statik](github.com/rakyll/statik) to be installed
```bash
go generate .
```
### build the tool
```bash
go build ./cmd/defunct
```
