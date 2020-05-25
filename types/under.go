package types

type Under struct {
	Arbitrary Thing
	Builtins  Funct
	Modules   Funct

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

	if f := u.Arbitrary.GetProp(under, name); name == "in" || name == "out" || f != Zilch {
		return f
	}

	if f := u.Builtins.GetProp(under, name); f != Zilch {
		return f
	}

	if f := u.Modules.GetProp(under, name); f != Zilch {
		return f
	}

	return Zilch
}

func (u *Under) SetProp(under Funct, name string, val Funct) Funct {
	u.Arbitrary.SetProp(under, name, val)

	return u
}

func (u *Under) String() string {
	return "{Under}"
}
