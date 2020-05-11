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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 10, 41, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 7, 3, 7, 5, 7, 32, 10, 7, 3, 8, 3, 8, 3, 9, 3, 9, 6,
	9, 38, 10, 9, 13, 9, 14, 9, 39, 2, 2, 10, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7,
	13, 8, 15, 9, 17, 10, 3, 2, 4, 4, 2, 11, 12, 34, 34, 8, 2, 11, 12, 34,
	34, 93, 93, 95, 95, 97, 97, 128, 128, 2, 42, 2, 3, 3, 2, 2, 2, 2, 5, 3,
	2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13,
	3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 3, 19, 3, 2, 2, 2, 5,
	21, 3, 2, 2, 2, 7, 23, 3, 2, 2, 2, 9, 25, 3, 2, 2, 2, 11, 27, 3, 2, 2,
	2, 13, 29, 3, 2, 2, 2, 15, 33, 3, 2, 2, 2, 17, 37, 3, 2, 2, 2, 19, 20,
	7, 97, 2, 2, 20, 4, 3, 2, 2, 2, 21, 22, 7, 93, 2, 2, 22, 6, 3, 2, 2, 2,
	23, 24, 7, 95, 2, 2, 24, 8, 3, 2, 2, 2, 25, 26, 7, 128, 2, 2, 26, 10, 3,
	2, 2, 2, 27, 28, 9, 2, 2, 2, 28, 12, 3, 2, 2, 2, 29, 31, 5, 9, 5, 2, 30,
	32, 9, 3, 2, 2, 31, 30, 3, 2, 2, 2, 32, 14, 3, 2, 2, 2, 33, 34, 10, 3,
	2, 2, 34, 16, 3, 2, 2, 2, 35, 38, 5, 13, 7, 2, 36, 38, 5, 15, 8, 2, 37,
	35, 3, 2, 2, 2, 37, 36, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 37, 3, 2, 2,
	2, 39, 40, 3, 2, 2, 2, 40, 18, 3, 2, 2, 2, 6, 2, 31, 37, 39, 2,
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
	"", "UNDER", "OPEN", "CLOSE", "ESC", "WS", "ESCD", "CHAR", "STR",
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
	DefunctLexerESCD  = 6
	DefunctLexerCHAR  = 7
	DefunctLexerSTR   = 8
)
