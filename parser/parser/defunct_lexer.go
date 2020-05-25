// Code generated from Defunct.g4 by ANTLR 4.8. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 8, 40, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 6, 9, 37, 10,
	9, 13, 9, 14, 9, 38, 2, 2, 10, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 2, 15,
	2, 17, 8, 3, 2, 4, 4, 2, 11, 12, 34, 34, 8, 2, 11, 12, 34, 34, 93, 93,
	95, 95, 97, 97, 128, 128, 2, 39, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2,
	7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2,
	3, 19, 3, 2, 2, 2, 5, 21, 3, 2, 2, 2, 7, 23, 3, 2, 2, 2, 9, 25, 3, 2, 2,
	2, 11, 27, 3, 2, 2, 2, 13, 29, 3, 2, 2, 2, 15, 32, 3, 2, 2, 2, 17, 36,
	3, 2, 2, 2, 19, 20, 7, 97, 2, 2, 20, 4, 3, 2, 2, 2, 21, 22, 7, 93, 2, 2,
	22, 6, 3, 2, 2, 2, 23, 24, 7, 95, 2, 2, 24, 8, 3, 2, 2, 2, 25, 26, 7, 128,
	2, 2, 26, 10, 3, 2, 2, 2, 27, 28, 9, 2, 2, 2, 28, 12, 3, 2, 2, 2, 29, 30,
	5, 9, 5, 2, 30, 31, 9, 3, 2, 2, 31, 14, 3, 2, 2, 2, 32, 33, 10, 3, 2, 2,
	33, 16, 3, 2, 2, 2, 34, 37, 5, 13, 7, 2, 35, 37, 5, 15, 8, 2, 36, 34, 3,
	2, 2, 2, 36, 35, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 38,
	39, 3, 2, 2, 2, 39, 18, 3, 2, 2, 2, 5, 2, 36, 38, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'_'", "'['", "']'", "'~'",
}

var lexerSymbolicNames = []string{
	"", "UNDER", "OPEN", "CLOSE", "ESC", "WS", "STR",
}

var lexerRuleNames = []string{
	"UNDER", "OPEN", "CLOSE", "ESC", "WS", "ESCD", "CHAR", "STR",
}

type DefunctLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewDefunctLexer(input antlr.CharStream) *DefunctLexer {

	l := new(DefunctLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Defunct.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// DefunctLexer tokens.
const (
	DefunctLexerUNDER = 1
	DefunctLexerOPEN  = 2
	DefunctLexerCLOSE = 3
	DefunctLexerESC   = 4
	DefunctLexerWS    = 5
	DefunctLexerSTR   = 6
)
