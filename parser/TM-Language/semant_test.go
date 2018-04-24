package parser_test

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mikkelmilo/Go-Turing-Machine/parser/TM-Language"
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
	tree := p.Start()
	return p, tree
}

func TestNewTMLBaseSemanticCheckerOnOnlyStartTransitionProgram(t *testing.T) {
	program_string := "(hs,a,_,1,>)"
	_, tree := setupParser(program_string)
	pt := antlr.ParseTreeWalkerDefault

	errors := parser.CheckSemantic(pt, tree)
	assert.Len(t, errors, 2, "Expected two errors but got something else (1 error for missing accept state, and one for missing reject state")
}

func TestNewTMLBaseSemanticCheckerOnMultipleStartStates(t *testing.T) {
	program_string := "(hs,a,_,1,>)(hs,a,_,1,>)" +
		"(ha,a,_,1,>)(hr,a,_,1,>)" //all special states are necessary so we add these to prevent errors unrelated to this test
	_, tree := setupParser(program_string)
	pt := antlr.ParseTreeWalkerDefault

	errors := parser.CheckSemantic(pt, tree)
	assert.Equal(t, 1, len(errors), "Expected one error.")
}

func TestNewTMLBaseSemanticCheckerOnMultipleAcceptStates(t *testing.T) {
	program_string := "(ha,a,_,1,>)(ha,a,_,1,>)" +
		"(hs,a,_,1,>)(hr,a,_,1,>)" //all special states are necessary so we add these to prevent errors unrelated to this test
	_, tree := setupParser(program_string)
	pt := antlr.ParseTreeWalkerDefault
	errors := parser.CheckSemantic(pt, tree)
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
	errors := parser.CheckSemantic(pt, tree)
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

	errors := parser.CheckSemantic(pt, tree)
	assert.Equal(t, 1, len(errors), "Expected one error.")
}

/*
func TestTMLBaseSemanticChecker_CheckSequentialProgram(t *testing.T) {
	program := []parser.Command{
		parser.Command{CurrentState: "hs", NewState: "a"},
		parser.Command{CurrentState: "a", NewState: "b"},
		parser.Command{CurrentState: "a", NewState: "c"},
		parser.Command{CurrentState: "c", NewState: "d"},
		parser.Command{CurrentState: "e", NewState: "a"}, // state e should be unreachable since it does not occur in 'NewState' in any command
		parser.Command{CurrentState: "e", NewState: "f"}, // state f should also be unreachable since it is only reachable from a, who is not reachable.
		parser.Command{CurrentState: "d", NewState: "ha"},
		parser.Command{CurrentState: "b", NewState: "hr"},
	}
	semant := parser.NewTMLBaseSemanticChecker()
	err, unreachable_indices := semant.CheckSequentialProgram(program)
	assert.Nil(t, err, "expected no errors but got one")
	assert.Equal(t, []int{4, 5}, unreachable_indices, "expected command indices 4 and 5 to be unreachable")
}

func TestTMLBaseSemanticChecker_CheckSequentialProgram_unreachable_hr(t *testing.T) {
	program := []parser.Command{
		parser.Command{CurrentState: "hs", NewState: "a"},
		parser.Command{CurrentState: "a", NewState: "ha"},
		parser.Command{CurrentState: "e", NewState: "hr"}, // state 'e' is not reachable, therefore hr is neither
	}
	semant := parser.NewTMLBaseSemanticChecker()
	err, unreachable_indices := semant.CheckSequentialProgram(program)
	assert.NotNil(t, err, "expected an error but got none")
	assert.Nil(t, unreachable_indices)
}

func TestTMLBaseSemanticChecker_CheckSequentialProgram_unreachable_ha(t *testing.T) {
	program := []parser.Command{
		parser.Command{CurrentState: "hs", NewState: "a"},
		parser.Command{CurrentState: "a", NewState: "hr"},
		parser.Command{CurrentState: "e", NewState: "ha"}, // state 'e' is not reachable, therefore ha is neither
	}
	semant := parser.NewTMLBaseSemanticChecker()
	err, unreachable_indices := semant.CheckSequentialProgram(program)
	assert.NotNil(t, err, "expected an error but got none")
	assert.Nil(t, unreachable_indices)
}
*/

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

	unreachables := parser.FindUnreachableNodes(*g, nodes)
	assert.Len(t, unreachables, 2)
}
