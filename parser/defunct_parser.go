// Code generated from Defunct.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Defunct

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 8, 35, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 2, 3, 2, 5,
	2, 15, 10, 2, 3, 3, 3, 3, 3, 4, 3, 4, 7, 4, 21, 10, 4, 12, 4, 14, 4, 24,
	11, 4, 3, 4, 3, 4, 3, 5, 3, 5, 7, 5, 30, 10, 5, 12, 5, 14, 5, 33, 11, 5,
	3, 5, 2, 2, 6, 2, 4, 6, 8, 2, 3, 4, 2, 3, 3, 6, 6, 2, 33, 2, 14, 3, 2,
	2, 2, 4, 16, 3, 2, 2, 2, 6, 18, 3, 2, 2, 2, 8, 27, 3, 2, 2, 2, 10, 15,
	7, 2, 2, 3, 11, 12, 5, 8, 5, 2, 12, 13, 7, 2, 2, 3, 13, 15, 3, 2, 2, 2,
	14, 10, 3, 2, 2, 2, 14, 11, 3, 2, 2, 2, 15, 3, 3, 2, 2, 2, 16, 17, 9, 2,
	2, 2, 17, 5, 3, 2, 2, 2, 18, 22, 7, 4, 2, 2, 19, 21, 5, 8, 5, 2, 20, 19,
	3, 2, 2, 2, 21, 24, 3, 2, 2, 2, 22, 20, 3, 2, 2, 2, 22, 23, 3, 2, 2, 2,
	23, 25, 3, 2, 2, 2, 24, 22, 3, 2, 2, 2, 25, 26, 7, 5, 2, 2, 26, 7, 3, 2,
	2, 2, 27, 31, 5, 4, 3, 2, 28, 30, 5, 6, 4, 2, 29, 28, 3, 2, 2, 2, 30, 33,
	3, 2, 2, 2, 31, 29, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 9, 3, 2, 2, 2,
	33, 31, 3, 2, 2, 2, 5, 14, 22, 31,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'_'", "'['", "']'",
}
var symbolicNames = []string{
	"", "UNDER", "OPEN", "CLOSE", "STR", "NEWLINE", "WHITESPACE",
}

var ruleNames = []string{
	"start", "funct", "args", "wrap",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type DefunctParser struct {
	*antlr.BaseParser
}

func NewDefunctParser(input antlr.TokenStream) *DefunctParser {
	this := new(DefunctParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Defunct.g4"

	return this
}

// DefunctParser tokens.
const (
	DefunctParserEOF        = antlr.TokenEOF
	DefunctParserUNDER      = 1
	DefunctParserOPEN       = 2
	DefunctParserCLOSE      = 3
	DefunctParserSTR        = 4
	DefunctParserNEWLINE    = 5
	DefunctParserWHITESPACE = 6
)

// DefunctParser rules.
const (
	DefunctParserRULE_start = 0
	DefunctParserRULE_funct = 1
	DefunctParserRULE_args  = 2
	DefunctParserRULE_wrap  = 3
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefunctParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefunctParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(DefunctParserEOF, 0)
}

func (s *StartContext) Wrap() IWrapContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWrapContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWrapContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefunctListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefunctListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *DefunctParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DefunctParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(12)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DefunctParserEOF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(8)
			p.Match(DefunctParserEOF)
		}

	case DefunctParserUNDER, DefunctParserSTR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(9)
			p.Wrap()
		}
		{
			p.SetState(10)
			p.Match(DefunctParserEOF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctContext is an interface to support dynamic dispatch.
type IFunctContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctContext differentiates from other interfaces.
	IsFunctContext()
}

type FunctContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctContext() *FunctContext {
	var p = new(FunctContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefunctParserRULE_funct
	return p
}

func (*FunctContext) IsFunctContext() {}

func NewFunctContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctContext {
	var p = new(FunctContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefunctParserRULE_funct

	return p
}

func (s *FunctContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctContext) UNDER() antlr.TerminalNode {
	return s.GetToken(DefunctParserUNDER, 0)
}

func (s *FunctContext) STR() antlr.TerminalNode {
	return s.GetToken(DefunctParserSTR, 0)
}

func (s *FunctContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefunctListener); ok {
		listenerT.EnterFunct(s)
	}
}

func (s *FunctContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefunctListener); ok {
		listenerT.ExitFunct(s)
	}
}

func (p *DefunctParser) Funct() (localctx IFunctContext) {
	localctx = NewFunctContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DefunctParserRULE_funct)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(14)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DefunctParserUNDER || _la == DefunctParserSTR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IArgsContext is an interface to support dynamic dispatch.
type IArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgsContext differentiates from other interfaces.
	IsArgsContext()
}

type ArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgsContext() *ArgsContext {
	var p = new(ArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefunctParserRULE_args
	return p
}

func (*ArgsContext) IsArgsContext() {}

func NewArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsContext {
	var p = new(ArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefunctParserRULE_args

	return p
}

func (s *ArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgsContext) OPEN() antlr.TerminalNode {
	return s.GetToken(DefunctParserOPEN, 0)
}

func (s *ArgsContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(DefunctParserCLOSE, 0)
}

func (s *ArgsContext) AllWrap() []IWrapContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWrapContext)(nil)).Elem())
	var tst = make([]IWrapContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWrapContext)
		}
	}

	return tst
}

func (s *ArgsContext) Wrap(i int) IWrapContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWrapContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWrapContext)
}

func (s *ArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefunctListener); ok {
		listenerT.EnterArgs(s)
	}
}

func (s *ArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefunctListener); ok {
		listenerT.ExitArgs(s)
	}
}

func (p *DefunctParser) Args() (localctx IArgsContext) {
	localctx = NewArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DefunctParserRULE_args)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.Match(DefunctParserOPEN)
	}
	p.SetState(20)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DefunctParserUNDER || _la == DefunctParserSTR {
		{
			p.SetState(17)
			p.Wrap()
		}

		p.SetState(22)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(23)
		p.Match(DefunctParserCLOSE)
	}

	return localctx
}

// IWrapContext is an interface to support dynamic dispatch.
type IWrapContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWrapContext differentiates from other interfaces.
	IsWrapContext()
}

type WrapContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWrapContext() *WrapContext {
	var p = new(WrapContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefunctParserRULE_wrap
	return p
}

func (*WrapContext) IsWrapContext() {}

func NewWrapContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WrapContext {
	var p = new(WrapContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefunctParserRULE_wrap

	return p
}

func (s *WrapContext) GetParser() antlr.Parser { return s.parser }

func (s *WrapContext) Funct() IFunctContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctContext)
}

func (s *WrapContext) AllArgs() []IArgsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgsContext)(nil)).Elem())
	var tst = make([]IArgsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgsContext)
		}
	}

	return tst
}

func (s *WrapContext) Args(i int) IArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *WrapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WrapContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WrapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefunctListener); ok {
		listenerT.EnterWrap(s)
	}
}

func (s *WrapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefunctListener); ok {
		listenerT.ExitWrap(s)
	}
}

func (p *DefunctParser) Wrap() (localctx IWrapContext) {
	localctx = NewWrapContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DefunctParserRULE_wrap)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)
		p.Funct()
	}
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DefunctParserOPEN {
		{
			p.SetState(26)
			p.Args()
		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
