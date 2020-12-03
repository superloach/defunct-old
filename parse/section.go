package parse

import (
	"io"

	"github.com/pkg/errors"
)

var (
	ErrNoSection    = errors.New("no section")
	ErrCloseSection = errors.New("unclosed section")
)

type Section Module

func (p *Parse) Section(start, end rune) (*Section, error) {
	loc := p.Loc

	r, err := p.skipPeek()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil, ErrNoSection
		}

		return nil, errors.Wrap(err, "peek")
	}

	if r != start {
		return nil, ErrNoSection
	}

	p.incr()

	m, err := p.Module("")
	if err != nil {
		return nil, errors.Wrap(err, "module")
	}

	m.Loc = loc

	r, err = p.skipPeek()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil, ErrCloseSection
		}

		return nil, errors.Wrap(err, "peek")
	}

	if r != end {
		return nil, ErrCloseSection
	}

	p.incr()

	s := (*Section)(m)

	return s, nil
}
