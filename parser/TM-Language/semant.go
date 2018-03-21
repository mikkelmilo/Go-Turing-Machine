package parser

import "github.com/antlr/antlr4/runtime/Go/antlr"

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
	Check(pt *antlr.ParseTreeWalker, tree IStartContext) []error
}

func NewTMLBaseSemanticChecker() TMLBaseSemanticChecker {
	return TMLBaseSemanticChecker{}
}

type TMLBaseSemanticChecker struct {
}

func (semant *TMLBaseSemanticChecker) Check(pt *antlr.ParseTreeWalker, tree IStartContext) []error {
	pt.Walk(&semantTreeListener{}, tree)
	return []error{}
}

type semantTreeListener struct {
}

func (semant *semantTreeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	panic("implement me")
}

func (semant *semantTreeListener) VisitTerminal(node antlr.TerminalNode) {
	panic("implement me")
}

func (semant *semantTreeListener) VisitErrorNode(node antlr.ErrorNode) {
	panic("implement me")
}

func (semant *semantTreeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	panic("implement me")
}

func (semant *semantTreeListener) EnterStart(c *StartContext) {
	panic("implement me")
}

func (semant *semantTreeListener) EnterProgram(c *ProgramContext) {
	panic("implement me")
}

func (semant *semantTreeListener) EnterStatement(c *StatementContext) {
	panic("implement me")
}

func (semant *semantTreeListener) EnterMacroApp(c *MacroAppContext) {
	panic("implement me")
}

func (semant *semantTreeListener) EnterMacroDef(c *MacroDefContext) {
	panic("implement me")
}

func (semant *semantTreeListener) EnterCommand(c *CommandContext) {
	panic("implement me")
}

func (semant *semantTreeListener) EnterStateLabel(c *StateLabelContext) {
	panic("implement me")
}

func (semant *semantTreeListener) EnterTapeSymbol(c *TapeSymbolContext) {
	panic("implement me")
}

func (semant *semantTreeListener) EnterDirection(c *DirectionContext) {
	panic("implement me")
}

func (semant *semantTreeListener) ExitStart(c *StartContext) {
	panic("implement me")
}

func (semant *semantTreeListener) ExitProgram(c *ProgramContext) {
	panic("implement me")
}

func (semant *semantTreeListener) ExitStatement(c *StatementContext) {
	panic("implement me")
}

func (semant *semantTreeListener) ExitMacroApp(c *MacroAppContext) {
	panic("implement me")
}

func (semant *semantTreeListener) ExitMacroDef(c *MacroDefContext) {
	panic("implement me")
}

func (semant *semantTreeListener) ExitCommand(c *CommandContext) {
	panic("implement me")
}

func (semant *semantTreeListener) ExitStateLabel(c *StateLabelContext) {
	panic("implement me")
}

func (semant *semantTreeListener) ExitTapeSymbol(c *TapeSymbolContext) {
	panic("implement me")
}

func (semant *semantTreeListener) ExitDirection(c *DirectionContext) {
	panic("implement me")
}
