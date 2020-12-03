package parse

import (
	"errors"
	"fmt"
	"io"
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

		return nil, fmt.Errorf("skip peek: %w", err)
	}

	if r != start {
		return nil, ErrNoSection
	}

	p.incr()

	const name = ""

	m, err := p.Module(name)
	if err != nil {
		return nil, fmt.Errorf("module %q: %w", name, err)
	}

	m.Loc = loc

	r, err = p.skipPeek()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil, ErrCloseSection
		}

		return nil, fmt.Errorf("skip peek: %w", err)
	}

	if r != end {
		return nil, ErrCloseSection
	}

	p.incr()

	s := (*Section)(m)

	return s, nil
}
