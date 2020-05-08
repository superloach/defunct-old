package functs

import "fmt"

type Error string

func (e Error) Call(args []Funct) Funct {
	return Error("can't call " + e.String())
}

func (e Error) GetProp(name string) Funct {
	return Error("can't get prop on " + e.String())
}

func (e Error) SetProp(name string, val Funct) Funct {
	return Error("can't set prop on " + e.String())
}

func (e Error) String() string {
	return fmt.Sprintf("{Error %#v}", string(e))
}

func (e Error) Error() string {
	return string(e)
}
