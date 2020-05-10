package defunct

import (
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/superloach/defunct/debug"
	"github.com/superloach/defunct/functs"
	"github.com/superloach/defunct/parser"
)

type Arena struct{}

func (a *Arena) RunProg(code string, in functs.Thing) (*functs.Under, error) {
	under := &functs.Under{
		In: in,
	}

	for _, line := range strings.Split(string(code), "\n") {
		is := antlr.NewInputStream(line)

		lexer := parser.NewDefunctLexer(is)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		p := parser.NewDefunctParser(stream)

		dl := &Listener{
			Under: under,
		}

		antlr.ParseTreeWalkerDefault.Walk(dl, p.Start())

		debug.Debugf("%s\n", dl)

		_, err := dl.Wrap.Run()
		if err != nil {
			return under, err
		}
	}

	return under, nil
}
