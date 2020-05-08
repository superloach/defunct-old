# defunct
an esoteric language based around Functs

## basic spec
### Funct
everything is one of these. in Go, this is the following interface:
```
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
_[_ ./X] -> loads and returns the Funct in the file ./X.df
```

## other builtins
there are other builtin types used by some implementations:

### Error
basically a String, but indicates that an error has occurred (implements a Go `error`)

### Native
a basic function (a Go `func(args []Funct) Funct`)

### Thing
a basic key-value map (a Go `map[string]Funct`)

### Zilch
no value, does nothing but return itself
