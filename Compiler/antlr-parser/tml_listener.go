// Generated from Compiler/antlr-parser//TML.g4 by ANTLR 4.7.

package parser // TML

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TMLListener is a complete listener for a parse tree produced by TMLParser.
type TMLListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterMacroApp is called when entering the macroApp production.
	EnterMacroApp(c *MacroAppContext)

	// EnterMacroDef is called when entering the macroDef production.
	EnterMacroDef(c *MacroDefContext)

	// EnterCommand is called when entering the command production.
	EnterCommand(c *CommandContext)

	// EnterStateLabel is called when entering the stateLabel production.
	EnterStateLabel(c *StateLabelContext)

	// EnterTapeSymbol is called when entering the tapeSymbol production.
	EnterTapeSymbol(c *TapeSymbolContext)

	// EnterDirection is called when entering the direction production.
	EnterDirection(c *DirectionContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitMacroApp is called when exiting the macroApp production.
	ExitMacroApp(c *MacroAppContext)

	// ExitMacroDef is called when exiting the macroDef production.
	ExitMacroDef(c *MacroDefContext)

	// ExitCommand is called when exiting the command production.
	ExitCommand(c *CommandContext)

	// ExitStateLabel is called when exiting the stateLabel production.
	ExitStateLabel(c *StateLabelContext)

	// ExitTapeSymbol is called when exiting the tapeSymbol production.
	ExitTapeSymbol(c *TapeSymbolContext)

	// ExitDirection is called when exiting the direction production.
	ExitDirection(c *DirectionContext)
}
