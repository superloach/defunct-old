package functs

type String string

func (s String) Call(args []Funct) Funct {
	println("TODO: String formatter")
	return Zilch{}
}

func (s String) GetProp(name string) Funct {
	println("TODO: methods on String")
	return Zilch{}
}

func (s String) SetProp(name string, val Funct) Funct {
	return Error("can't set prop on " + s.String())
}

func (s String) String() string {
	return string(s)
}
