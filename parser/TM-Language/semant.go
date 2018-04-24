package parser

import (
	"errors"
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

type semanticChecker = func(pt *antlr.ParseTreeWalker, tree IStartContext) []TMLError

func CheckSemantic(pt *antlr.ParseTreeWalker, tree IStartContext) []TMLError {

	// first walk the concrete syntax tree and report any errors found
	walker := semantTreeListener{}
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
	//TODO: build main program graph and check for unreachable nodes in it.
	return nil
}

//TODO: implement. make a node for each normal command, and add edges between nodes in macro applications.
func BuildMainProgramGraph(pt *antlr.ParseTreeWalker, tree IStartContext) ([]graph.Node, graph.Graph) {
	return nil, graph.Graph{}

	//make a treelistener implementation which traverses the concrete syntax tree and does as described above.
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
	errors             []TMLError
	inMacro            bool
	seenStartState     bool
	seenAcceptState    bool
	seenRejectState    bool
	startStateChanged  bool
	acceptStateChanged bool
	rejectStateChanged bool
}

func (semant *semantTreeListener) VisitTerminal(node antlr.TerminalNode) {
}
func (semant *semantTreeListener) VisitErrorNode(node antlr.ErrorNode) {
}
func (semant *semantTreeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
}
func (semant *semantTreeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
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
	switch c.GetCurrentState().GetText() {
	case "hs":
		if !semant.seenStartState {
			semant.seenStartState = true
			if !semant.inMacro {
				semant.startStateChanged = true
			}
		} else {
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
		} else {
			semant.AppendErrorMsg("Multiple accept states defined", c.GetStart())
		}
	case "hr":
		if !semant.seenRejectState {
			semant.seenRejectState = true
			if !semant.inMacro {
				semant.rejectStateChanged = true
			}
		} else {
			semant.AppendErrorMsg("Multiple reject states defined", c.GetStart())
		}
	}

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

/*
	Checks a TML program (in sequential form) for the following reachability properties:
	The accept and reject states must be reachable from the start state (all of which are assumed to exist in the given input)
	Furthermore, it returns a list of indices into the program list where it holds that the current state of the command is unreachable.
*/
func CheckReachability(program []Command) (error, []int) {
	visited := make([]bool, len(program))
	for i, _ := range visited {
		visited[i] = false
	}

	startIndex := -1
	for i, c := range program {
		if c.CurrentState == "hs" {
			startIndex = i
			break
		}
	}
	if startIndex == -1 {
		panic("Unexpected situation: start state not found in program while performing state reachability analysis")
	}
	remaining := []int{startIndex} // a list of currently marked nodes whose lambda closure has not yet been found

	for len(remaining) != 0 {
		i := remaining[len(remaining)-1] // take out last element in 'remaining'
		visited[i] = true
		remaining = remaining[0 : len(remaining)-1] //shrink array by one - ie remove last element
		next := program[i].NewState
		// search for all occurances of the newState as 'currentState' in other commands
		// then add these to the remaining list
		for j, c := range program {
			if c.CurrentState == next {
				remaining = append(remaining, j)
			}
		}
	}

	// check if both accept and reject states were reachable from the start state. If not, return an error.
	for i, isVisited := range visited {
		if !isVisited && program[i].NewState == "ha" {
			err := errors.New("The accept state is not reachable from the start state")
			return err, nil
		}
		if !isVisited && program[i].NewState == "hr" {
			err := errors.New("The reject state is not reachable from the start state")
			return err, nil
		}
	}

	// else if no errors return all those indices (in program) that have not been visited
	result := []int{}
	for i, isVisited := range visited {
		if !isVisited {
			result = append(result, i)
		}
	}
	return nil, result
}
