package Compiler

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mikkelmilo/Go-Turing-Machine/Compiler/antlr-parser"
	"github.com/stretchr/testify/assert"
	"github.com/twmb/algoimpl/go/graph"
	"testing"
)

func setupParser(program_string string) (*parser.TMLParser, parser.IStartContext) {
	is := antlr.NewInputStream(program_string)
	// Create the Lexer
	lexer := parser.NewTMLLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// Create the Parser
	p := parser.NewTMLParser(stream)
	p.RemoveErrorListeners()
	tree := p.Start()

	return p, tree
}

func TestNewTMLBaseSemanticCheckerOnOnlyStartTransitionProgram(t *testing.T) {
	program_string := "(hs,a,_,1,>)"
	_, tree := setupParser(program_string)
	pt := antlr.ParseTreeWalkerDefault

	errors := CheckSemantic(pt, tree)
	assert.Len(t, errors, 2, "Expected two errors but got something else (1 error for missing accept state, and one for missing reject state")
}

func TestNewTMLBaseSemanticCheckerOnMultipleStartStates(t *testing.T) {
	program_string := "(hs,ha,_,1,>)(hs,hr,_,1,>)"
	_, tree := setupParser(program_string)
	pt := antlr.ParseTreeWalkerDefault

	errors := CheckSemantic(pt, tree)
	fmt.Println(errors)
	assert.Equal(t, 1, len(errors), "Expected one error.")
}

func TestNewTMLBaseSemanticCheckerOnMultipleStartStatesInMacro(t *testing.T) {
	program_string :=
		"(hs,hr,_,1,>)(a,ha,_,1,>)" +
			"define macro m {" +
			"(hs,hr,_,1,>)" +
			"(hs,ha,_,1,>)" +
			"}"
	_, tree := setupParser(program_string)
	pt := antlr.ParseTreeWalkerDefault
	errors := CheckSemantic(pt, tree)
	fmt.Println(errors)

	assert.Equal(t, 1, len(errors), "Expected one error.")
}

func TestNewTMLBaseSemanticCheckerOnIllegalTransitionFromAcceptStateInMacro(t *testing.T) {
	program_string :=
		"(hs,hr,_,1,>)(a,ha,_,1,>)" +
			"define macro m {" +
			"(ha,a,_,1,>)" +
			"(a,ha,_,1,>)" +
			"(a,hr,_,1,>)" +
			"(hs,a,_,1,>)" +
			"}"
	_, tree := setupParser(program_string)
	pt := antlr.ParseTreeWalkerDefault
	errors := CheckSemantic(pt, tree)
	fmt.Println(errors)

	assert.Equal(t, 1, len(errors), "Expected one error.")
}

func TestNewTMLBaseSemanticCheckerOnMissingMacroStates(t *testing.T) {
	program_string :=
		"(hs,a,_,1,>)" +
			"(a,hr,_,1,>)" +
			"(a,ha,_,1,>)" +
			"define macro m {" +
			"(a,ha,_,1,>)" +
			"(a,hr,_,1,>)" +
			"}"
	_, tree := setupParser(program_string)
	pt := antlr.ParseTreeWalkerDefault
	errors := CheckSemantic(pt, tree)
	fmt.Printf("%v\n", errors)
	assert.Equal(t, 1, len(errors), "Expected one error.")
}

func TestCheckSemantic_unreachable_states(t *testing.T) {
	program_string :=
		"(hs,a,_,1,>)" +
			"(a,b,_,1,>)" +
			"(c,a,_,1,>)" + // c and hr are unreachable
			"(c,hr,_,1,>)" +
			"(a,ha,_,1,>)"
	_, tree := setupParser(program_string)
	pt := antlr.ParseTreeWalkerDefault

	errors := CheckSemantic(pt, tree)
	fmt.Printf("%v\n", errors)
	assert.Equal(t, 2, len(errors), "Expected one error.")
}

func TestCheckSemantic_unreachable_macrostates1(t *testing.T) {
	program_string :=
		"(ha,a,_,1,>)" +
			"define macro m {" +
			"(hs,a,_,1,>)" +
			"(a,ha,_,1,>)" +
			"(c,a,_,1,>)" + // c and hr are unreachable
			"(c,hr,_,1,>)" +
			"}" +
			"(hs,a,_,1,>)" +
			"(hr,a,_,1,>)"

	_, tree := setupParser(program_string)
	pt := antlr.ParseTreeWalkerDefault

	errors := CheckSemantic(pt, tree)
	fmt.Printf("%v\n", errors)
	assert.Equal(t, 2, len(errors), "Expected one error.")
}

func TestCheckSemantic_unreachable_macrostates2(t *testing.T) {
	program_string :=
		"(hs,a,_,1,>)" +
			"(a,ha,_,1,>)" +
			"(a,hr,_,1,>)" +
			"define macro m {" +
			"(hs,a,_,1,>)" +
			"(a,ha,_,1,>)" +
			"(c,a,_,1,>)" + // c is unreachable
			"(c,hr,_,1,>)" +
			"(a,hr,_,1,>)" + // but hr is not since a can reach it, and a is reachable
			"}"
	_, tree := setupParser(program_string)
	pt := antlr.ParseTreeWalkerDefault

	errors := CheckSemantic(pt, tree)
	fmt.Printf("%v\n", errors)
	assert.Equal(t, 1, len(errors), "Expected one error.")
}

func TestFindUnreachableNodes(t *testing.T) {
	g := graph.New(graph.Directed)
	nodes := make(map[string]graph.Node, 0)
	nodes["hs"] = g.MakeNode()
	*nodes["hs"].Value = "hs"
	nodes["hr"] = g.MakeNode()
	*nodes["hr"].Value = "hr"
	nodes["ha"] = g.MakeNode()
	*nodes["ha"].Value = "ha"
	nodes["a"] = g.MakeNode()
	*nodes["a"].Value = "a"
	nodes["b"] = g.MakeNode()
	*nodes["b"].Value = "b"
	nodes["c"] = g.MakeNode()
	*nodes["c"].Value = "c"

	g.MakeEdge(nodes["hs"], nodes["a"])
	g.MakeEdge(nodes["a"], nodes["hr"])
	g.MakeEdge(nodes["a"], nodes["ha"])
	g.MakeEdge(nodes["b"], nodes["c"]) // b and c are unreachable
	g.MakeEdge(nodes["c"], nodes["a"])

	unreachables := FindUnreachableNodes(*g, nodes)
	assert.Len(t, unreachables, 2)
}

func TestBuildMainProgramGraph(t *testing.T) {
	program_string := "" +
		"(hs,a,_,1,>)" +
		"define macro testMacro {" +
		"}" +
		"(b, 0)testMacro(c,a)" +
		"(c, 0)testMacro(a,b)" +
		"(a,b,_,0,>) " +
		"(b,ha,_,1,_)" +
		"(c,hr,_,1,_)"

	_, tree := setupParser(program_string)
	pt := antlr.ParseTreeWalkerDefault
	var mainprogbuilder MainProgramGraphBuilder
	pt.Walk(&mainprogbuilder, tree)
	assert.Len(t, mainprogbuilder.Nodes, 6)

}
