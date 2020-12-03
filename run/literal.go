package run

import (
	"github.com/pkg/errors"

	"github.com/superloach/defunct/parse"
)

var ErrUnknownLiteral = errors.New("unknown literal")

func Literal(l *parse.Literal, u *Under) (Funct, error) {
	switch {
	case l.Under:
		return u, nil
	case len(l.Text) > 0:
		return Text(l.Text), nil
	case l.Block != nil:
		return (*Block)(l.Block), nil
	}

	return nil, ErrUnknownLiteral
}
