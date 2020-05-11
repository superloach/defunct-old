package liner

import (
	"github.com/superloach/defunct/functs"
	"github.com/superloach/defunct/parser"
	"github.com/superloach/defunct/wrap"
)

const (
	f1 = "%s %#v\n :: %s\n"
	f2 = " -> %s\n\n"
)

type Liner struct {
	*parser.BaseDefunctListener

	Debugf func(string, ...interface{})
	Wrap   *wrap.Wrap
}

func (l *Liner) EnterFunct(c *parser.FunctContext) {
	t := c.GetText()
	l.Debugf(f1, "enter funct", t, l.Wrap)

	if t == "_" {
		l.Wrap.Funct = &functs.Under{}
	} else {
		l.Wrap.Funct = functs.String(t)
	}

	l.Debugf(f2, l.Wrap)
}

func (l *Liner) EnterArgs(c *parser.ArgsContext) {
	t := c.GetText()
	l.Debugf(f1, "enter args", t, l.Wrap)

	args := make(wrap.Args, 0)
	l.Wrap.Argss = append(l.Wrap.Argss, &args)

	l.Debugf(f2, l.Wrap)
}

func (l *Liner) EnterWrap(c *parser.WrapContext) {
	t := c.GetText()
	l.Debugf(f1, "enter wrap", t, l.Wrap)

	nc := &wrap.Wrap{
		Funct:  nil,
		Argss:  nil,
		Parent: l.Wrap,
		Debugf: l.Debugf,
	}
	if args := l.Wrap.Args(); args != nil {
		*args = append(*args, nc)
	}
	l.Wrap = nc

	l.Debugf(f2, l.Wrap)
}

func (l *Liner) ExitFunct(c *parser.FunctContext) {
	t := c.GetText()
	l.Debugf(f1, "exit funct", t, l.Wrap)

	l.Debugf(f2, l.Wrap)
}

func (l *Liner) ExitArgs(c *parser.ArgsContext) {
	t := c.GetText()
	l.Debugf(f1, "exit args", t, l.Wrap)

	l.Debugf(f2, l.Wrap)
}

func (l *Liner) ExitWrap(c *parser.WrapContext) {
	t := c.GetText()
	l.Debugf(f1, "exit wrap", t, l.Wrap)

	if l.Wrap.Parent != nil {
		l.Wrap = l.Wrap.Parent
	}

	l.Debugf(f2, l.Wrap)
}
