package run

import (
	"strconv"

	"github.com/pkg/errors"
)

type Numberer interface {
	Number() (Number, error)
}

type Number int

func NumberOf(f Funct) (Number, error) {
	if n, ok := f.(Numberer); ok {
		return n.Number()
	}

	i, err := strconv.Atoi(f.String())
	if err != nil {
		return 0, errors.Wrap(err, "atoi")
	}

	return Number(i), nil
}

func (n Number) String() string {
	return strconv.Itoa(int(n))
}

func (n Number) Call(_ *Under, fs ...Funct) (Funct, error) {
	panic("number call stub")
}

func (n Number) Number() (Number, error) {
	return n, nil
}
