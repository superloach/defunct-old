package defunct

import (
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/superloach/defunct/functs"
	"github.com/superloach/defunct/parser"
)

type Arena struct {
	Under functs.Funct
	Debug bool
}

func NewArena(in functs.Funct) *Arena {
	a := &Arena{}

	a.Debug = false
	a.Under = &functs.Under{
		In:    in,
		Out:   functs.Zilch,
		Debug: &(a.Debug),
	}

	return a
}

func (a *Arena) Debugf(f string, args ...interface{}) {
	if a.Debug {
		fmt.Printf(f, args...)
	}
}

func (a *Arena) RunBytes(code []byte) (functs.Funct, error) {
	return a.RunString(string(code))
}

func (a *Arena) RunString(code string) (functs.Funct, error) {
	for _, line := range strings.Split(string(code), "\n") {
		_, err := a.RunLine(line)
		if err != nil {
			return a.Under.GetProp("out"), err
		}
	}

	return a.Under.GetProp("out"), nil
}

func (a *Arena) RunLine(line string) (functs.Funct, error) {
	is := antlr.NewInputStream(line)

	lexer := parser.NewDefunctLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewDefunctParser(stream)

	liner := &Liner{Arena: a}

	antlr.ParseTreeWalkerDefault.Walk(liner, p.Start())

	a.Debugf("%s\n", liner)

	return liner.Wrap.Run()
}
