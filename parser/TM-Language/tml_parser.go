// Generated from TM-Language/TML.g4 by ANTLR 4.7.

package parser // TML

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 20, 76, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 2, 3, 3, 6, 3, 25, 10,
	3, 13, 3, 14, 3, 26, 3, 4, 3, 4, 3, 4, 5, 4, 32, 10, 4, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 7, 6, 51, 10, 6, 12, 6, 14, 6, 54, 11, 6, 3, 6, 3, 6, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 2, 2, 11, 2, 4, 6, 8, 10, 12,
	14, 16, 18, 2, 5, 4, 2, 5, 7, 16, 16, 4, 2, 8, 8, 16, 17, 4, 2, 8, 8, 14,
	15, 2, 70, 2, 20, 3, 2, 2, 2, 4, 24, 3, 2, 2, 2, 6, 31, 3, 2, 2, 2, 8,
	33, 3, 2, 2, 2, 10, 45, 3, 2, 2, 2, 12, 57, 3, 2, 2, 2, 14, 69, 3, 2, 2,
	2, 16, 71, 3, 2, 2, 2, 18, 73, 3, 2, 2, 2, 20, 21, 5, 4, 3, 2, 21, 22,
	7, 2, 2, 3, 22, 3, 3, 2, 2, 2, 23, 25, 5, 6, 4, 2, 24, 23, 3, 2, 2, 2,
	25, 26, 3, 2, 2, 2, 26, 24, 3, 2, 2, 2, 26, 27, 3, 2, 2, 2, 27, 5, 3, 2,
	2, 2, 28, 32, 5, 12, 7, 2, 29, 32, 5, 10, 6, 2, 30, 32, 5, 8, 5, 2, 31,
	28, 3, 2, 2, 2, 31, 29, 3, 2, 2, 2, 31, 30, 3, 2, 2, 2, 32, 7, 3, 2, 2,
	2, 33, 34, 7, 9, 2, 2, 34, 35, 5, 14, 8, 2, 35, 36, 7, 13, 2, 2, 36, 37,
	5, 16, 9, 2, 37, 38, 7, 10, 2, 2, 38, 39, 7, 16, 2, 2, 39, 40, 7, 9, 2,
	2, 40, 41, 5, 14, 8, 2, 41, 42, 7, 13, 2, 2, 42, 43, 5, 14, 8, 2, 43, 44,
	7, 10, 2, 2, 44, 9, 3, 2, 2, 2, 45, 46, 7, 3, 2, 2, 46, 47, 7, 4, 2, 2,
	47, 48, 7, 16, 2, 2, 48, 52, 7, 11, 2, 2, 49, 51, 5, 6, 4, 2, 50, 49, 3,
	2, 2, 2, 51, 54, 3, 2, 2, 2, 52, 50, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53,
	55, 3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 55, 56, 7, 12, 2, 2, 56, 11, 3, 2,
	2, 2, 57, 58, 7, 9, 2, 2, 58, 59, 5, 14, 8, 2, 59, 60, 7, 13, 2, 2, 60,
	61, 5, 14, 8, 2, 61, 62, 7, 13, 2, 2, 62, 63, 5, 16, 9, 2, 63, 64, 7, 13,
	2, 2, 64, 65, 5, 16, 9, 2, 65, 66, 7, 13, 2, 2, 66, 67, 5, 18, 10, 2, 67,
	68, 7, 10, 2, 2, 68, 13, 3, 2, 2, 2, 69, 70, 9, 2, 2, 2, 70, 15, 3, 2,
	2, 2, 71, 72, 9, 3, 2, 2, 72, 17, 3, 2, 2, 2, 73, 74, 9, 4, 2, 2, 74, 19,
	3, 2, 2, 2, 5, 26, 31, 52,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'define'", "'macro'", "'hs'", "'ha'", "'hr'", "'_'", "'('", "')'",
	"'{'", "'}'", "','", "'<'", "'>'",
}
var symbolicNames = []string{
	"", "DEFINE", "MACRO", "STARTSTATE", "ACCEPTSTATE", "REJECTSTATE", "UNDERSCORE",
	"LPAREN", "RPAREN", "LBRACE", "RBRACE", "COMMA", "LT", "GT", "ID", "DECIMAL",
	"BlockComment", "LineComment", "WHITESPACE",
}

