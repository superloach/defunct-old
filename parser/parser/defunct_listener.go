// Code generated from Defunct.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Defunct

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DefunctListener is a complete listener for a parse tree produced by DefunctParser.
type DefunctListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterWrap is called when entering the wrap production.
	EnterWrap(c *WrapContext)

	// EnterFunct is called when entering the funct production.
	EnterFunct(c *FunctContext)

	// EnterArgs is called when entering the args production.
	EnterArgs(c *ArgsContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitWrap is called when exiting the wrap production.
	ExitWrap(c *WrapContext)

	// ExitFunct is called when exiting the funct production.
	ExitFunct(c *FunctContext)

	// ExitArgs is called when exiting the args production.
	ExitArgs(c *ArgsContext)
}
