package parse

import (
	"strings"
)

type Call struct {
	Args []*Expr `parser:"ArgsOpen @@* ArgsClose"`
}

func (c *Call) String() string {
	ss := make([]string, 0, len(c.Args))

	for _, arg := range c.Args {
		ss = append(ss, arg.String())
	}

	return "[" + strings.Join(ss, " ") + "]"
}
