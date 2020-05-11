package types

type String string

func (s String) Call(under Funct, args []Funct) Funct {
	println("TODO: String formatter")
	return Zilch
}

func (s String) GetProp(under Funct, name string) Funct {
	println("TODO: methods on String")
	return Zilch
}

func (s String) SetProp(under Funct, name string, val Funct) Funct {
	return Error("can't set prop on " + s.String())
}

func (s String) String() string {
	return string(s)
}
