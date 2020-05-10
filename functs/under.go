package functs

import (
	"fmt"
)

type Under struct {
	In    Funct
	Out   Funct
	Debug *bool
}

func (u *Under) Debugf(f string, args ...interface{}) {
	if *u.Debug {
		fmt.Printf(f, args...)
	}
}

func (u *Under) Call(args []Funct) Funct {
	u.Debugf("under call %s\n", args)

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
	u.Debugf("under getprop %#v\n", name)

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
		u.Debugf("TODO: Under get prop\n")
		return Zilch
	}
}

func (u *Under) SetProp(name string, val Funct) Funct {
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
	return "{Under}"
}
