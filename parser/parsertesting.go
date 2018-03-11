package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mikkelmilo/Go-Turing-Machine/parser/TM-Language"
)

func main() {
	// Setup the input
	is := antlr.NewInputStream("" +
		"/* test comment */\n" +
		"// another comment\n" +
		"(hs,a,_,1,>)\n" +
		"define macro testMacro {\n" +
		"(a,b,2,1,<)\n" +
		"}\n" +
		"(b, 0)testMacro(b,a)\n" +
		"(a,b,_,0,>)\n" +
		"(b,ha,_,1,_)\n")

	// Create the Lexer
	lexer := parser.NewTMLLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewTMLParser(stream)
	p.RemoveErrorListeners()
	var errListener parser.TMLerrorListener
	p.AddErrorListener(&errListener)
	tree := p.Start()
	var treePrinter parser.TMLTreePrinterListener
	// Finally parse the expression
	pt := antlr.ParseTreeWalkerDefault
	pt.Walk(&treePrinter, tree)
	if len(errListener.Errors) != 0 {
		for _, err := range errListener.Errors {
			println(err.Error())
		}
	}
}
