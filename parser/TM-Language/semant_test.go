package parser_test

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mikkelmilo/Go-Turing-Machine/parser/TM-Language"
	"github.com/stretchr/testify/assert"
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
	errors := semant.Check(pt, tree)
	assert.Len(t, errors, 2, "Expected two errors but got something else (1 error for missing accept state, and one for missing reject state")
}

func TestNewTMLBaseSemanticCheckerOnMultipleStartStates(t *testing.T) {
	program_string := "(hs,a,_,1,>)(hs,a,_,1,>)" +
		"(ha,a,_,1,>)(hr,a,_,1,>)" //all special states are necessary so we add these to prevent errors unrelated to this test
	_, tree := setupParser(program_string)
	pt := antlr.ParseTreeWalkerDefault

	semant := parser.NewTMLBaseSemanticChecker()
	errors := semant.Check(pt, tree)
	assert.Equal(t, 1, len(errors), "Expected one error.")
}

func TestNewTMLBaseSemanticCheckerOnMultipleAcceptStates(t *testing.T) {
	program_string := "(ha,a,_,1,>)(ha,a,_,1,>)" +
		"(hs,a,_,1,>)(hr,a,_,1,>)" //all special states are necessary so we add these to prevent errors unrelated to this test
	_, tree := setupParser(program_string)
	pt := antlr.ParseTreeWalkerDefault

	semant := parser.NewTMLBaseSemanticChecker()
	errors := semant.Check(pt, tree)
	assert.Equal(t, 1, len(errors), "Expected one error.")
}

func TestNewTMLBaseSemanticCheckerOnMultipleAcceptStatesInMacro(t *testing.T) {
	program_string :=
		"(hs,a,_,1,>)(hr,a,_,1,>)(ha,a,_,1,>)" +
			"define macro m {" +
			"(ha,a,_,1,>)" +
			"(ha,a,_,1,>)" +
			"(hs,a,_,1,>)" +
			"(hr,a,_,1,>)" +
			"}"
	_, tree := setupParser(program_string)
	pt := antlr.ParseTreeWalkerDefault

	semant := parser.NewTMLBaseSemanticChecker()
	errors := semant.Check(pt, tree)
	assert.Equal(t, 1, len(errors), "Expected one error.")
}

func TestNewTMLBaseSemanticCheckerOnMissingMacroStates(t *testing.T) {
	program_string :=
		"(ha,a,_,1,>)" +
			"define macro m {" +
			"(ha,a,_,1,>)" +
			"(hr,a,_,1,>)" +
			"}" +
			"(hs,a,_,1,>)" +
			"(hr,a,_,1,>)"
	_, tree := setupParser(program_string)
	pt := antlr.ParseTreeWalkerDefault

	semant := parser.NewTMLBaseSemanticChecker()
	errors := semant.Check(pt, tree)
	assert.Equal(t, 1, len(errors), "Expected one error.")
}
