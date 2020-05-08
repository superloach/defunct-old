package functs

type Zilch struct{}

func (z Zilch) Call(args []Funct) Funct {
	return z
}

func (z Zilch) GetProp(name string) Funct {
	return z
}

func (z Zilch) SetProp(name string, val Funct) Funct {
	return z
}

func (z Zilch) String() string {
	return "{Zilch}"
}
