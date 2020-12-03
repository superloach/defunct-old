package parse

import (
	"strings"

	"github.com/pkg/errors"
)

var ErrNoArguments = errors.New("no arguments")

type Arguments Section

func (p *Parse) Arguments() (*Arguments, error) {
	s, err := p.Section(RuneOpen, RuneClose)
	if err != nil {
		if errors.Is(err, ErrNoSection) {
			return nil, ErrNoArguments
		}

		return nil, errors.Wrap(err, "section")
	}

	as := (*Arguments)(s)

	return as, nil
}

func (as *Arguments) String() string {
	ss := make([]string, len(as.Exprs))

	for i, e := range as.Exprs {
		ss[i] = e.String()
	}

	return string(RuneOpen) +
		strings.Join(ss, " ") +
		string(RuneClose)
}
