package parse

import (
	"errors"
	"fmt"
	"strings"
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

			return nil, fmt.Errorf("expr: %w", err)
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
