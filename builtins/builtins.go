package builtins

import "github.com/superloach/defunct/types"

var Builtins = types.Thing{
	"concat": Concat,
	"print":  Print,
}
