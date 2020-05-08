package functs

type Native func(args []Funct) Funct

func (n Native) Call(args []Funct) Funct {
	return n(args)
}

func (n Native) GetProp(name string) Funct {
	return Error("can't get prop on " + n.String())
}

func (n Native) SetProp(name string, val Funct) Funct {
	return Error("can't set prop on " + n.String())
}

func (n Native) String() string {
	return "{Native}"
}
