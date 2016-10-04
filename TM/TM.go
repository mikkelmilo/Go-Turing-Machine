package TM

import "fmt"

// State struct for representing a state in the TM
type State struct {
	Name string
}

//TM a Turing Machine where the input- and output alphabets consist of 1 and 0.
type TM struct {
	StartState   *State
	AcceptState  *State
	RejectState  *State
	CurrentState *State
	Transitions  []Transition
	Tape         []uint8
	Head         int
}

// Transition : a representation of a transition
type Transition struct {
	curState  *State
	newState  *State
	curSymbol uint8
	newSymbol uint8
	dir       uint8
}

// One representation
const One uint8 = 49

// Zero representation
const Zero uint8 = 48

//Empty underscore
const Empty uint8 = 95

// LeftBracket representation
const LeftBracket = 123

// RightBracket representation
const RightBracket = 125

// Left arrow
const Left = 60

// Right arrow
const Right = 62

// NewTM constructor for the TM struct
func NewTM(s []uint8) TM {
	tm := TM{}
	tm.Head = 0
	tm.Tape = []uint8{Empty}
	if s != nil && len(s) != 0 {
		tm.Tape = append(tm.Tape, s...)
	} else {
		tm.Tape = []uint8{Empty, Empty}

	}
	return tm
}

/*
// AddInput add an input string to the TM
func (tm *TM) AddInput(s string) {
	a := 1
	for _, i := range s {
		if tm.Head+a >= len(tm.Tape) {
			tm.Tape = expandTape(tm.Tape)
		}
		tm.Tape[tm.Head+a] = i
		a++
	}
}*/

// doubles the size of the tape and places "_" symbols on the new slots
func expandTape(s []uint8) []uint8 {
	length := len(s)
	a := make([]uint8, length)
	for i := range a {
		a[i] = Empty
	}
	return append(s, a...)
}

// Run the TM until it halts (we assume it always halts)
func (tm *TM) Run(state, quit chan int) {
	var steps uint64
	steps = 0
	for tm.CurrentState != tm.AcceptState {
		select {
		case <-state:
			PrintTM(tm)
		case <-quit:
			quit <- 1
			return
		default:
			err := tm.Step()
			steps++
			if err != nil {
				fmt.Println(err)
				quit <- -1
			}
		}
	}
	quit <- 1
	return
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

func (tm *TM) makeTransition(s *State, symbol uint8) error {
	found := false
	for _, t := range tm.Transitions {
		if t.curState == s && t.curSymbol == symbol {
			tm.CurrentState = t.newState
			tm.Tape[tm.Head] = t.newSymbol
			if t.dir == One {
				tm.Head++
				if tm.Head >= len(tm.Tape) {
					tm.Tape = expandTape(tm.Tape)
				}
			} else if t.dir == Zero {
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
		return fmt.Errorf("no transitions found on state %s with symbol %d", s.Name, symbol)
	}
	return nil
}

//AddTransition asd
func (tm *TM) AddTransition(curState *State, newState *State, curSymbol string, newSymbol string, dir string) error {
	cSymbol := mapInput(curSymbol)
	nSymbol := mapInput(newSymbol)
	var cdir uint8
	if dir == "<" {
		cdir = Zero
	} else if dir == ">" {
		cdir = One
	} else if dir == "_" {
		cdir = Empty
	} else {
		return fmt.Errorf("Illegal argument: %s . Must be either <, > or _", dir)
	}

	if cSymbol == 3 {
		return fmt.Errorf("Illegal argument: %s. Must be 0 or 1 or _", curSymbol)
	} else if nSymbol == 3 {
		return fmt.Errorf("Illegal argument: %s . Must be 0 or 1 or _", newSymbol)
	}
	tm.Transitions = append(tm.Transitions, Transition{curState, newState, cSymbol, nSymbol, cdir})
	return nil
}

func mapInput(a string) uint8 {
	if a == "0" {
		return Zero
	} else if a == "1" {
		return One
	} else if a == "_" {
		return Empty
	}
	return 3
}

// SetStartState set the start state
func (tm *TM) SetStartState(s *State) {
	tm.StartState = s
}

// SetAcceptState set the accept state
func (tm *TM) SetAcceptState(s *State) {
	tm.AcceptState = s
}

// PrintTM prints tm tape
func PrintTM(tm *TM) {
	fmt.Println("Tape:")
	a := tm.Tape[tm.Head]
	i := tm.Tape[0:tm.Head]
	j := tm.Tape[tm.Head+1 : len(tm.Tape)]
	p := make([]uint8, len(i))
	copy(p, i)
	p = append(p, LeftBracket)
	p = append(p, a)
	p = append(p, RightBracket)
	p = append(p, j...)
	fmt.Printf("%c \n", p)
}
