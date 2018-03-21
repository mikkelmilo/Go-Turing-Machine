package parser_test

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mikkelmilo/Go-Turing-Machine/parser/TM-Language"
	"testing"
)

func setupParser(program_string string) (*parser.TMLParser, parser.IStartContext) {
	is := antlr.NewInputStream(program_string)
	// Create the Lexer
	lexer := parser.NewTMLLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// Create the Parser
	p := parser.NewTMLParser(stream)
	tree := p.Start()
	return p, tree
}

func TestNewTMLBaseSemanticCheckerOnOnlyStartTransitionProgram(t *testing.T) {
	program_string := "(hs,a,_,1,>)"
	_, tree := setupParser(program_string)
	pt := antlr.ParseTreeWalkerDefault

	semant := parser.NewTMLBaseSemanticChecker()
	semant.Check(pt, tree)
}
