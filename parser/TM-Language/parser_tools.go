package parser

/*
	This file contains tools for working with the generated parser.
	This includes listeners for syntax Errors, and pretty printing of a TML AST.
*/

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strconv"
	"strings"
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
	name := c.GetToken(TMLLexerID, 0).GetText()
	print(nestingLvlString(t.nesting_lvl) + "Macro Application: " + name + "(")
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

type Command struct {
	CurrentState  string
	NewState      string
	CurrentSymbol string
	NewSymbol     string
	Direction     string
}

// A TML Tree listener which constructs a sequential program from a given TML Tree by unfolding all macro definitions
type TMLMacroUnfolder struct {
	Program      []Command
	macros       map[string][]Command //maps from the macro name to its list of tuples
	currentMacro string
	uniqueNr     int
}

func (*TMLMacroUnfolder) VisitTerminal(node antlr.TerminalNode) {
}

func (*TMLMacroUnfolder) VisitErrorNode(node antlr.ErrorNode) {
}

func (*TMLMacroUnfolder) EnterEveryRule(ctx antlr.ParserRuleContext) {
}

func (*TMLMacroUnfolder) ExitEveryRule(ctx antlr.ParserRuleContext) {
}

func (*TMLMacroUnfolder) EnterStart(c *StartContext) {
}

func (t *TMLMacroUnfolder) EnterProgram(c *ProgramContext) {
	t.Program = make([]Command, 0)
	t.macros = make(map[string][]Command)
	t.currentMacro = ""
	t.uniqueNr = 0
}

func (*TMLMacroUnfolder) EnterStatement(c *StatementContext) {
}

func (t *TMLMacroUnfolder) EnterMacroApp(c *MacroAppContext) {
	macroName := c.GetToken(TMLLexerID, 0).GetText()
	text := c.GetText()[1 : len(c.GetText())-1] //get text and remove surrounding ( and )
	text = strings.Replace(text, ")", ",", -1)
	text = strings.Replace(text, "(", ",", -1)
	//elems is a list of the different "elements" of this macro application
	elems := strings.Split(text, ",")

	curStateName := elems[0]
	curSymbol := elems[1]
	acceptState := elems[3]
	rejectState := elems[4]

	// make transition rules from curStatName to start starte of macro, and transitions when macro halts in accept or reject

	macro_hs_trans := Command{
		CurrentState:  curStateName,
		NewState:      "hs_" + macroName + strconv.Itoa(t.uniqueNr),
		CurrentSymbol: curSymbol,
		NewSymbol:     curSymbol,
		Direction:     "_",
	}
	macro_ha_trans := Command{
		CurrentState:  "ha_" + macroName + strconv.Itoa(t.uniqueNr),
		NewState:      acceptState,
		CurrentSymbol: "_",
		NewSymbol:     "_",
		Direction:     "_",
	}
	macro_hr_trans := Command{
		CurrentState:  "hr_" + macroName + strconv.Itoa(t.uniqueNr),
		NewState:      rejectState,
		CurrentSymbol: "_",
		NewSymbol:     "_",
		Direction:     "_",
	}

	// append these commands to the program
	t.Program = append(t.Program, macro_hs_trans, macro_ha_trans, macro_hr_trans)
	// generate the macro commands with unique state names
	macroCommands := t.GenerateUniqueStates(t.macros[macroName], macroName, t.uniqueNr)
	t.uniqueNr++
	// append the macro to the program
	t.Program = append(t.Program, macroCommands...)
}

func (t *TMLMacroUnfolder) EnterMacroDef(c *MacroDefContext) {
	t.currentMacro = c.GetToken(TMLLexerID, 0).GetText()
}

func (t *TMLMacroUnfolder) EnterCommand(c *CommandContext) {
	//therefore elems[0] will contain the current state string, elems[1] the new state string, elems[2] the current symbol, etc.
	//since we have already syntax checked the program, we can assume that this command is syntactically correct
	command := c.GetText()[1 : len(c.GetText())-1]
	elems := strings.Split(command, ",")
	cm := Command{
		CurrentState:  elems[0],
		NewState:      elems[1],
		CurrentSymbol: elems[2],
		NewSymbol:     elems[3],
		Direction:     elems[4],
	}
	// if we are not in a macro definition, add to t.Program list, else add to t.macros[t.currentMacro]
	if t.currentMacro == "" {
		t.Program = append(t.Program, cm)
	} else {
		t.macros[t.currentMacro] = append(t.macros[t.currentMacro], cm)
	}
}

func (t *TMLMacroUnfolder) EnterStateLabel(c *StateLabelContext) {
}

func (t *TMLMacroUnfolder) EnterTapeSymbol(c *TapeSymbolContext) {
}

func (t *TMLMacroUnfolder) EnterDirection(c *DirectionContext) {
	t.Program[len(t.Program)-1].Direction = c.GetText()
}

func (*TMLMacroUnfolder) ExitStart(c *StartContext) {
}

func (*TMLMacroUnfolder) ExitProgram(c *ProgramContext) {
}

func (*TMLMacroUnfolder) ExitStatement(c *StatementContext) {
}

func (*TMLMacroUnfolder) ExitMacroApp(c *MacroAppContext) {
}

func (t *TMLMacroUnfolder) ExitMacroDef(c *MacroDefContext) {
	t.currentMacro = ""
}

func (*TMLMacroUnfolder) ExitCommand(c *CommandContext) {
}

func (*TMLMacroUnfolder) ExitStateLabel(c *StateLabelContext) {
}

func (*TMLMacroUnfolder) ExitTapeSymbol(c *TapeSymbolContext) {
}

func (*TMLMacroUnfolder) ExitDirection(c *DirectionContext) {
}

func (t *TMLMacroUnfolder) GenerateUniqueStates(c []Command, macroName string, seed int) []Command {
	res := make([]Command, len(c))
	for i, command := range c {
		res[i] = command
		res[i].CurrentState = res[i].CurrentState + "_" + macroName + strconv.Itoa(seed)
		res[i].NewState = res[i].NewState + "_" + macroName + strconv.Itoa(seed)
	}

	return res
}
