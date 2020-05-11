package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/superloach/defunct/parser/parser"
)

func ParseLines(
	code string,
	debugf func(string, ...interface{}),
) *Listener {
	is := antlr.NewInputStream(code)

	lexer := parser.NewDefunctLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewDefunctParser(stream)

	l := &Listener{
		Debugf: debugf,
	}

	debugf("parsing\n")
	antlr.ParseTreeWalkerDefault.Walk(l, p.Start())
	debugf("parsed\n")

	return l
}
