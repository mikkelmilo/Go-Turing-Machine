package Compiler

/*
	This file contains specific parse-tree walkers. (all implementing the TMLListener interface)
	This includes listeners for syntax Errors, pretty printing of the parse tree, unfolding of macro usages.
	These tools are solely for internal use.
*/

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mikkelmilo/Go-Turing-Machine/Compiler/antlr-parser"
	"github.com/twmb/algoimpl/go/graph"
	"strconv"
	"strings"
)

type TMLTreePrinterListener struct {
	*parser.BaseTMLListener
	nesting_lvl int
	errors      []antlr.ErrorNode
}

func nestingLvlString(lvl int) string {
	str := ""
	for i := 0; i < lvl; i++ {
		str += "  "
	}
	return str
}

func (t *TMLTreePrinterListener) EnterProgram(c *parser.ProgramContext) {
	t.nesting_lvl = 0
	println("Program:")
	t.nesting_lvl++
}

func (t *TMLTreePrinterListener) EnterMacroApp(c *parser.MacroAppContext) {
	name := c.GetToken(parser.TMLLexerID, 0).GetText()
	print(nestingLvlString(t.nesting_lvl) + "Macro Application: " + name + "(")
}

func (t *TMLTreePrinterListener) EnterMacroDef(c *parser.MacroDefContext) {
	println(nestingLvlString(t.nesting_lvl) + "Macro Definition: " +
		c.GetToken(parser.TMLLexerID, 0).GetText())
	t.nesting_lvl++
}

func (t *TMLTreePrinterListener) EnterCommand(c *parser.CommandContext) {
	print(nestingLvlString(t.nesting_lvl) + "Command: (")
}

func (t *TMLTreePrinterListener) EnterStateLabel(c *parser.StateLabelContext) {
	print(c.GetText() + " ")
}

func (t *TMLTreePrinterListener) EnterTapeSymbol(c *parser.TapeSymbolContext) {
	print(c.GetText() + " ")
}

func (t *TMLTreePrinterListener) EnterDirection(c *parser.DirectionContext) {
	print(c.GetText() + " ")
}

func (t *TMLTreePrinterListener) ExitProgram(c *parser.ProgramContext) {
	t.nesting_lvl--
}

func (t *TMLTreePrinterListener) ExitMacroApp(c *parser.MacroAppContext) {
	println(")")
}

func (t *TMLTreePrinterListener) ExitMacroDef(c *parser.MacroDefContext) {
	t.nesting_lvl--
}

func (t *TMLTreePrinterListener) ExitCommand(c *parser.CommandContext) {
	println(")")
}

type Command struct {
	CurrentState  string
	NewState      string
	CurrentSymbol string
	NewSymbol     string
	Direction     string
}

// A TML Tree listener which constructs a sequential program from a given TML Tree by unfolding all macro definitions
type TMLMacroUnfolder struct {
	*parser.BaseTMLListener
	Program      []Command
	Macros       map[string][]Command //maps from the macro name to its list of tuples
	currentMacro string
	uniqueNr     int
}

func (t *TMLMacroUnfolder) EnterProgram(c *parser.ProgramContext) {
	t.Program = make([]Command, 0)
	t.Macros = make(map[string][]Command)
	t.currentMacro = ""
	t.uniqueNr = 0
}

func (t *TMLMacroUnfolder) EnterMacroApp(c *parser.MacroAppContext) {
	macroName := c.GetToken(parser.TMLLexerID, 0).GetText()
	text := c.GetText()[1 : len(c.GetText())-1] //get text and remove surrounding ( and )
	text = strings.Replace(text, ")", ",", -1)
	text = strings.Replace(text, "(", ",", -1)
	//elems is a list of the different "elements" of this macro application
	elems := strings.Split(text, ",")

	curStateName := elems[0]
	curSymbol := elems[1]
	acceptState := elems[3]
	rejectState := elems[4]

	// make transition rules from curStatName to start starte of macro, and transitions when macro halts in accept or reject

	macro_hs_trans := Command{
		CurrentState:  curStateName,
		NewState:      "hs_" + macroName + strconv.Itoa(t.uniqueNr),
		CurrentSymbol: curSymbol,
		NewSymbol:     curSymbol,
		Direction:     "_",
	}
	macro_ha_trans := Command{
		CurrentState:  "ha_" + macroName + strconv.Itoa(t.uniqueNr),
		NewState:      acceptState,
		CurrentSymbol: "_",
		NewSymbol:     "_",
		Direction:     "_",
	}
	macro_hr_trans := Command{
		CurrentState:  "hr_" + macroName + strconv.Itoa(t.uniqueNr),
		NewState:      rejectState,
		CurrentSymbol: "_",
		NewSymbol:     "_",
		Direction:     "_",
	}

	// append these commands to the program
	t.Program = append(t.Program, macro_hs_trans, macro_ha_trans, macro_hr_trans)
	// generate the macro commands with unique state names
	macroCommands := t.GenerateUniqueStates(t.Macros[macroName], macroName, t.uniqueNr)
	t.uniqueNr++
	// append the macro to the program
	t.Program = append(t.Program, macroCommands...)
}

