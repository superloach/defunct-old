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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 8, 55, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 7,
	2, 15, 10, 2, 12, 2, 14, 2, 18, 11, 2, 7, 2, 20, 10, 2, 12, 2, 14, 2, 23,
	11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 7, 4, 31, 10, 4, 12, 4, 14,
	4, 34, 11, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 6, 6, 41, 10, 6, 13, 6, 14,
	6, 42, 3, 6, 7, 6, 46, 10, 6, 12, 6, 14, 6, 49, 11, 6, 5, 6, 51, 10, 6,
	3, 6, 3, 6, 3, 6, 2, 2, 7, 2, 4, 6, 8, 10, 2, 3, 4, 2, 3, 3, 8, 8, 2, 55,
	2, 21, 3, 2, 2, 2, 4, 26, 3, 2, 2, 2, 6, 28, 3, 2, 2, 2, 8, 35, 3, 2, 2,
	2, 10, 37, 3, 2, 2, 2, 12, 16, 5, 4, 3, 2, 13, 15, 7, 7, 2, 2, 14, 13,
	3, 2, 2, 2, 15, 18, 3, 2, 2, 2, 16, 14, 3, 2, 2, 2, 16, 17, 3, 2, 2, 2,
	17, 20, 3, 2, 2, 2, 18, 16, 3, 2, 2, 2, 19, 12, 3, 2, 2, 2, 20, 23, 3,
	2, 2, 2, 21, 19, 3, 2, 2, 2, 21, 22, 3, 2, 2, 2, 22, 24, 3, 2, 2, 2, 23,
	21, 3, 2, 2, 2, 24, 25, 7, 2, 2, 3, 25, 3, 3, 2, 2, 2, 26, 27, 5, 6, 4,
	2, 27, 5, 3, 2, 2, 2, 28, 32, 5, 8, 5, 2, 29, 31, 5, 10, 6, 2, 30, 29,
	3, 2, 2, 2, 31, 34, 3, 2, 2, 2, 32, 30, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2,
	33, 7, 3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 35, 36, 9, 2, 2, 2, 36, 9, 3, 2,
	2, 2, 37, 50, 7, 4, 2, 2, 38, 47, 5, 6, 4, 2, 39, 41, 7, 7, 2, 2, 40, 39,
	3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2,
	43, 44, 3, 2, 2, 2, 44, 46, 5, 6, 4, 2, 45, 40, 3, 2, 2, 2, 46, 49, 3,
	2, 2, 2, 47, 45, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 51, 3, 2, 2, 2, 49,
	47, 3, 2, 2, 2, 50, 38, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 52, 3, 2, 2,
	2, 52, 53, 7, 5, 2, 2, 53, 11, 3, 2, 2, 2, 8, 16, 21, 32, 42, 47, 50,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'_'", "'['", "']'", "'~'",
}
var symbolicNames = []string{
	"", "UNDER", "OPEN", "CLOSE", "ESC", "WS", "STR",
}

var ruleNames = []string{
	"start", "line", "wrap", "funct", "args",
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
	DefunctParserEOF   = antlr.TokenEOF
	DefunctParserUNDER = 1
	DefunctParserOPEN  = 2
	DefunctParserCLOSE = 3
	DefunctParserESC   = 4
	DefunctParserWS    = 5
	DefunctParserSTR   = 6
)

// DefunctParser rules.
const (
	DefunctParserRULE_start = 0
	DefunctParserRULE_line  = 1
	DefunctParserRULE_wrap  = 2
	DefunctParserRULE_funct = 3
	DefunctParserRULE_args  = 4
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

func (s *StartContext) AllLine() []ILineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineContext)(nil)).Elem())
	var tst = make([]ILineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineContext)
		}
	}

	return tst
}

func (s *StartContext) Line(i int) ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *StartContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(DefunctParserWS)
}

func (s *StartContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(DefunctParserWS, i)
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
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DefunctParserUNDER || _la == DefunctParserSTR {
		{
			p.SetState(10)
			p.Line()
		}
		p.SetState(14)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == DefunctParserWS {
			{
				p.SetState(11)
				p.Match(DefunctParserWS)
			}

			p.SetState(16)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(22)
		p.Match(DefunctParserEOF)
	}

	return localctx
}

// ILineContext is an interface to support dynamic dispatch.
type ILineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLineContext differentiates from other interfaces.
	IsLineContext()
}

type LineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineContext() *LineContext {
	var p = new(LineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefunctParserRULE_line
	return p
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefunctParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) Wrap() IWrapContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWrapContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWrapContext)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefunctListener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefunctListener); ok {
		listenerT.ExitLine(s)
	}
}

func (p *DefunctParser) Line() (localctx ILineContext) {
	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DefunctParserRULE_line)

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
		p.SetState(24)
		p.Wrap()
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
	p.EnterRule(localctx, 4, DefunctParserRULE_wrap)
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
		p.SetState(26)
		p.Funct()
	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DefunctParserOPEN {
		{
			p.SetState(27)
			p.Args()
		}

		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.EnterRule(localctx, 6, DefunctParserRULE_funct)
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
		p.SetState(33)
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

func (s *ArgsContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(DefunctParserWS)
}

func (s *ArgsContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(DefunctParserWS, i)
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
	p.EnterRule(localctx, 8, DefunctParserRULE_args)
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
		p.SetState(35)
		p.Match(DefunctParserOPEN)
	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DefunctParserUNDER || _la == DefunctParserSTR {
		{
			p.SetState(36)
			p.Wrap()
		}
		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == DefunctParserWS {
			p.SetState(38)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == DefunctParserWS {
				{
					p.SetState(37)
					p.Match(DefunctParserWS)
				}

				p.SetState(40)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(42)
				p.Wrap()
			}

			p.SetState(47)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(50)
		p.Match(DefunctParserCLOSE)
	}

	return localctx
}
