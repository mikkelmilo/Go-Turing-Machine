package Compiler

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mikkelmilo/Go-Turing-Machine/Compiler/antlr-parser"
	"strconv"
)

type TMLError struct {
	line   int
	column int
	msg    string
}

func (s TMLError) String() string {
	line_str := strconv.Itoa(s.line)
	column_str := strconv.Itoa(s.column)
	return s.msg + " in line " + line_str + ", column " + column_str
}

func (s TMLError) Error() string {
	return s.String()
}

type Command struct {
	CurrentState  string
	NewState      string
	CurrentSymbol string
	NewSymbol     string
	Direction     string
}

/*
	This TMLTreeListener pretty prints the syntax tree as it traverses the tree.
*/
type TMLTreePrinterListener struct {
	*parser.BaseTMLListener
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

func (t *TMLTreePrinterListener) EnterProgram(c *parser.ProgramContext) {
	t.nesting_lvl = 0
	println("Program:")
	t.nesting_lvl++
}

func (t *TMLTreePrinterListener) EnterMacroApp(c *parser.MacroAppContext) {
	name := c.GetToken(parser.TMLLexerID, 0).GetText()
	print(nestingLvlString(t.nesting_lvl) + "Macro Application: " + name + "(")
}

func (t *TMLTreePrinterListener) EnterMacroDef(c *parser.MacroDefContext) {
	println(nestingLvlString(t.nesting_lvl) + "Macro Definition: " +
		c.GetToken(parser.TMLLexerID, 0).GetText())
	t.nesting_lvl++
}

func (t *TMLTreePrinterListener) EnterCommand(c *parser.CommandContext) {
	print(nestingLvlString(t.nesting_lvl) + "Command: (")
}

func (t *TMLTreePrinterListener) EnterStateLabel(c *parser.StateLabelContext) {
	print(c.GetText() + " ")
}

func (t *TMLTreePrinterListener) EnterTapeSymbol(c *parser.TapeSymbolContext) {
	print(c.GetText() + " ")
}

func (t *TMLTreePrinterListener) EnterDirection(c *parser.DirectionContext) {
	print(c.GetText() + " ")
}

func (t *TMLTreePrinterListener) ExitProgram(c *parser.ProgramContext) {
	t.nesting_lvl--
}

func (t *TMLTreePrinterListener) ExitMacroApp(c *parser.MacroAppContext) {
	println(")")
}

func (t *TMLTreePrinterListener) ExitMacroDef(c *parser.MacroDefContext) {
	t.nesting_lvl--
}

func (t *TMLTreePrinterListener) ExitCommand(c *parser.CommandContext) {
	println(")")
}
