package parse

import (
	"errors"
	"fmt"
	"strings"
)

var ErrNoBlock = errors.New("no block")

type Block Section

func (p *Parse) Block() (*Block, error) {
	s, err := p.Section(RuneStart, RuneEnd)
	if err != nil {
		if errors.Is(err, ErrNoSection) {
			return nil, ErrNoBlock
		}

		return nil, fmt.Errorf("section %q %q: %w", RuneStart, RuneEnd, err)
	}

	b := (*Block)(s)

	return b, nil
}

func (b *Block) String() string {
	ss := make([]string, len(b.Exprs))

	for i, e := range b.Exprs {
		ss[i] = e.String()
	}

	return string(RuneStart) +
		strings.Join(ss, " ") +
		string(RuneEnd)
}
