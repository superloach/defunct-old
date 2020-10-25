package parse

import (
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/participle/lexer/regex"
)

var Lexer = lexer.Must(regex.New(`
	Under = [_]
	ArgsOpen = [[]
	ArgsClose = [\]]
	BlockOpen = [{]
	BlockClose = [}]
	Space = [ \t\r\n]+
	Text = ([~][_[\]{} \t\r\n]|[^_[\]{} \t\r\n])+
`))
