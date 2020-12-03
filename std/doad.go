package std

import (
	"errors"

	"github.com/superloach/defunct/run"
)

var ErrDoad = errors.New("not a typewriter")

type Doad struct{}

func (d Doad) String() string {
	return "doad"
}

func (d Doad) Call(_ *run.Under, _ run.Queue) (run.Funct, error) {
	return nil, ErrDoad
}

func (d Doad) Praxis() []string {
	return []string{`[] -> halt and catch fire`}
}
