package TM

import "fmt"

// State struct for representing a state in the TM
type State struct {
	Name string
}

//TM a Turing Machine where the input- and output alphabets consist of 1 and 0.
type TM struct {
	States       []*State
	StartState   *State
	AcceptState  *State
	CurrentState *State
	Transitions  []Transition
	Tape         []string
	Head         int
}

// Transition : a representation of a transition
type Transition struct {
	curState  *State
	newState  *State
	curSymbol string
	newSymbol string
	dir       string
}

// NewTM constructor for the TM struct
func NewTM(s string) TM {
	tm := TM{}
	tm.Head = 0
	tm.StartState = &State{"q_0"}
	tm.Tape = []string{"_"}
	if s != "" {
		tm.AddInput(s)
	}
	return tm
}

// AddInput add an input string to the TM
func (tm *TM) AddInput(s string) {
	a := 1
	for _, i := range s {
		if tm.Head+a >= len(tm.Tape) {
			tm.Tape = expandTape(tm.Tape)
		}
		tm.Tape[tm.Head+a] = fmt.Sprintf("%c", i)
		a++
	}
}

// doubles the size of the tape and places "_" symbols on the new slots
func expandTape(s []string) []string {
	length := len(s)
	a := make([]string, length)
	for i := range a {
		a[i] = "_"
	}
	return append(s, a...)
}

// Run the TM until it halts (we assume it always halts)
func (tm *TM) Run() error {
	printTM(tm)
	for tm.CurrentState != tm.AcceptState {
		err := tm.Step()
		if err != nil {
			return err
		}
		printTM(tm)

	}
	return nil
}

// Step : takes one step
func (tm *TM) Step() error {
	if tm.CurrentState == nil {
		tm.Head = 1
		tm.CurrentState = tm.StartState
	} else {
		symbol := tm.Tape[tm.Head]
		return tm.makeTransition(tm.CurrentState, symbol)
	}
	return nil
}

func (tm *TM) makeTransition(s *State, symbol string) error {
	found := false
	for _, t := range tm.Transitions {
		if t.curState == s && t.curSymbol == symbol {
			tm.CurrentState = t.newState
			tm.Tape[tm.Head] = t.newSymbol
			if t.dir == ">" {
				tm.Head++
				if tm.Head >= len(tm.Tape) {
					tm.Tape = expandTape(tm.Tape)
				}
			} else if t.dir == "<" {
				if tm.Head <= 0 {
					return fmt.Errorf("tried to move < out of bounds")
				}
				tm.Head--
			}
			found = true
			break
		}
	}
	if found == false && tm.CurrentState != tm.AcceptState {
		return fmt.Errorf("no transitions found on state %s with symbol %s", s.Name, symbol)
	}
	return nil
}

//AddTransition asd
func (tm *TM) AddTransition(curState *State, newState *State, curSymbol string, newSymbol string, dir string) error {
	if curSymbol != "0" && curSymbol != "1" && curSymbol != "_" {
		return fmt.Errorf("Illegal argument: %s. Must be a 0 or 1", curSymbol)
	} else if newSymbol != "0" && newSymbol != "1" && newSymbol != "_" {
		return fmt.Errorf("Illegal argument: %s . Must be a 0 or 1", newSymbol)
	} else if dir != "<" && dir != ">" && dir != "_" {
		return fmt.Errorf("Illegal argument: %s . Must be either <, > or _", dir)
	}
	tm.Transitions = append(tm.Transitions, Transition{curState, newState, curSymbol, newSymbol, dir})
	return nil
}

// SetStartState set the start state
func (tm *TM) SetStartState(s *State) {
	tm.StartState = s
}

// SetAcceptState set the accept state
func (tm *TM) SetAcceptState(s *State) {
	tm.AcceptState = s
}

// AddState add a state to the TM
func (tm *TM) AddState(s *State) {
	tm.States = append(tm.States, s)
}

func printTM(tm *TM) {
	fmt.Println("Tape:")
	a := tm.Tape[tm.Head]
	tm.Tape[tm.Head] = "{" + a + "}"
	fmt.Println(tm.Tape)
	tm.Tape[tm.Head] = a
}
