package parse

import (
	"strings"

	"github.com/pkg/errors"
)

func (p *Parse) peek() (rune, error) {
	r, _, err := p.Buf.ReadRune()
	if err != nil {
		return -1, errors.Wrap(err, "read rune")
	}

	err = p.Buf.UnreadRune()
	if err != nil {
		return -1, errors.Wrap(err, "unread rune")
	}

	return r, nil
}

func (p *Parse) skipPeek() (rune, error) {
	for {
		r, err := p.peek()
		if err != nil {
			return -1, errors.Wrap(err, "peek")
		}

		if !strings.ContainsRune(StringSpace, r) {
			return r, nil
		}

		p.incr()
	}
}

func (p *Parse) incr() {
	r, _, _ := p.Buf.ReadRune()

	p.Loc.Char++
	p.Loc.Col++

	if r == '\n' {
		p.Loc.Col = 0
		p.Loc.Line++
	}
}
