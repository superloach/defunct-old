package run

import (
	"github.com/pkg/errors"

	"github.com/superloach/defunct/parse"
)

func Expr(e *parse.Expr, u *Under) (Funct, error) {
	li, err := Literal(e.Literal, u)
	if err != nil {
		return nil, errors.Wrap(err, "literal")
	}

	for _, as := range e.Arguments {
		in, err := Arguments(as, u)
		if err != nil {
			return nil, errors.Wrap(err, "arguments")
		}

		nli, err := li.Call(u, in)
		if err != nil {
			return nil, errors.Wrap(err, "call")
		}

		li = nli
	}

	return li, nil
}
