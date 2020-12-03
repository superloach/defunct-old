package run

import (
	"errors"
	"strings"
)

var (
	ErrQueueArgs    = errors.New("queue takes 0, 1, or 2 args")
	ErrQueueUnknown = errors.New("unknown queue command")
)

type Queue []Funct

func StringsQueue(ss ...string) Queue {
	q := make(Queue, len(ss))

	for i, s := range ss {
		q[i] = Text(s)
	}

	return q
}

func (q Queue) String() string {
	ss := make([]string, len(q))

	for i, f := range q {
		ss[i] = f.String()
	}

	return "[" + strings.Join(ss, " ") + "]"
}

func (q Queue) Call(_ *Under, in Queue) (Funct, error) {
	switch len(in) {
	case 0:
		panic("copy stub")

	case 1:
		switch in[0].String() {
		case "map":
			panic("map stub")
		}

		return nil, ErrQueueUnknown

	case 2:
		switch in[0].String() {
		case "get":
			panic("get stub")

		case "set":
			panic("set stub")
		}

		return nil, ErrQueueUnknown
	}

	return nil, ErrQueueArgs
}
