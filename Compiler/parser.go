package Compiler

import (
	"bytes"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mikkelmilo/Go-Turing-Machine/Compiler/antlr-parser"
)

type TMLParseTree = parser.IStartContext

type TMLParser func(bytes.Buffer) ([]TMLError, TMLParseTree)

func ParseTMLProgram(program bytes.Buffer) ([]TMLError, TMLParseTree) {
	programStr := program.String()
	// first lex and parse the program, and report any errors

	is := antlr.NewInputStream(programStr)
	// Create the Lexer
	lexer := parser.NewTMLLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// Create the Parser
	p := parser.NewTMLParser(stream)
	// add a listener for parse errors and walk the tree
	var parserErrorsListener TMLParserErrorsListener
	p.AddErrorListener(&parserErrorsListener)
	tree := p.Start()

	if errs := parserErrorsListener.Errors; len(errs) != 0 {
		return errs, nil
	}

	return nil, tree
}

type TMLParserErrorsListener struct {
	Errors []TMLError
}

func (el *TMLParserErrorsListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{},
	line, column int, msg string, e antlr.RecognitionException) {
	el.Errors = append(el.Errors, TMLError{line: line, column: column, msg: msg})
}

// we assert this never happens because this is inherently a property of the grammar, not the specific instance.
func (el *TMLParserErrorsListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	panic("implement me")
}

// we assert this never happens because this is inherently a property of the grammar, not the specific instance.
func (el *TMLParserErrorsListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	panic("implement me")
}

// we assert this never happens because this is inherently a property of the grammar, not the specific instance.
func (el *TMLParserErrorsListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	panic("implement me")
}
