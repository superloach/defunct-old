package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/superloach/defunct/parser/parser"
	"github.com/superloach/defunct/types"
	"github.com/superloach/defunct/wrap"
)

const (
	f1 = "%s %#v\n :: %s\n"
	f2 = " -> %s\n\n"
)

type Listener struct {
	Debugf func(string, ...interface{})
	Lines  []*wrap.Wrap
	Wrap   *wrap.Wrap
}

func (l *Listener) VisitTerminal(node antlr.TerminalNode) {
	l.Debugf("%#v\n", node)
}

func (l *Listener) VisitErrorNode(node antlr.ErrorNode) {
	l.Debugf("%#v\n", node)
}

func (l *Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	l.Debugf("%#v\n", ctx)
}

func (l *Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	l.Debugf("%#v\n", ctx)
}

func (l *Listener) EnterStart(c *parser.StartContext) {
	t := c.GetText()
	l.Debugf(f1, "enter start", t, l.Wrap)

	l.Lines = make([]*wrap.Wrap, 0)
}

func (l *Listener) EnterLine(c *parser.LineContext) {
	t := c.GetText()
	l.Debugf(f1, "enter line", t, l.Wrap)

	l.Wrap = nil
}

func (l *Listener) EnterFunct(c *parser.FunctContext) {
	t := c.GetText()
	l.Debugf(f1, "enter funct", t, l.Wrap)

	if t == "_" {
		l.Wrap.Funct = &types.Under{}
	} else {
		rs := []rune(t)
		nrs := make([]rune, 0)
		num := 0
		for _, r := range rs {
			if num == 1 {
				num = 0
			}
			if r == '~' && num == 0 {
				num = 1
			} else {
				nrs = append(nrs, r)
			}
		}
		l.Wrap.Funct = types.String(nrs)
	}

	l.Debugf(f2, l.Wrap)
}

func (l *Listener) EnterArgs(c *parser.ArgsContext) {
	t := c.GetText()
	l.Debugf(f1, "enter args", t, l.Wrap)

	args := make(wrap.Args, 0)
	l.Wrap.Argss = append(l.Wrap.Argss, &args)

	l.Debugf(f2, l.Wrap)
}

func (l *Listener) EnterWrap(c *parser.WrapContext) {
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

func (l *Listener) ExitStart(c *parser.StartContext) {
	t := c.GetText()
	l.Debugf(f1, "exit start", t, l.Wrap)

	l.Debugf(f2, l.Wrap)
}

func (l *Listener) ExitLine(c *parser.LineContext) {
	t := c.GetText()
	l.Debugf(f1, "exit line", t, l.Wrap)

	l.Lines = append(l.Lines, l.Wrap)

	l.Debugf(f2, l.Wrap)
}

func (l *Listener) ExitFunct(c *parser.FunctContext) {
	t := c.GetText()
	l.Debugf(f1, "exit funct", t, l.Wrap)

	l.Debugf(f2, l.Wrap)
}

func (l *Listener) ExitArgs(c *parser.ArgsContext) {
	t := c.GetText()
	l.Debugf(f1, "exit args", t, l.Wrap)

	l.Debugf(f2, l.Wrap)
}

func (l *Listener) ExitWrap(c *parser.WrapContext) {
	t := c.GetText()
	l.Debugf(f1, "exit wrap", t, l.Wrap)

	if l.Wrap.Parent != nil {
		l.Wrap = l.Wrap.Parent
	}

	l.Debugf(f2, l.Wrap)
}
