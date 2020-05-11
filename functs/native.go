package functs

type Native struct {
	CallFn    func(under Funct, args []Funct) Funct
	GetPropFn func(under Funct, name string) Funct
	SetPropFn func(under Funct, name string, val Funct) Funct
	StringFn  func() string
}

func (n *Native) Call(under Funct, args []Funct) Funct {
	if n.CallFn == nil {
		return Zilch
	}

	return n.CallFn(under, args)
}

func (n *Native) GetProp(under Funct, name string) Funct {
	if n.GetPropFn == nil {
		return Zilch
	}

	return n.GetPropFn(under, name)
}

func (n *Native) SetProp(under Funct, name string, val Funct) Funct {
	if n.SetPropFn == nil {
		return Zilch
	}

	return n.SetPropFn(under, name, val)
}

func (n *Native) String() string {
	if n.StringFn == nil {
		return "{Native}"
	}

	return n.StringFn()
}
