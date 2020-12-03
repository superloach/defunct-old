package run

import (
	"errors"
	"io"
)

var ErrUnknownStd = errors.New("unknown std funct")

type Under struct {
	Parent *Under

	In Queue

	Var Map
	Std Map
	Imp Map

	R io.Reader
	W io.Writer
}

func NewUnder(
	in Queue,
	std Map,
	r io.Reader,
	w io.Writer,
) *Under {
	return &Under{
		Parent: nil,

		In: in,

		Var: Map{},
		Std: std,
		Imp: Map{},

		R: r,
		W: w,
	}
}

func (u *Under) String() string {
	return "_"
}

func (u *Under) Call(_ *Under, in Queue) (Funct, error) {
	const (
		runeStd = '@'
		runeImp = '.'
	)

	switch len(in) {
	case 0:
		return Zilch, nil

	case 1:
		name := TextOf(in[0])

		switch name[0] {
		case runeStd:
			f, ok := u.Std[string(name[1:])]
			if !ok {
				return nil, ErrUnknownStd
			}

			return u.Got(f), nil

		case runeImp:
			f, ok := u.Std[string(name[1:])]
			if !ok {
				return u.Import(string(name[1:]))
			}

			return u.Got(f), nil
		}

		f, ok := u.Var[string(name)]
		if !ok {
			return Zilch, nil
		}

		return u.Got(f), nil

	case 2:
		name := TextOf(in[0])
		u.Var[string(name)] = in[1]

		return u.Got(in[1]), nil
	}

	return nil, PraxisOf(u)
}

func (u *Under) Child(in Queue) *Under {
	c := *u
	c.Parent = u
	c.Var = Map{}
	return &c
}

func (u *Under) Praxis() []string {
	return []string{
		`[] -> Zilch`,
		`[x] -> get x`,
		`[x y] -> set x to y`,
		``,
		`@x -> stdlib x`,
		`.x -> import x`,
	}
}
