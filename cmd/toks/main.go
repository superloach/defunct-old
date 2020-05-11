package main

import (
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/superloach/defunct/parser"
)

func main() {
	// Setup the input
	is := antlr.NewInputStream(os.Args[1])

	// Create the Lexer
	lexer := parser.NewDefunctLexer(is)

	// Read all tokens
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}
