package builtins

import (
	"strconv"

	"github.com/superloach/defunct/types"
)

var Apply = &types.Native{
	CallFn: func(under types.Funct, args []types.Funct) types.Funct {
		if len(args) == 1 {
			args = append(args, types.Zilch)
		}
		if len(args) != 2 {
			return types.Error("usage: apply[fn args]")
		}

		fn := args[0]
		nargs := args[1]

		gnargs := make([]types.Funct, 0)
		for i := 0; ; i++ {
			k := strconv.Itoa(i)
			prop := nargs.GetProp(under, k)
			if prop == nil || prop == types.Zilch {
				break
			}
			if err, ok := prop.(types.Error); ok {
				return err
			}
			gnargs = append(gnargs, prop)
		}

		return fn.Call(under, gnargs)
	},
	StringFn: BuiltinString("Apply"),
}
