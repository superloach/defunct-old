package defunct

import (
	"github.com/superloach/defunct/debug"
	"github.com/superloach/defunct/functs"
	"github.com/superloach/defunct/parser"
)

const (
	f1 = "%s %#v\n :: %s\n"
	f2 = " -> %s\n\n"
)

type Listener struct {
	*parser.BaseDefunctListener

	Under *functs.Under
	Wrap  *Wrap
}

func (l *Listener) EnterFunct(c *parser.FunctContext) {
	t := c.GetText()
	debug.Debugf(f1, "enter funct", t, l.Wrap)

	if t == "_" {
		l.Wrap.Funct = l.Under
	} else {
		l.Wrap.Funct = functs.String(t)
	}

	debug.Debugf(f2, l.Wrap)
}

func (l *Listener) EnterArgs(c *parser.ArgsContext) {
	t := c.GetText()
	debug.Debugf(f1, "enter args", t, l.Wrap)

	args := make(Args, 0)
	l.Wrap.Argss = append(l.Wrap.Argss, &args)

	debug.Debugf(f2, l.Wrap)
}

func (l *Listener) EnterWrap(c *parser.WrapContext) {
	t := c.GetText()
	debug.Debugf(f1, "enter wrap", t, l.Wrap)

	nc := &Wrap{
		Funct:  nil,
		Argss:  nil,
		Parent: l.Wrap,
	}
	if args := l.Wrap.Args(); args != nil {
		*args = append(*args, nc)
	}
	l.Wrap = nc

	debug.Debugf(f2, l.Wrap)
}

func (l *Listener) ExitFunct(c *parser.FunctContext) {
	t := c.GetText()
	debug.Debugf(f1, "exit funct", t, l.Wrap)

	debug.Debugf(f2, l.Wrap)
}

func (l *Listener) ExitArgs(c *parser.ArgsContext) {
	t := c.GetText()
	debug.Debugf(f1, "exit args", t, l.Wrap)

	debug.Debugf(f2, l.Wrap)
}

func (l *Listener) ExitWrap(c *parser.WrapContext) {
	t := c.GetText()
	debug.Debugf(f1, "exit wrap", t, l.Wrap)

	if l.Wrap.Parent != nil {
		l.Wrap = l.Wrap.Parent
	}

	debug.Debugf(f2, l.Wrap)
}
