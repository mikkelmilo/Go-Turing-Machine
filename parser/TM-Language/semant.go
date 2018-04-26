package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/twmb/algoimpl/go/graph"
)

/*
 * This file contains a Tree walker which checks the semantics of the program.
 * This includes:
 *    - the program contains hs, ha, and hr states
 *    - all states must be reachable from hs (the same is true for Macros and their states)
 *    - Macros must also contain hs, ha, and hr states
 *    - Nice-to-have: warnings will be produced (if enabled) if there are unbreakable cycles, ie.
 *      Nice-to-have: cycles which have no sequence of transitions to either ha or hr
 */

type SemanticChecker func(pt *antlr.ParseTreeWalker, tree IStartContext) []TMLError

func CheckSemantic(pt *antlr.ParseTreeWalker, tree IStartContext) []TMLError {

	// first walk the concrete syntax tree and report any parser errors found
	var walker semantTreeListener
	pt.Walk(&walker, tree)

	if len(walker.errors) != 0 {
		return walker.errors
	}

	// if no errors were found in the tree walk above, check for reachability properties of hs, ha, and hr (in both main program and macros)
	var unfolder TMLMacroUnfolder
	pt.Walk(&unfolder, tree)

	// for each local macro, check reachability of nodes and report and error if there are any.
	for macroname, m := range unfolder.Macros {
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
			errs[i] = TMLError{msg: "state " + (*n.Value).(string) + " is unreachable in macro: " + macroname}
		}
		if len(errs) != 0 {
			return errs
		}
	}

	// check the main program for reachability
	var mainprogbuilder MainProgramGraphBuilder
	pt.Walk(&mainprogbuilder, tree)

	unreachables := FindUnreachableNodes(*mainprogbuilder.Graph, mainprogbuilder.Nodes)
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
	for n, was_visited := range visited {
		if !was_visited {
			unvisited = append(unvisited, n)
		}
	}
	return unvisited
}

type semantTreeListener struct {
	*BaseTMLListener
	errors             []TMLError
	inMacro            bool
	seenStartState     bool
	seenAcceptState    bool
	seenRejectState    bool
	startStateChanged  bool
	acceptStateChanged bool
	rejectStateChanged bool
}

func (semant *semantTreeListener) EnterProgram(c *ProgramContext) {
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

func (semant *semantTreeListener) EnterMacroDef(c *MacroDefContext) {
	semant.seenRejectState = false
	semant.seenAcceptState = false
	semant.seenStartState = false
	semant.inMacro = true

}

func (semant *semantTreeListener) EnterCommand(c *CommandContext) {
	check_command := func(state_name string, state_type string) {
		switch state_name {
		case "hs":
			if !semant.seenStartState {
				semant.seenStartState = true
				if !semant.inMacro {
					semant.startStateChanged = true
				}
			} else if state_type == "currentstate" {
				// else if the current state symbol is the start state, but not the first one
				// seen in this scope, add an error.
				semant.AppendErrorMsg("Multiple start states defined", c.GetStart())
			}
		case "ha":
			if !semant.seenAcceptState {
				semant.seenAcceptState = true
				if !semant.inMacro {
					semant.acceptStateChanged = true
				}
			} else if state_type == "currentstate" {
				semant.AppendErrorMsg("Accept state cannot have transitions to other states", c.GetStart())
			}
		case "hr":
			if !semant.seenRejectState {
				semant.seenRejectState = true
				if !semant.inMacro {
					semant.rejectStateChanged = true
				}
			} else if state_type == "currentstate" {
				semant.AppendErrorMsg("Reject state cannot have transitions to other states", c.GetStart())
			}
		}
	}

	check_command(c.GetCurrentState().GetText(), "currentstate")
	check_command(c.GetNewState().GetText(), "newstate")
}

func (semant *semantTreeListener) ExitProgram(c *ProgramContext) {
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

func (semant *semantTreeListener) ExitMacroDef(c *MacroDefContext) {
	semant.inMacro = false
	macroName := c.GetToken(TMLParserID, 0).GetText()
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
