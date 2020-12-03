package std

import (
	"github.com/superloach/defunct/run"
)

type Parent struct{}

func (p Parent) String() string {
	return "parent"
}

func (p Parent) Call(u *run.Under, in run.Queue) (run.Funct, error) {
	return nil, nil
}

func (p Parent) Got(u *run.Under) run.Funct {
	if u.Parent == nil {
		return run.Zilch
	}

	return u.Parent
}
