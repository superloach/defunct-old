package run

import (
	"fmt"
	"strconv"
)

type Numberer interface {
	Number() (Number, error)
}

type Number int

func NumberOf(f Funct) (Number, error) {
	if n, ok := f.(Numberer); ok {
		return n.Number()
	}

	s := f.String()

	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("atoi %q: %w", s, err)
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
