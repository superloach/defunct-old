// Code generated from Defunct.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Defunct

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDefunctListener is a complete listener for a parse tree produced by DefunctParser.
type BaseDefunctListener struct{}

var _ DefunctListener = &BaseDefunctListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDefunctListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDefunctListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDefunctListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDefunctListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseDefunctListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseDefunctListener) ExitStart(ctx *StartContext) {}

// EnterLine is called when production line is entered.
func (s *BaseDefunctListener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BaseDefunctListener) ExitLine(ctx *LineContext) {}

// EnterWrap is called when production wrap is entered.
func (s *BaseDefunctListener) EnterWrap(ctx *WrapContext) {}

// ExitWrap is called when production wrap is exited.
func (s *BaseDefunctListener) ExitWrap(ctx *WrapContext) {}

// EnterFunct is called when production funct is entered.
func (s *BaseDefunctListener) EnterFunct(ctx *FunctContext) {}

// ExitFunct is called when production funct is exited.
func (s *BaseDefunctListener) ExitFunct(ctx *FunctContext) {}

// EnterArgs is called when production args is entered.
func (s *BaseDefunctListener) EnterArgs(ctx *ArgsContext) {}

// ExitArgs is called when production args is exited.
func (s *BaseDefunctListener) ExitArgs(ctx *ArgsContext) {}
