package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

/*
 * This file contains a Tree walker which checks the semantics of the program.
 * This includes:
 *    - the program contains hs, ha, and hr states
 *    - all states must be reachable from hs (the same is true for macros and their states)
 *    - macros must also contain hs, ha, and hr states
 *    - warnings will be produced (if enabled) if there are unbreakable cycles, ie.
 *      cycles which has no sequence of transitions to either ha or hr
 */

type TMLSemanticChecker interface {
	Check(pt *antlr.ParseTreeWalker, tree IStartContext) []TMLError
}

func NewTMLBaseSemanticChecker() TMLBaseSemanticChecker {
	return TMLBaseSemanticChecker{}
}

type TMLBaseSemanticChecker struct {
}

func (semant *TMLBaseSemanticChecker) Check(pt *antlr.ParseTreeWalker, tree IStartContext) []TMLError {
	walker := semantTreeListener{}
	pt.Walk(&walker, tree)
	return walker.errors
}

type semantTreeListener struct {
	errors             []TMLError
	inMacro            bool
	seenStartState     bool
	seenAcceptState    bool
	seenRejectState    bool
	startStateChanged  bool
	acceptStateChanged bool
	rejectStateChanged bool
}

func (semant *semantTreeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
}

func (semant *semantTreeListener) VisitTerminal(node antlr.TerminalNode) {
}

func (semant *semantTreeListener) VisitErrorNode(node antlr.ErrorNode) {
}

func (semant *semantTreeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
}

func (semant *semantTreeListener) EnterStart(c *StartContext) {
}

func (semant *semantTreeListener) EnterProgram(c *ProgramContext) {
	semant.errors = []TMLError{}
	semant.inMacro = false
	// predicates to determine if certain states have been seen so far in the current scope
	semant.seenAcceptState = false
	semant.seenStartState = false
	semant.seenRejectState = false
	semant.acceptStateChanged = false
	semant.startStateChanged = false
	semant.rejectStateChanged = false
}

func (semant *semantTreeListener) EnterStatement(c *StatementContext) {
}

func (semant *semantTreeListener) EnterMacroApp(c *MacroAppContext) {
}

func (semant *semantTreeListener) EnterMacroDef(c *MacroDefContext) {
	semant.seenRejectState = false
	semant.seenAcceptState = false
	semant.seenStartState = false
	semant.inMacro = true

}

func (semant *semantTreeListener) EnterCommand(c *CommandContext) {
	switch c.GetCurrentState().GetText() {
	case "hs":
		if !semant.seenStartState {
			semant.seenStartState = true
			if !semant.inMacro {
				semant.startStateChanged = true
			}
		} else {
			// else if the current state symbol is the start state, but not the first one
			// seen in this scope, add an error.
			semant.AppendErrorMsg("Multiple start states defined", c.GetStart())
		}
	case "ha":
		if !semant.seenAcceptState {
			semant.seenAcceptState = true
			if !semant.inMacro {
				semant.acceptStateChanged = true
			}
		} else {
			semant.AppendErrorMsg("Multiple accept states defined", c.GetStart())
		}
	case "hr":
		if !semant.seenRejectState {
			semant.seenRejectState = true
			if !semant.inMacro {
				semant.rejectStateChanged = true
			}
		} else {
			semant.AppendErrorMsg("Multiple reject states defined", c.GetStart())
		}
	}

}

func (semant *semantTreeListener) EnterStateLabel(c *StateLabelContext) {
}

func (semant *semantTreeListener) EnterTapeSymbol(c *TapeSymbolContext) {
}

func (semant *semantTreeListener) EnterDirection(c *DirectionContext) {
}

func (semant *semantTreeListener) ExitStart(c *StartContext) {
}

func (semant *semantTreeListener) ExitProgram(c *ProgramContext) {
	// check if the main TM contained a start, accept, and reject state. If not, report an error for each missing state.
	if !semant.seenStartState {
		semant.AppendErrorMsg("Missing start state", c.GetStart())
	}
	if !semant.seenAcceptState {
		semant.AppendErrorMsg("Missing accept state", c.GetStart())
	}
	if !semant.seenRejectState {
		semant.AppendErrorMsg("Missing reject state", c.GetStart())
	}
}

func (semant *semantTreeListener) ExitStatement(c *StatementContext) {
}

func (semant *semantTreeListener) ExitMacroApp(c *MacroAppContext) {
}

func (semant *semantTreeListener) ExitMacroDef(c *MacroDefContext) {
	semant.inMacro = false
	macroName := c.GetToken(TMLParserID, 0).GetText()
	// check if macro contained a start, accept, and reject state. If not, report an error for each missing state.
	if !semant.seenStartState {
		semant.AppendErrorMsg("Missing start state in macro: "+macroName, c.GetStart())
	}
	if !semant.seenAcceptState {
		semant.AppendErrorMsg("Missing accept state in macro: "+macroName, c.GetStart())
	}
	if !semant.seenRejectState {
		semant.AppendErrorMsg("Missing reject state in macro: "+macroName, c.GetStart())
	}
	semant.seenStartState = semant.startStateChanged
	semant.seenAcceptState = semant.acceptStateChanged
	semant.seenRejectState = semant.rejectStateChanged
}

func (semant *semantTreeListener) ExitCommand(c *CommandContext) {
}

func (semant *semantTreeListener) ExitStateLabel(c *StateLabelContext) {
}

func (semant *semantTreeListener) ExitTapeSymbol(c *TapeSymbolContext) {
}

func (semant *semantTreeListener) ExitDirection(c *DirectionContext) {

}

func (semant *semantTreeListener) AppendErrorMsg(msg string, c antlr.Token) {
	semant.errors = append(semant.errors,
		TMLError{
			column: c.GetColumn(),
			line:   c.GetLine(),
			msg:    msg,
		})
}