var ruleNames = []string{
	"start", "program", "statement", "macroApp", "macroDef", "command", "stateLabel",
	"tapeSymbol", "direction",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type TMLParser struct {
	*antlr.BaseParser
}

func NewTMLParser(input antlr.TokenStream) *TMLParser {
	this := new(TMLParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "TML.g4"

	return this
}

// TMLParser tokens.
const (
	TMLParserEOF          = antlr.TokenEOF
	TMLParserDEFINE       = 1
	TMLParserMACRO        = 2
	TMLParserSTARTSTATE   = 3
	TMLParserACCEPTSTATE  = 4
	TMLParserREJECTSTATE  = 5
	TMLParserUNDERSCORE   = 6
	TMLParserLPAREN       = 7
	TMLParserRPAREN       = 8
	TMLParserLBRACE       = 9
	TMLParserRBRACE       = 10
	TMLParserCOMMA        = 11
	TMLParserLT           = 12
	TMLParserGT           = 13
	TMLParserID           = 14
	TMLParserDECIMAL      = 15
	TMLParserBlockComment = 16
	TMLParserLineComment  = 17
	TMLParserWHITESPACE   = 18
)

// TMLParser rules.
const (
	TMLParserRULE_start      = 0
	TMLParserRULE_program    = 1
	TMLParserRULE_statement  = 2
	TMLParserRULE_macroApp   = 3
	TMLParserRULE_macroDef   = 4
	TMLParserRULE_command    = 5
	TMLParserRULE_stateLabel = 6
	TMLParserRULE_tapeSymbol = 7
	TMLParserRULE_direction  = 8
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
	p.RuleIndex = TMLParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TMLParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Program() IProgramContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProgramContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProgramContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(TMLParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TMLListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TMLListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *TMLParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TMLParserRULE_start)

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
		p.SetState(18)
		p.Program()
	}
	{
		p.SetState(19)
		p.Match(TMLParserEOF)
	}

	return localctx
}

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TMLParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TMLParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TMLListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TMLListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *TMLParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TMLParserRULE_program)
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
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == TMLParserDEFINE || _la == TMLParserLPAREN {
		{
			p.SetState(21)
			p.Statement()
		}

		p.SetState(24)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TMLParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TMLParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Command() ICommandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *StatementContext) MacroDef() IMacroDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMacroDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMacroDefContext)
}

func (s *StatementContext) MacroApp() IMacroAppContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMacroAppContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMacroAppContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TMLListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TMLListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *TMLParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TMLParserRULE_statement)

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

	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(26)
			p.Command()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(27)
			p.MacroDef()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(28)
			p.MacroApp()
		}

	}

	return localctx
}

// IMacroAppContext is an interface to support dynamic dispatch.
type IMacroAppContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMacroAppContext differentiates from other interfaces.
	IsMacroAppContext()
}

type MacroAppContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMacroAppContext() *MacroAppContext {
	var p = new(MacroAppContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TMLParserRULE_macroApp
	return p
}

func (*MacroAppContext) IsMacroAppContext() {}

func NewMacroAppContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MacroAppContext {
	var p = new(MacroAppContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TMLParserRULE_macroApp

	return p
}

func (s *MacroAppContext) GetParser() antlr.Parser { return s.parser }

func (s *MacroAppContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(TMLParserLPAREN)
}

func (s *MacroAppContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(TMLParserLPAREN, i)
}

func (s *MacroAppContext) AllStateLabel() []IStateLabelContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStateLabelContext)(nil)).Elem())
	var tst = make([]IStateLabelContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStateLabelContext)
		}
	}

	return tst
}

func (s *MacroAppContext) StateLabel(i int) IStateLabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStateLabelContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStateLabelContext)
}

func (s *MacroAppContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TMLParserCOMMA)
}

func (s *MacroAppContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TMLParserCOMMA, i)
}

func (s *MacroAppContext) TapeSymbol() ITapeSymbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITapeSymbolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITapeSymbolContext)
}

func (s *MacroAppContext) AllRPAREN() []antlr.TerminalNode {
	return s.GetTokens(TMLParserRPAREN)
}

func (s *MacroAppContext) RPAREN(i int) antlr.TerminalNode {
	return s.GetToken(TMLParserRPAREN, i)
}

func (s *MacroAppContext) ID() antlr.TerminalNode {
	return s.GetToken(TMLParserID, 0)
}

func (s *MacroAppContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MacroAppContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MacroAppContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TMLListener); ok {
		listenerT.EnterMacroApp(s)
	}
}

func (s *MacroAppContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TMLListener); ok {
		listenerT.ExitMacroApp(s)
	}
}

func (p *TMLParser) MacroApp() (localctx IMacroAppContext) {
	localctx = NewMacroAppContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TMLParserRULE_macroApp)

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
		p.SetState(31)
		p.Match(TMLParserLPAREN)
	}
	{
		p.SetState(32)
		p.StateLabel()
	}
	{
		p.SetState(33)
		p.Match(TMLParserCOMMA)
	}
	{
		p.SetState(34)
		p.TapeSymbol()
	}
	{
		p.SetState(35)
		p.Match(TMLParserRPAREN)
	}
	{
		p.SetState(36)
		p.Match(TMLParserID)
	}
	{
		p.SetState(37)
		p.Match(TMLParserLPAREN)
	}
	{
		p.SetState(38)
		p.StateLabel()
	}
	{
		p.SetState(39)
		p.Match(TMLParserCOMMA)
	}
	{
		p.SetState(40)
		p.StateLabel()
	}
	{
		p.SetState(41)
		p.Match(TMLParserRPAREN)
	}

	return localctx
}

