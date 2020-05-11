package functs

type Funct interface {
	Call(under Funct, args []Funct) Funct
	GetProp(under Funct, name string) Funct
	SetProp(under Funct, name string, val Funct) Funct
	String() string
}
