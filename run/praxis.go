package run

import "strings"

type Praxiser interface {
	Praxis() []string
}

type Praxis struct {
	Funct Funct
	Lines []string
}

func PraxisOf(f Funct) *Praxis {
	ls := []string{"unknown"}

	if p, ok := f.(Praxiser); ok {
		ls = p.Praxis()
	}

	return &Praxis{
		Funct: f,
		Lines: ls,
	}
}

func (p *Praxis) Error() string {
	return "praxis of " + p.Funct.String() + ":\n\t" + strings.Join(p.Lines, "\n\t")
}
