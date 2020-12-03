package parse

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

var ErrNoLiteral = errors.New("no literal")

type Literal struct {
	Under bool
	Text  Text
	Block *Block
}

func (p *Parse) Literal() (*Literal, error) {
	r, err := p.skipPeek()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil, ErrNoLiteral
		}

		return nil, fmt.Errorf("skip peek: %w", err)
	}

	switch {
	case r == RuneStart:
		b, err := p.Block()
		if err != nil {
			if errors.Is(err, ErrNoBlock) {
				return nil, ErrNoLiteral
			}

			return nil, fmt.Errorf("block: %w", err)
		}

		return &Literal{
			Block: b,
		}, nil

	case r == RuneUnder:
		p.incr()

		return &Literal{
			Under: true,
		}, nil

	case !strings.ContainsRune(StringEnder, r):
		t, err := p.Text()
		if err != nil {
			if errors.Is(err, ErrNoText) {
				return nil, ErrNoLiteral
			}

			return nil, fmt.Errorf("text: %w", err)
		}

		return &Literal{
			Text: t,
		}, nil
	}

	return nil, ErrNoLiteral
}

func (l *Literal) String() string {
	switch {
	case l.Under:
		return string(RuneUnder)
	case len(l.Text) > 0:
		return l.Text.String()
	case l.Block != nil:
		return l.Block.String()
	}

	return "literal???"
}
