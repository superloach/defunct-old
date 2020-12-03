package parse

import (
	"strings"

	"github.com/pkg/errors"
)

type Module struct {
	Name  string
	Loc   Loc
	Exprs []*Expr
}

func (p *Parse) Module(name string) (*Module, error) {
	m := &Module{
		Name:  name,
		Loc:   p.Loc,
		Exprs: nil,
	}

	for {
		e, err := p.Expr()
		if err != nil {
			if errors.Is(err, ErrNoExpr) {
				return m, nil
			}

			return nil, errors.Wrap(err, "expr")
		}

		m.Exprs = append(m.Exprs, e)
	}
}

func (m *Module) String() string {
	ss := make([]string, len(m.Exprs))

	for i, e := range m.Exprs {
		ss[i] = e.String()
	}

	return strings.Join(ss, "\n")
}
