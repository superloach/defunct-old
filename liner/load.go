package liner

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/superloach/defunct/parser"
)

func LoadLine(
	line string,
	debugf func(string, ...interface{}),
) *Liner {
	is := antlr.NewInputStream(line)

	lexer := parser.NewDefunctLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewDefunctParser(stream)

	l := &Liner{
		Debugf: debugf,
	}

	antlr.ParseTreeWalkerDefault.Walk(l, p.Start())

	return l
}
