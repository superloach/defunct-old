# defunct
an esoteric language based around Functs

## basic spec
everything is a Funct

there are 2 kinds of functs you can actually use in programs, Strings and Under

Strings are just any text that isn't `_`, `[`, or `]` (you can escape with `~`)

Under is a sort of predefined, global, magic Funct. it behaves as such:
```
_[] -> Zilch
_[X] -> String representation of X
_[X Y] -> X's property named by Y
_[X Y Z] -> X, after setting its property named by Y to the value Z
```

Under also does special things when used on itself:
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
