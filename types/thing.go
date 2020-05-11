package types

import "fmt"

type Thing map[string]Funct

func (t Thing) Call(under Funct, args []Funct) Funct {
	switch len(args) {
	case 0:
		return t
	case 1:
		return t.GetProp(under, args[0].String())
	case 2:
		return t.SetProp(under, args[0].String(), args[1])
	default:
		return Error("too many args to Thing")
	}
}

func (t Thing) GetProp(under Funct, name string) Funct {
	if p, ok := t[name]; ok {
		return p
	}
	return Zilch
}

func (t Thing) SetProp(under Funct, name string, val Funct) Funct {
	t[name] = val
	return t
}

func (t Thing) String() string {
	m := fmt.Sprintf("%s", (map[string]Funct)(t))
	return "{Thing " + m[4:len(m)-1] + "}"
}
