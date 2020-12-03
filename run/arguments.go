package run

import (
	"fmt"

	"github.com/superloach/defunct/parse"
)

func Arguments(as *parse.Arguments, u *Under) (Queue, error) {
	q := make(Queue, len(as.Exprs))

	for i, e := range as.Exprs {
		f, err := Expr(e, u)
		if err != nil {
			return nil, fmt.Errorf("expr %q: %w", e, err)
		}

		q[i] = f
	}

	return q, nil
}
