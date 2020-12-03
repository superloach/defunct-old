package run

import (
	"fmt"

	"github.com/superloach/defunct/parse"
)

func Expr(e *parse.Expr, u *Under) (Funct, error) {
	li, err := Literal(e.Literal, u)
	if err != nil {
		return nil, fmt.Errorf("literal %q: %w", e.Literal, err)
	}

	for _, as := range e.Arguments {
		in, err := Arguments(as, u)
		if err != nil {
			return nil, fmt.Errorf("arguments %q: %w", as, err)
		}

		nli, err := li.Call(u, in)
		if err != nil {
			return nil, fmt.Errorf("call %s%s: %w", li, in, err)
		}

		li = nli
	}

	return li, nil
}
