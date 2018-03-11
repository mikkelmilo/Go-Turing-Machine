package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mikkelmilo/Go-Turing-Machine/parser/TM-Language"
)

type tmlListener struct {
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

/*
func (t *tmlListener) VisitTerminal(node antlr.TerminalNode) {
	panic("implement me")
}
*/

func (t *tmlListener) VisitErrorNode(node antlr.ErrorNode) {
	// TODO ErrorNodes do not contain position. Should change implementation to store something else.
	t.errors = append(t.errors, node)
}

/*
func (t *tmlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	panic("implement me")
}

func (t *tmlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	panic("implement me")
}
*/
func (t *tmlListener) EnterStart(c *parser.StartContext) {
	t.nesting_lvl = 0
}

func (t *tmlListener) EnterProgram(c *parser.ProgramContext) {
	println("Program:")
	t.nesting_lvl++

}

func (t *tmlListener) EnterStatement(c *parser.StatementContext) {
}

func (t *tmlListener) EnterMacroApp(c *parser.MacroAppContext) {
	print(nestingLvlString(t.nesting_lvl) + "Macro Application: (")
}

func (t *tmlListener) EnterMacroDef(c *parser.MacroDefContext) {
	println(nestingLvlString(t.nesting_lvl) + "Macro Definition: " +
		c.GetToken(parser.TMLLexerID, 0).GetText())
	t.nesting_lvl++
}

func (t *tmlListener) EnterCommand(c *parser.CommandContext) {
	print(nestingLvlString(t.nesting_lvl) + "Command: (")
}

func (t *tmlListener) EnterStateLabel(c *parser.StateLabelContext) {
	print(c.GetText() + " ")
}

func (t *tmlListener) EnterTapeSymbol(c *parser.TapeSymbolContext) {
	print(c.GetText() + " ")
}

func (t *tmlListener) EnterDirection(c *parser.DirectionContext) {
	print(c.GetText() + " ")
}

func (t *tmlListener) ExitStart(c *parser.StartContext) {
}

func (t *tmlListener) ExitProgram(c *parser.ProgramContext) {
	t.nesting_lvl--
}

func (t *tmlListener) ExitStatement(c *parser.StatementContext) {
}

func (t *tmlListener) ExitMacroApp(c *parser.MacroAppContext) {
	println(")")
}

func (t *tmlListener) ExitMacroDef(c *parser.MacroDefContext) {
	t.nesting_lvl--
}

func (t *tmlListener) ExitCommand(c *parser.CommandContext) {
	println(")")
}

func (t *tmlListener) ExitStateLabel(c *parser.StateLabelContext) {
}

func (t *tmlListener) ExitTapeSymbol(c *parser.TapeSymbolContext) {
}

func (t *tmlListener) ExitDirection(c *parser.DirectionContext) {
}

func main() {
	// Setup the input
	is := antlr.NewInputStream("" +
		"/* test comment */" +
		"// another comment" +
		"(hs,a,_,1,>)\n" +
		"define macro testMacro {\n" +
		"(a,b,2,1,<\n" +
		"}\n" +
		"(b, 0)testMacro(b,a)\n" +
		"(a,b,_,0,>)\n" +
		"(b,ha,_,1,_)\n")

	// Create the Lexer
	lexer := parser.NewTMLLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewTMLParser(stream)
	tree := p.Start()
	var tmlListener tmlListener
	// Finally parse the expression
	pt := antlr.ParseTreeWalkerDefault
	pt.Walk(&tmlListener, tree)
	if len(tmlListener.errors) != 0 {
		for _, node := range tmlListener.errors {
			println(node.GetText())
		}
	}
}
