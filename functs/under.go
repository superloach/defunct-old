package functs

import (
	"fmt"
)

type Under struct {
	In     Funct
	Out    Funct
	Debugf func(string, ...interface{})
}

func (u *Under) Call(under Funct, args []Funct) Funct {
	u.Debugf("under call %s\n", args)

	switch len(args) {
	case 0:
		return Zilch
	case 1:
		return String(args[0].String())
	case 2:
		if key, ok := args[1].(String); ok {
			return args[0].GetProp(under, string(key))
		} else {
			return Error("can't use " + args[1].String() + " as prop")
		}
	case 3:
		if key, ok := args[1].(String); ok {
			return args[0].SetProp(under, string(key), args[2])
		} else {
			return Error("can't use " + args[1].String() + " as prop")
		}
	default:
		return Error("too many args to _")
	}
}

func (u *Under) GetProp(under Funct, name string) Funct {
	u.Debugf("under getprop %#v\n", name)

	switch name {
	case "in":
		return u.In
	case "print":
		return &Native{
			CallFn: func(under Funct, args []Funct) Funct {
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
			},
		}
	default:
		u.Debugf("TODO: Under get prop\n")
		return Zilch
	}
}

func (u *Under) SetProp(under Funct, name string, val Funct) Funct {
	switch name {
	case "args":
		u.In = val
		return u
	default:
		u.Debugf("TODO: Under set prop\n")
		return Zilch
	}
}

func (u *Under) String() string {
	return "{Under " + u.In.String() + " " + u.Out.String() + "}"
}
