package Compiler

import (
	"bytes"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mikkelmilo/Go-Turing-Machine/Compiler/antlr-parser"
	"github.com/mikkelmilo/Go-Turing-Machine/TM"
	"strconv"
	"strings"
)

/*
	This file contains the compiler.
	The compiler takes a source file and checks whether the program is valid.
	If so, it will return a Turing Machine which emulates this program when run.
*/

type TMLCompiler func(bytes.Buffer, TMLParser, TMLSemanticChecker) ([]TMLError, TM.TM)

func CompileTMLProgram(program bytes.Buffer, parser TMLParser, semantChecker TMLSemanticChecker) ([]TMLError, TM.TM) {
	// parse the program, and report any errors.
	errs, syntaxTree := parser(program)
	if len(errs) != 0 {
		return errs, TM.TM{}
	}
	// do semantic checking, and report any errors.
	errs = semantChecker(antlr.ParseTreeWalkerDefault, syntaxTree)
	if len(errs) != 0 {
		return errs, TM.TM{}
	}

	// assuming no other errors, construct the TM. First we need to unfold macro applications.
	var macroUnfolder TMLMacroUnfolder
	antlr.ParseTreeWalkerDefault.Walk(&macroUnfolder, syntaxTree)
	// unfoldedProgram is a list of commands where all macro applications have been replaced by its macro definition.
	// and everything has been neatly merged into a sequential program.
	unfoldedProgram := macroUnfolder.Program
	// before constructing the TM we must first identify the alphabet and the states.
	alphabet := make(map[string]bool, 0)
	states := make(map[string]*TM.State)

	for _, command := range unfoldedProgram {
		cursym := command.CurrentSymbol
		newsym := command.NewSymbol
		// add symbols to alphabet if they don't already exist.
		if _, exists := alphabet[cursym]; !exists && cursym != "_" {
			alphabet[cursym] = true
		}
		if _, exists := alphabet[newsym]; !exists && newsym != "_" {
			alphabet[newsym] = true
		}

		// add states whenever we see new ones
		curState := command.CurrentState
		newState := command.NewState
		if _, exists := states[curState]; !exists {
			states[curState] = &TM.State{Name: curState}
		}
		if _, exists := states[newState]; !exists {
			states[newState] = &TM.State{Name: newState}
		}
	}
	alphabet_list := make([]string, len(alphabet))
	i := 0
	for key := range alphabet {
		alphabet_list[i] = key
		i++
	}

	err, tm := TM.NewTM(alphabet_list, nil)
	if err != nil {
		return []TMLError{{msg: err.Error()}}, tm
	}
	// add transitions to the TM
	for _, c := range unfoldedProgram {
		tm.AddTransition(states[c.CurrentState], states[c.NewState], c.CurrentSymbol, c.NewSymbol, c.Direction)
		if c.CurrentState == "hs" {
			tm.SetStartState(states[c.CurrentState])
		}
		if c.NewState == "ha" {
			tm.SetAcceptState(states[c.NewState])
		}
		if c.NewState == "hr" {
			tm.SetRejctState(states[c.NewState])
		}
	}
	return nil, tm
}

//---------------------------------------------------------------------------------------

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
	cm := Command{
		CurrentState:  c.GetCurrentState().GetText(),
		NewState:      c.GetNewState().GetText(),
		CurrentSymbol: c.GetCurrentSymbol().GetText(),
		NewSymbol:     c.GetNewSymbol().GetText(),
		Direction:     c.GetDir().GetText(),
	}
	// if we are not in a macro definition, add to t.Program list, else add to t.Macros[t.currentMacro]
	if t.currentMacro == "" {
		t.Program = append(t.Program, cm)
	} else {
		t.Macros[t.currentMacro] = append(t.Macros[t.currentMacro], cm)
	}
}

func (t *TMLMacroUnfolder) EnterDirection(c *parser.DirectionContext) {
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
