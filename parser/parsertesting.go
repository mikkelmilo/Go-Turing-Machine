package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mikkelmilo/Go-Turing-Machine/parser/TM-Language"
	"io/ioutil"
)

func main() {
	// Setup the input

	b, err := ioutil.ReadFile("tests/testmacro.txt") // just pass the file name
	if err != nil {
		panic(err)
	}
	str := string(b) // convert content to a 'string'
	is := antlr.NewInputStream(str)
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
	var unfolder parser.TMLMacroUnfolder
	pt.Walk(&unfolder, tree)
	println("Program after unfolding:")
	for _, c := range unfolder.Program {
		println("(" + c.CurrentState + "," + c.NewState + "," + c.CurrentSymbol + "," + c.NewSymbol + "," + c.Direction + ")")
	}
	println("-----------------------")
	for name, m := range unfolder.Macros {
		println("macro " + name + ":")
		for _, c := range m {
			println("(" + c.CurrentState + "," + c.NewState + "," + c.CurrentSymbol + "," + c.NewSymbol + "," + c.Direction + ")")
		}
	}
}
