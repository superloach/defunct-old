package functs

import "fmt"

type Under struct {
	Args Funct
}

func (u *Under) Call(args []Funct) Funct {
	switch len(args) {
	case 0:
		return Zilch{}
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
	println("getprop", name)

	switch name {
	case "args":
		return u.Args
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

			return Zilch{}
		})
	default:
		println("TODO: Under get prop")
		return Zilch{}
	}
}

func (u *Under) SetProp(name string, val Funct) Funct {
	switch name {
	case "args":
		u.Args = val
		return u
	default:
		println("TODO: Under set prop")
		return Zilch{}
	}
}

func (u *Under) String() string {
	return "{Under}"
}
