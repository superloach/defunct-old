package run

import (
	"github.com/pkg/errors"

	"github.com/superloach/defunct/parse"
)

type Module parse.Module

func (m *Module) String() string {
	return (*parse.Module)(m).String()
}

func (m *Module) Call(u *Under, in Queue) (Funct, error) {
	return u.Child(in).Run(m)
}

func (u *Under) Run(m *Module) (Funct, error) {
	ret, err := Zilch, error(nil)

	for _, e := range m.Exprs {
		ret, err = Expr(e, u)
		if err != nil {
			return nil, errors.Wrap(err, "expr")
		}
	}

	return ret, nil
}
