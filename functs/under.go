package functs

import (
	"fmt"

	"github.com/superloach/defunct/debug"
)

type Under struct {
	In Funct
}

func (u *Under) Call(args []Funct) Funct {
	debug.Debugf("under call %s\n", args)

	switch len(args) {
	case 0:
		return Zilch
	case 1:
		return String(args[0].String())
	case 2:
		return args[0].GetProp(args[1].String())
	case 3:
		return args[0].SetProp(args[1].String(), args[2])
	default:
		return Error("too many args to _")
	}
}

func (u *Under) GetProp(name string) Funct {
	debug.Debugf("under getprop %#v\n", name)

	switch name {
	case "in":
		return u.In
	case "print":
		return Native(func(args []Funct) Funct {
			iargs := make([]interface{}, len(args))
			for i, arg := range args {
				if s, ok := arg.(String); ok {
					iargs[i] = string(s)
				} else {
					iargs[i] = arg
				}
			}

			fmt.Println(iargs...)

			return Zilch
		})
	default:
		debug.Debugf("TODO: Under get prop")
		return Zilch
	}
}

func (u *Under) SetProp(name string, val Funct) Funct {
	switch name {
	case "args":
		u.In = val
		return u
	default:
		debug.Debugf("TODO: Under set prop")
		return Zilch
	}
}

func (u *Under) String() string {
	return "{Under}"
}
