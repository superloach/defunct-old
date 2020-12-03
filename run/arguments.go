package run

import (
	"github.com/pkg/errors"

	"github.com/superloach/defunct/parse"
)

func Arguments(as *parse.Arguments, u *Under) (Queue, error) {
	q := make(Queue, len(as.Exprs))

	for i, e := range as.Exprs {
		f, err := Expr(e, u)
		if err != nil {
			return nil, errors.Wrap(err, "argument")
		}

		q[i] = f
	}

	return q, nil
}
