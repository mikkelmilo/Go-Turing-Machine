package Compiler

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mikkelmilo/Go-Turing-Machine/Compiler/antlr-parser"
	"github.com/twmb/algoimpl/go/graph"
)

/*
 * This file contains a Tree walker which checks the semantics of the program.
 * This includes:
 *    - the program contains hs, ha, and hr states
 *    - all states must be reachable from hs (the same is true for Macros and their states)
 *    - Macros must also contain hs, ha, and hr states
 *    - Nice-to-have: warnings will be produced (if enabled) if there are unbreakable cycles, ie.
 */

type TMLSemanticChecker func(*antlr.ParseTreeWalker, TMLParseTree) []TMLError

func CheckSemantic(pt *antlr.ParseTreeWalker, tree TMLParseTree) []TMLError {

	// first walk the concrete syntax tree and report any basic semantic errors such as missing start/accept/reject states
	var walker semantTreeListener
	pt.Walk(&walker, tree)

	if len(walker.errors) != 0 {
		return walker.errors
	}

	// if no errors were found in the tree walk above, check for reachability properties of hs, ha, and hr (in both main program and macros)
	var unfolder TMLMacroUnfolder
	pt.Walk(&unfolder, tree)

	// for each local macro, check reachability of nodes and report and error if there are any.
	for macroName, m := range unfolder.Macros {
		g := graph.New(graph.Directed)
		nodes := make(map[string]graph.Node, 0)

		// make nodes for each state in the command list
		for _, c := range m {
			nodes[c.CurrentState] = g.MakeNode()
			*nodes[c.CurrentState].Value = c.CurrentState
			nodes[c.NewState] = g.MakeNode()
			*nodes[c.NewState].Value = c.NewState
		}
		// add edges
		for _, c := range m {
			g.MakeEdge(nodes[c.CurrentState], nodes[c.NewState])
		}
		// find unreachable states
		unreachables := FindUnreachableNodes(*g, nodes)
		errs := make([]TMLError, len(unreachables))
		for i, n := range unreachables {
			errs[i] = TMLError{msg: "state " + (*n.Value).(string) + " is unreachable in macro: " + macroName}
		}
		if len(errs) != 0 {
			return errs
		}
	}

	// check the main program for reachability
	var mainprogBuilder MainProgramGraphBuilder
	pt.Walk(&mainprogBuilder, tree)

	unreachables := FindUnreachableNodes(*mainprogBuilder.Graph, mainprogBuilder.Nodes)
	errs := make([]TMLError, len(unreachables))
	for i, n := range unreachables {
		errs[i] = TMLError{msg: "state " + (*n.Value).(string) + " is unreachable in main program"}
	}
	if len(errs) != 0 {
		return errs
	}
	return nil
}

func FindUnreachableNodes(g graph.Graph, nodes map[string]graph.Node) []graph.Node {
	visited := make(map[graph.Node]bool)
	for _, n := range nodes {
		visited[n] = false
	}
	visited[nodes["hs"]] = true
	remaining := g.Neighbors(nodes["hs"])

	for len(remaining) != 0 {
		// take out last node in remaining list
		n := remaining[len(remaining)-1]
		visited[n] = true                        // add this node as visited
		remaining = remaining[:len(remaining)-1] // remove last element

		// for each of the neighbours of n, add them to remaining list
		// only if they haven't already been visited
		for _, neighbour := range g.Neighbors(n) {
			if !visited[neighbour] {
				remaining = append(remaining, neighbour)
			}
		}
	}
	unvisited := make([]graph.Node, 0)
	for n, wasVisited := range visited {
		if !wasVisited {
			unvisited = append(unvisited, n)
		}
	}
	return unvisited
}

// -------------------------------------------------------------------------

/*
	semantTreeListener is a TMLTreListener that does basic semantic checking on a given TML concrete syntax tree.
	It checks that the program contains exactly one instance the necessary states:
	start, accept, and reject states (and does the same for all macro definitions).
*/
type semantTreeListener struct {
	*parser.BaseTMLListener
	errors             []TMLError
	inMacro            bool
	seenStartState     bool
	seenAcceptState    bool
	seenRejectState    bool
	startStateChanged  bool
	acceptStateChanged bool
	rejectStateChanged bool
}

func (semant *semantTreeListener) EnterProgram(c *parser.ProgramContext) {
	semant.errors = []TMLError{}
	semant.inMacro = false
	// predicates to determine if certain states have been seen so far in the current scope
	semant.seenAcceptState = false
	semant.seenStartState = false
	semant.seenRejectState = false
	semant.acceptStateChanged = false
	semant.startStateChanged = false
	semant.rejectStateChanged = false
}

func (semant *semantTreeListener) EnterMacroDef(c *parser.MacroDefContext) {
	semant.seenRejectState = false
	semant.seenAcceptState = false
	semant.seenStartState = false
	semant.inMacro = true

}

func (semant *semantTreeListener) EnterCommand(c *parser.CommandContext) {
	checkCommand := func(stateName string, stateType string) {
		switch stateName {
		case "hs":
			if !semant.seenStartState {
				semant.seenStartState = true
				if !semant.inMacro {
					semant.startStateChanged = true
				}
			}
		case "ha":
			if !semant.seenAcceptState {
				semant.seenAcceptState = true
				if !semant.inMacro {
					semant.acceptStateChanged = true
				}
			} else if stateType == "currentstate" {
				semant.AppendErrorMsg("Accept state cannot have transitions to other states", c.GetStart())
			}
		case "hr":
			if !semant.seenRejectState {
				semant.seenRejectState = true
				if !semant.inMacro {
					semant.rejectStateChanged = true
				}
			} else if stateType == "currentstate" {
				semant.AppendErrorMsg("Reject state cannot have transitions to other states", c.GetStart())
			}
		}
	}

	checkCommand(c.GetCurrentState().GetText(), "currentstate")
	checkCommand(c.GetNewState().GetText(), "newstate")
}

