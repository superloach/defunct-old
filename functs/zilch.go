package functs

type ZilchType struct{}

func (z ZilchType) Call(args []Funct) Funct {
	return z
}

func (z ZilchType) GetProp(name string) Funct {
	return z
}

func (z ZilchType) SetProp(name string, val Funct) Funct {
	return z
}

func (z ZilchType) String() string {
	return "{Zilch}"
}

var Zilch = ZilchType{}
