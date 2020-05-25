package builtins

import "github.com/superloach/defunct/types"

var Builtins = types.Thing{
	"apply":  Apply,
	"concat": Concat,
	"print":  Print,
}

func BuiltinString(name string) func() string {
	return func() string {
		return "{Builtin " + name + "}"
	}
}