func (semant *semantTreeListener) ExitProgram(c *parser.ProgramContext) {
	// check if the main TM contained a start, accept, and reject state. If not, report an error for each missing state.
	if !semant.seenStartState {
		semant.AppendErrorMsg("Missing start state", c.GetStart())
	}
	if !semant.seenAcceptState {
		semant.AppendErrorMsg("Missing accept state", c.GetStart())
	}
	if !semant.seenRejectState {
		semant.AppendErrorMsg("Missing reject state", c.GetStart())
	}
}

func (semant *semantTreeListener) ExitMacroDef(c *parser.MacroDefContext) {
	semant.inMacro = false
	macroName := c.GetToken(parser.TMLParserID, 0).GetText()
	// check if macro contained a start, accept, and reject state. If not, report an error for each missing state.
	if !semant.seenStartState {
		semant.AppendErrorMsg("Missing start state in macro: "+macroName, c.GetStart())
	}
	if !semant.seenAcceptState {
		semant.AppendErrorMsg("Missing accept state in macro: "+macroName, c.GetStart())
	}
	if !semant.seenRejectState {
		semant.AppendErrorMsg("Missing reject state in macro: "+macroName, c.GetStart())
	}
	semant.seenStartState = semant.startStateChanged
	semant.seenAcceptState = semant.acceptStateChanged
	semant.seenRejectState = semant.rejectStateChanged
}

func (semant *semantTreeListener) AppendErrorMsg(msg string, c antlr.Token) {
	semant.errors = append(semant.errors,
		TMLError{
			column: c.GetColumn(),
			line:   c.GetLine(),
			msg:    msg,
		})
}

// -------------------------------------------------------------------------------------

// this struct builds a graph of the main program, but ignores the content of the macros
// and instead just assumes all macros internally satisfy the reachability requirements of states
type MainProgramGraphBuilder struct {
	*parser.BaseTMLListener
	Graph   *graph.Graph
	Nodes   map[string]graph.Node
	inMacro bool
}

func (m *MainProgramGraphBuilder) EnterStart(c *parser.StartContext) {
	m.inMacro = false
	m.Graph = graph.New(graph.Directed)
	m.Nodes = make(map[string]graph.Node, 0)

}

func (m *MainProgramGraphBuilder) EnterMacroApp(c *parser.MacroAppContext) {
	// make transitions from entering state to accepting and rejecting state of macro

	// only add commands that are not inside macros
	if !m.inMacro {
		// add new nodes for current and new state, and make an edge between them.
		enteringstateName := c.GetEnteringState().GetText()
		// if entering state doesn't exist as a node in the graph, add it
		if _, ok := m.Nodes[enteringstateName]; !ok {
			n1 := m.Graph.MakeNode()
			*n1.Value = enteringstateName
			m.Nodes[enteringstateName] = n1

		}
		macroacceptName := c.GetAcceptState().GetText()
		// if macro accept state doesn't exist as a node in the graph, add it
		if _, ok := m.Nodes[macroacceptName]; !ok {
			n2 := m.Graph.MakeNode()
			*n2.Value = macroacceptName
			m.Nodes[macroacceptName] = n2
		}

		macrorejectName := c.GetRejectState().GetText()
		// if macro accept state doesn't exist as a node in the graph, add it
		if _, ok := m.Nodes[macrorejectName]; !ok {
			n2 := m.Graph.MakeNode()
			*n2.Value = macrorejectName
			m.Nodes[macrorejectName] = n2
		}

		m.Graph.MakeEdge(m.Nodes[enteringstateName], m.Nodes[macroacceptName])
		m.Graph.MakeEdge(m.Nodes[enteringstateName], m.Nodes[macrorejectName])
	}
}

func (m *MainProgramGraphBuilder) EnterMacroDef(c *parser.MacroDefContext) {
	m.inMacro = true
}

func (m *MainProgramGraphBuilder) EnterCommand(c *parser.CommandContext) {
	// only add commands that are not inside macros
	if !m.inMacro {
		// add new nodes for current and new state, and make an edge between them.
		currentStateName := c.GetCurrentState().GetText()
		// if currentstate doesn't exist as a node in the graph, add it
		if _, ok := m.Nodes[currentStateName]; !ok {
			n1 := m.Graph.MakeNode()
			*n1.Value = currentStateName
			m.Nodes[currentStateName] = n1

		}
		newstateName := c.GetNewState().GetText()
		// if newstate doesn't exist as a node in the graph, add it
		if _, ok := m.Nodes[newstateName]; !ok {
			n2 := m.Graph.MakeNode()
			*n2.Value = newstateName
			m.Nodes[newstateName] = n2
		}

		// finally add a transition between them
		m.Graph.MakeEdge(m.Nodes[currentStateName], m.Nodes[newstateName])
	}
}

func (m *MainProgramGraphBuilder) ExitMacroDef(c *parser.MacroDefContext) {
	m.inMacro = false
}
