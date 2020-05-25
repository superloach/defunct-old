package builtins

import "github.com/superloach/defunct/types"

var Concat = &types.Native{
	CallFn: func(under types.Funct, args []types.Funct) types.Funct {
		s := ""

		for _, arg := range args {
			s += arg.String()
		}

		return types.String(s)
	},
	StringFn: BuiltinString("Concat"),
}
