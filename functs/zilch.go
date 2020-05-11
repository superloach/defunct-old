package functs

type ZilchType struct{}

func (z ZilchType) Call(under Funct, args []Funct) Funct {
	return z
}

func (z ZilchType) GetProp(under Funct, name string) Funct {
	return z
}

func (z ZilchType) SetProp(under Funct, name string, val Funct) Funct {
	return z
}

func (z ZilchType) String() string {
	return "{Zilch}"
}

var Zilch = ZilchType{}
