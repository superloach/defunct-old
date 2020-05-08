package functs

type Funct interface {
	Call(args []Funct) Funct
	GetProp(name string) Funct
	SetProp(name string, val Funct) Funct
	String() string
}