func (t *TMLMacroUnfolder) EnterMacroDef(c *parser.MacroDefContext) {
	t.currentMacro = c.GetToken(parser.TMLLexerID, 0).GetText()
}

func (t *TMLMacroUnfolder) EnterCommand(c *parser.CommandContext) {
	//therefore elems[0] will contain the current state string, elems[1] the new state string, elems[2] the current symbol, etc.
	//since we have already syntax checked the program, we can assume that this command is syntactically correct
	command := c.GetText()[1 : len(c.GetText())-1]
	elems := strings.Split(command, ",")
	cm := Command{
		CurrentState:  elems[0],
		NewState:      elems[1],
		CurrentSymbol: elems[2],
		NewSymbol:     elems[3],
		Direction:     elems[4],
	}
	// if we are not in a macro definition, add to t.Program list, else add to t.Macros[t.currentMacro]
	if t.currentMacro == "" {
		t.Program = append(t.Program, cm)
	} else {
		t.Macros[t.currentMacro] = append(t.Macros[t.currentMacro], cm)
	}
}

func (t *TMLMacroUnfolder) EnterDirection(c *parser.DirectionContext) {
	t.Program[len(t.Program)-1].Direction = c.GetText()
}

func (t *TMLMacroUnfolder) ExitMacroDef(c *parser.MacroDefContext) {
	t.currentMacro = ""
}

func (t *TMLMacroUnfolder) GenerateUniqueStates(c []Command, macroName string, seed int) []Command {
	res := make([]Command, len(c))
	for i, command := range c {
		res[i] = command
		res[i].CurrentState = res[i].CurrentState + "_" + macroName + strconv.Itoa(seed)
		res[i].NewState = res[i].NewState + "_" + macroName + strconv.Itoa(seed)
	}

	return res
}

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
		enteringstate_name := c.GetEnteringState().GetText()
		// if entering state doesn't exist as a node in the graph, add it
		if _, ok := m.Nodes[enteringstate_name]; !ok {
			n1 := m.Graph.MakeNode()
			*n1.Value = enteringstate_name
			m.Nodes[enteringstate_name] = n1

		}
		macroaccept_name := c.GetAcceptState().GetText()
		// if macro accept state doesn't exist as a node in the graph, add it
		if _, ok := m.Nodes[macroaccept_name]; !ok {
			n2 := m.Graph.MakeNode()
			*n2.Value = macroaccept_name
			m.Nodes[macroaccept_name] = n2
		}

		macroreject_name := c.GetRejectState().GetText()
		// if macro accept state doesn't exist as a node in the graph, add it
		if _, ok := m.Nodes[macroreject_name]; !ok {
			n2 := m.Graph.MakeNode()
			*n2.Value = macroreject_name
			m.Nodes[macroreject_name] = n2
		}

		m.Graph.MakeEdge(m.Nodes[enteringstate_name], m.Nodes[macroaccept_name])
		m.Graph.MakeEdge(m.Nodes[enteringstate_name], m.Nodes[macroreject_name])
	}
}

func (m *MainProgramGraphBuilder) EnterMacroDef(c *parser.MacroDefContext) {
	m.inMacro = true
}

func (m *MainProgramGraphBuilder) EnterCommand(c *parser.CommandContext) {
	// only add commands that are not inside macros
	if !m.inMacro {
		// add new nodes for current and new state, and make an edge between them.
		curstate_name := c.GetCurrentState().GetText()
		// if currentstate doesn't exist as a node in the graph, add it
		if _, ok := m.Nodes[curstate_name]; !ok {
			n1 := m.Graph.MakeNode()
			*n1.Value = curstate_name
			m.Nodes[curstate_name] = n1

		}
		newstate_name := c.GetNewState().GetText()
		// if newstate doesn't exist as a node in the graph, add it
		if _, ok := m.Nodes[newstate_name]; !ok {
			n2 := m.Graph.MakeNode()
			*n2.Value = newstate_name
			m.Nodes[newstate_name] = n2
		}

		// finally add a transition between them
		m.Graph.MakeEdge(m.Nodes[curstate_name], m.Nodes[newstate_name])
	}
}

func (m *MainProgramGraphBuilder) ExitMacroDef(c *parser.MacroDefContext) {
	m.inMacro = false
}
