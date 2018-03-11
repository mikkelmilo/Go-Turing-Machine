package parser

/*
	This file contains tools for working with the generated parser.
	This includes listeners for syntax Errors, and pretty printing of a TML AST.
*/

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strconv"
)

type syntaxError struct {
	line   int
	column int
	msg    string
}

func (s syntaxError) String() string {
	line_str := strconv.Itoa(s.line)
	column_str := strconv.Itoa(s.column)
	return s.msg + " in line " + line_str + ", column " + column_str
}

func (s syntaxError) Error() string {
	return s.String()
}

type TMLTreePrinterListener struct {
	*BaseTMLListener
	nesting_lvl int
	errors      []antlr.ErrorNode
}

func nestingLvlString(lvl int) string {
	str := ""
	for i := 0; i < lvl; i++ {
		str += "  "
	}
	return str
}

/*
func (t *TMLTreePrinterListener) VisitTerminal(node antlr.TerminalNode) {
	panic("implement me")
}

func (t *TMLTreePrinterListener) VisitErrorNode(node antlr.ErrorNode) {
}

func (t *TMLTreePrinterListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	panic("implement me")
}

func (t *TMLTreePrinterListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	panic("implement me")
}
*/
func (t *TMLTreePrinterListener) EnterStart(c *StartContext) {
	t.nesting_lvl = 0
}

func (t *TMLTreePrinterListener) EnterProgram(c *ProgramContext) {
	println("Program:")
	t.nesting_lvl++

}

func (t *TMLTreePrinterListener) EnterStatement(c *StatementContext) {
}

func (t *TMLTreePrinterListener) EnterMacroApp(c *MacroAppContext) {
	print(nestingLvlString(t.nesting_lvl) + "Macro Application: (")
}

func (t *TMLTreePrinterListener) EnterMacroDef(c *MacroDefContext) {
	println(nestingLvlString(t.nesting_lvl) + "Macro Definition: " +
		c.GetToken(TMLLexerID, 0).GetText())
	t.nesting_lvl++
}

func (t *TMLTreePrinterListener) EnterCommand(c *CommandContext) {
	print(nestingLvlString(t.nesting_lvl) + "Command: (")
}

func (t *TMLTreePrinterListener) EnterStateLabel(c *StateLabelContext) {
	print(c.GetText() + " ")
}

func (t *TMLTreePrinterListener) EnterTapeSymbol(c *TapeSymbolContext) {
	print(c.GetText() + " ")
}

func (t *TMLTreePrinterListener) EnterDirection(c *DirectionContext) {
	print(c.GetText() + " ")
}

func (t *TMLTreePrinterListener) ExitStart(c *StartContext) {
}

func (t *TMLTreePrinterListener) ExitProgram(c *ProgramContext) {
	t.nesting_lvl--
}

func (t *TMLTreePrinterListener) ExitStatement(c *StatementContext) {
}

func (t *TMLTreePrinterListener) ExitMacroApp(c *MacroAppContext) {
	println(")")
}

func (t *TMLTreePrinterListener) ExitMacroDef(c *MacroDefContext) {
	t.nesting_lvl--
}

func (t *TMLTreePrinterListener) ExitCommand(c *CommandContext) {
	println(")")
}

func (t *TMLTreePrinterListener) ExitStateLabel(c *StateLabelContext) {
}

func (t *TMLTreePrinterListener) ExitTapeSymbol(c *TapeSymbolContext) {
}

func (t *TMLTreePrinterListener) ExitDirection(c *DirectionContext) {
}

type TMLerrorListener struct {
	Errors []syntaxError
}

func (el *TMLerrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{},
	line, column int, msg string, e antlr.RecognitionException) {
	el.Errors = append(el.Errors, syntaxError{line: line, column: column, msg: msg})
}

// we assert this never happens because this is inherently a property of the grammar, not the specific instance.
func (el *TMLerrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	panic("implement me")
}

// we assert this never happens because this is inherently a property of the grammar, not the specific instance.
func (el *TMLerrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	panic("implement me")
}

// we assert this never happens because this is inherently a property of the grammar, not the specific instance.
func (el *TMLerrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	panic("implement me")
}
