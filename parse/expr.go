package parse

import (
	"errors"
	"fmt"
)

var ErrNoExpr = errors.New("no expr")

type Expr struct {
	Literal   *Literal
	Arguments []*Arguments
}

func (p *Parse) Expr() (*Expr, error) {
	li, err := p.Literal()
	if err != nil {
		if errors.Is(err, ErrNoLiteral) {
			return nil, ErrNoExpr
		}

		return nil, fmt.Errorf("literal: %w", err)
	}

	e := &Expr{
		Literal: li,
	}

	for {
		as, err := p.Arguments()
		if err != nil {
			if errors.Is(err, ErrNoArguments) {
				return e, nil
			}

			return nil, fmt.Errorf("arguments: %w", err)
		}

		e.Arguments = append(e.Arguments, as)
	}
}

func (e *Expr) String() string {
	s := e.Literal.String()

	for _, as := range e.Arguments {
		s += as.String()
	}

	return s
}