// IMacroDefContext is an interface to support dynamic dispatch.
type IMacroDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMacroDefContext differentiates from other interfaces.
	IsMacroDefContext()
}

type MacroDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMacroDefContext() *MacroDefContext {
	var p = new(MacroDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TMLParserRULE_macroDef
	return p
}

func (*MacroDefContext) IsMacroDefContext() {}

func NewMacroDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MacroDefContext {
	var p = new(MacroDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TMLParserRULE_macroDef

	return p
}

func (s *MacroDefContext) GetParser() antlr.Parser { return s.parser }

func (s *MacroDefContext) DEFINE() antlr.TerminalNode {
	return s.GetToken(TMLParserDEFINE, 0)
}

func (s *MacroDefContext) MACRO() antlr.TerminalNode {
	return s.GetToken(TMLParserMACRO, 0)
}

func (s *MacroDefContext) ID() antlr.TerminalNode {
	return s.GetToken(TMLParserID, 0)
}

func (s *MacroDefContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(TMLParserLBRACE, 0)
}

func (s *MacroDefContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(TMLParserRBRACE, 0)
}

func (s *MacroDefContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *MacroDefContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *MacroDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MacroDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MacroDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TMLListener); ok {
		listenerT.EnterMacroDef(s)
	}
}

func (s *MacroDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TMLListener); ok {
		listenerT.ExitMacroDef(s)
	}
}

func (p *TMLParser) MacroDef() (localctx IMacroDefContext) {
	localctx = NewMacroDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TMLParserRULE_macroDef)
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
		p.SetState(43)
		p.Match(TMLParserDEFINE)
	}
	{
		p.SetState(44)
		p.Match(TMLParserMACRO)
	}
	{
		p.SetState(45)
		p.Match(TMLParserID)
	}
	{
		p.SetState(46)
		p.Match(TMLParserLBRACE)
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TMLParserDEFINE || _la == TMLParserLPAREN {
		{
			p.SetState(47)
			p.Statement()
		}

		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(53)
		p.Match(TMLParserRBRACE)
	}

	return localctx
}

// ICommandContext is an interface to support dynamic dispatch.
type ICommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandContext() *CommandContext {
	var p = new(CommandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TMLParserRULE_command
	return p
}

func (*CommandContext) IsCommandContext() {}

func NewCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandContext {
	var p = new(CommandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TMLParserRULE_command

	return p
}

func (s *CommandContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TMLParserLPAREN, 0)
}

func (s *CommandContext) AllStateLabel() []IStateLabelContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStateLabelContext)(nil)).Elem())
	var tst = make([]IStateLabelContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStateLabelContext)
		}
	}

	return tst
}

func (s *CommandContext) StateLabel(i int) IStateLabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStateLabelContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStateLabelContext)
}

func (s *CommandContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TMLParserCOMMA)
}

func (s *CommandContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TMLParserCOMMA, i)
}

func (s *CommandContext) AllTapeSymbol() []ITapeSymbolContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITapeSymbolContext)(nil)).Elem())
	var tst = make([]ITapeSymbolContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITapeSymbolContext)
		}
	}

	return tst
}

func (s *CommandContext) TapeSymbol(i int) ITapeSymbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITapeSymbolContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITapeSymbolContext)
}

func (s *CommandContext) Direction() IDirectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectionContext)
}

func (s *CommandContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TMLParserRPAREN, 0)
}

func (s *CommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TMLListener); ok {
		listenerT.EnterCommand(s)
	}
}

func (s *CommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TMLListener); ok {
		listenerT.ExitCommand(s)
	}
}

func (p *TMLParser) Command() (localctx ICommandContext) {
	localctx = NewCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TMLParserRULE_command)

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
		p.SetState(55)
		p.Match(TMLParserLPAREN)
	}
	{
		p.SetState(56)
		p.StateLabel()
	}
	{
		p.SetState(57)
		p.Match(TMLParserCOMMA)
	}
	{
		p.SetState(58)
		p.StateLabel()
	}
	{
		p.SetState(59)
		p.Match(TMLParserCOMMA)
	}
	{
		p.SetState(60)
		p.TapeSymbol()
	}
	{
		p.SetState(61)
		p.Match(TMLParserCOMMA)
	}
	{
		p.SetState(62)
		p.TapeSymbol()
	}
	{
		p.SetState(63)
		p.Match(TMLParserCOMMA)
	}
	{
		p.SetState(64)
		p.Direction()
	}
	{
		p.SetState(65)
		p.Match(TMLParserRPAREN)
	}

	return localctx
}

