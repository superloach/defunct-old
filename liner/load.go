package liner

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/superloach/defunct/parser"
)

func LoadLines(
	code string,
	debugf func(string, ...interface{}),
) *Liner {
	is := antlr.NewInputStream(code)

	lexer := parser.NewDefunctLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewDefunctParser(stream)

	l := &Liner{
		Debugf: debugf,
	}

	debugf("parsing\n")
	antlr.ParseTreeWalkerDefault.Walk(l, p.Start())
	debugf("parsed\n")

	return l
}
