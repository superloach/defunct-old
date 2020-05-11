package builtins

import (
	"fmt"

	"github.com/superloach/defunct/types"
)

var Print = &types.Native{
	CallFn: func(under types.Funct, args []types.Funct) types.Funct {
		iargs := make([]interface{}, len(args))
		for i, arg := range args {
			iargs[i] = arg
		}

		fmt.Println(iargs...)

		return types.Zilch
	},
}
