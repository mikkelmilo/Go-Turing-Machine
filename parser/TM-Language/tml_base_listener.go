// Generated from TM-Language/TML.g4 by ANTLR 4.7.

package parser // TML

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTMLListener is a complete listener for a parse tree produced by TMLParser.
type BaseTMLListener struct{}

var _ TMLListener = &BaseTMLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTMLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTMLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTMLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTMLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseTMLListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseTMLListener) ExitStart(ctx *StartContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseTMLListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseTMLListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseTMLListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseTMLListener) ExitStatement(ctx *StatementContext) {}

// EnterMacroApp is called when production macroApp is entered.
func (s *BaseTMLListener) EnterMacroApp(ctx *MacroAppContext) {}

// ExitMacroApp is called when production macroApp is exited.
func (s *BaseTMLListener) ExitMacroApp(ctx *MacroAppContext) {}

// EnterMacroDef is called when production macroDef is entered.
func (s *BaseTMLListener) EnterMacroDef(ctx *MacroDefContext) {}

// ExitMacroDef is called when production macroDef is exited.
func (s *BaseTMLListener) ExitMacroDef(ctx *MacroDefContext) {}

// EnterCommand is called when production command is entered.
func (s *BaseTMLListener) EnterCommand(ctx *CommandContext) {}

// ExitCommand is called when production command is exited.
func (s *BaseTMLListener) ExitCommand(ctx *CommandContext) {}

// EnterStateLabel is called when production stateLabel is entered.
func (s *BaseTMLListener) EnterStateLabel(ctx *StateLabelContext) {}

// ExitStateLabel is called when production stateLabel is exited.
func (s *BaseTMLListener) ExitStateLabel(ctx *StateLabelContext) {}

// EnterTapeSymbol is called when production tapeSymbol is entered.
func (s *BaseTMLListener) EnterTapeSymbol(ctx *TapeSymbolContext) {}

// ExitTapeSymbol is called when production tapeSymbol is exited.
func (s *BaseTMLListener) ExitTapeSymbol(ctx *TapeSymbolContext) {}

// EnterDirection is called when production direction is entered.
func (s *BaseTMLListener) EnterDirection(ctx *DirectionContext) {}

// ExitDirection is called when production direction is exited.
func (s *BaseTMLListener) ExitDirection(ctx *DirectionContext) {}