// IStateLabelContext is an interface to support dynamic dispatch.
type IStateLabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStateLabelContext differentiates from other interfaces.
	IsStateLabelContext()
}

type StateLabelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStateLabelContext() *StateLabelContext {
	var p = new(StateLabelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TMLParserRULE_stateLabel
	return p
}

func (*StateLabelContext) IsStateLabelContext() {}

func NewStateLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StateLabelContext {
	var p = new(StateLabelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TMLParserRULE_stateLabel

	return p
}

func (s *StateLabelContext) GetParser() antlr.Parser { return s.parser }

func (s *StateLabelContext) ID() antlr.TerminalNode {
	return s.GetToken(TMLParserID, 0)
}

func (s *StateLabelContext) STARTSTATE() antlr.TerminalNode {
	return s.GetToken(TMLParserSTARTSTATE, 0)
}

func (s *StateLabelContext) REJECTSTATE() antlr.TerminalNode {
	return s.GetToken(TMLParserREJECTSTATE, 0)
}

func (s *StateLabelContext) ACCEPTSTATE() antlr.TerminalNode {
	return s.GetToken(TMLParserACCEPTSTATE, 0)
}

func (s *StateLabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StateLabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StateLabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TMLListener); ok {
		listenerT.EnterStateLabel(s)
	}
}

func (s *StateLabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TMLListener); ok {
		listenerT.ExitStateLabel(s)
	}
}

func (p *TMLParser) StateLabel() (localctx IStateLabelContext) {
	localctx = NewStateLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TMLParserRULE_stateLabel)
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
	p.SetState(67)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TMLParserSTARTSTATE)|(1<<TMLParserACCEPTSTATE)|(1<<TMLParserREJECTSTATE)|(1<<TMLParserID))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// ITapeSymbolContext is an interface to support dynamic dispatch.
type ITapeSymbolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTapeSymbolContext differentiates from other interfaces.
	IsTapeSymbolContext()
}

type TapeSymbolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTapeSymbolContext() *TapeSymbolContext {
	var p = new(TapeSymbolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TMLParserRULE_tapeSymbol
	return p
}

func (*TapeSymbolContext) IsTapeSymbolContext() {}

func NewTapeSymbolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TapeSymbolContext {
	var p = new(TapeSymbolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TMLParserRULE_tapeSymbol

	return p
}

func (s *TapeSymbolContext) GetParser() antlr.Parser { return s.parser }

func (s *TapeSymbolContext) ID() antlr.TerminalNode {
	return s.GetToken(TMLParserID, 0)
}

func (s *TapeSymbolContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(TMLParserDECIMAL, 0)
}

func (s *TapeSymbolContext) UNDERSCORE() antlr.TerminalNode {
	return s.GetToken(TMLParserUNDERSCORE, 0)
}

func (s *TapeSymbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TapeSymbolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TapeSymbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TMLListener); ok {
		listenerT.EnterTapeSymbol(s)
	}
}

func (s *TapeSymbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TMLListener); ok {
		listenerT.ExitTapeSymbol(s)
	}
}

func (p *TMLParser) TapeSymbol() (localctx ITapeSymbolContext) {
	localctx = NewTapeSymbolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TMLParserRULE_tapeSymbol)
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
	p.SetState(69)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TMLParserUNDERSCORE)|(1<<TMLParserID)|(1<<TMLParserDECIMAL))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IDirectionContext is an interface to support dynamic dispatch.
type IDirectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectionContext differentiates from other interfaces.
	IsDirectionContext()
}

type DirectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectionContext() *DirectionContext {
	var p = new(DirectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TMLParserRULE_direction
	return p
}

func (*DirectionContext) IsDirectionContext() {}

func NewDirectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectionContext {
	var p = new(DirectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TMLParserRULE_direction

	return p
}

func (s *DirectionContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectionContext) LT() antlr.TerminalNode {
	return s.GetToken(TMLParserLT, 0)
}

func (s *DirectionContext) GT() antlr.TerminalNode {
	return s.GetToken(TMLParserGT, 0)
}

func (s *DirectionContext) UNDERSCORE() antlr.TerminalNode {
	return s.GetToken(TMLParserUNDERSCORE, 0)
}

func (s *DirectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TMLListener); ok {
		listenerT.EnterDirection(s)
	}
}

func (s *DirectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TMLListener); ok {
		listenerT.ExitDirection(s)
	}
}

func (p *TMLParser) Direction() (localctx IDirectionContext) {
	localctx = NewDirectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TMLParserRULE_direction)
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
	p.SetState(71)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TMLParserUNDERSCORE)|(1<<TMLParserLT)|(1<<TMLParserGT))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}
