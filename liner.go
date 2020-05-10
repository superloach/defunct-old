package defunct

import (
	"github.com/superloach/defunct/functs"
	"github.com/superloach/defunct/parser"
)

const (
	f1 = "%s %#v\n :: %s\n"
	f2 = " -> %s\n\n"
)

type Liner struct {
	*parser.BaseDefunctListener

	Arena *Arena
	Wrap  *Wrap
}

func (l *Liner) EnterFunct(c *parser.FunctContext) {
	t := c.GetText()
	l.Arena.Debugf(f1, "enter funct", t, l.Wrap)

	if t == "_" {
		l.Wrap.Funct = l.Arena.Under
	} else {
		l.Wrap.Funct = functs.String(t)
	}

	l.Arena.Debugf(f2, l.Wrap)
}

func (l *Liner) EnterArgs(c *parser.ArgsContext) {
	t := c.GetText()
	l.Arena.Debugf(f1, "enter args", t, l.Wrap)

	args := make(Args, 0)
	l.Wrap.Argss = append(l.Wrap.Argss, &args)

	l.Arena.Debugf(f2, l.Wrap)
}

func (l *Liner) EnterWrap(c *parser.WrapContext) {
	t := c.GetText()
	l.Arena.Debugf(f1, "enter wrap", t, l.Wrap)

	nc := &Wrap{
		Arena:  l.Arena,
		Funct:  nil,
		Argss:  nil,
		Parent: l.Wrap,
	}
	if args := l.Wrap.Args(); args != nil {
		*args = append(*args, nc)
	}
	l.Wrap = nc

	l.Arena.Debugf(f2, l.Wrap)
}

func (l *Liner) ExitFunct(c *parser.FunctContext) {
	t := c.GetText()
	l.Arena.Debugf(f1, "exit funct", t, l.Wrap)

	l.Arena.Debugf(f2, l.Wrap)
}

func (l *Liner) ExitArgs(c *parser.ArgsContext) {
	t := c.GetText()
	l.Arena.Debugf(f1, "exit args", t, l.Wrap)

	l.Arena.Debugf(f2, l.Wrap)
}

func (l *Liner) ExitWrap(c *parser.WrapContext) {
	t := c.GetText()
	l.Arena.Debugf(f1, "exit wrap", t, l.Wrap)

	if l.Wrap.Parent != nil {
		l.Wrap = l.Wrap.Parent
	}

	l.Arena.Debugf(f2, l.Wrap)
}
