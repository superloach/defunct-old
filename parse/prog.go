package parse

import (
	"strings"

	"github.com/alecthomas/participle"
)

var Parser = participle.MustBuild(
	&Prog{},
	participle.Lexer(Lexer),
	participle.Elide("Space"),
)

type Prog struct {
	Exprs []*Expr `parser:"@@*"`
}

func (p *Prog) String() string {
	ss := make([]string, 0, len(p.Exprs))

	for _, c := range p.Exprs {
		ss = append(ss, c.String())
	}

	return strings.Join(ss, "\n")
}
