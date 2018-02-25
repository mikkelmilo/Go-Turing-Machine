package TM

import (
	"fmt"
	"github.com/micro/go-micro/errors"
)

// State struct for representing a state in the TM
type State struct {
	Name string
}

func (s *State) String() string {
	return s.Name
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

func (t Transition) String() string {
	return "(" +
		t.curState.String() +
		"," + t.newState.String() +
		"," + fmt.Sprintf("%c", t.curSymbol) +
		"," + fmt.Sprintf("%c", t.newSymbol) +
		"," + fmt.Sprintf("%c", t.dir) + ")"
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
// TODO: inject an alphabet (with max size |uint8| - 3) and let that alphabet be the
// allowed characters on the tape aside from the special characters '<', '>', and '_'
// build a mapping from the alphabet to uint8
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

// TODO: this operation may be very expensive and possibly redundant as well
// because the underlying array automatically expands when append() is called on a filled slice.
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
func (tm *TM) Run(state, quit chan int) error {
	var steps uint64
	steps = 0
	for tm.CurrentState != tm.AcceptState {
		select {
		case <-state:
			PrintTM(tm)
		case <-quit:
			quit <- 1
			return nil
		default:
			err := tm.Step()
			steps++
			if err != nil {
				quit <- -1
				return err
			}
		}
	}
	quit <- 1
	return nil
}

// Step : takes one step
func (tm *TM) Step() error {
	// first check if current state is nil; then set current state to the start state
	// else check if current state is either accept or reject. If so, report an error.
	if tm.CurrentState == nil {
		tm.Head = 1
		tm.CurrentState = tm.StartState

	} else if tm.CurrentState == tm.AcceptState {
		return errors.New(
			"1", "TM is already at the accept state and cannot make further transitions", 1)
	} else if tm.CurrentState == tm.RejectState {
		return errors.New(
			"2", "TM is already at the reject state and cannot make further transitions", -1)
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
		return fmt.Errorf("no transitions found on state %s with symbol %c", s.Name, symbol)
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
// TODO: implement the Stringer interface: change function signature to (tm *TM)String() -> string and return a string representation instead
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

func (tm *TM) String() string {
	a := tm.Tape[tm.Head]
	i := tm.Tape[0:tm.Head]
	j := tm.Tape[tm.Head+1 : len(tm.Tape)]
	p := make([]uint8, len(i))
	copy(p, i)
	p = append(p, LeftBracket)
	p = append(p, a)
	p = append(p, RightBracket)
	p = append(p, j...)
	nil_string := func(s *State) string {
		if s == nil {
			return "None"
		}
		return s.String()
	}
	tape := "Tape:\n" + fmt.Sprintf("%c \n", p)
	return "TM:\n" +
		"Reject state: " + nil_string(tm.RejectState) + "\n" +
		"Current state: " + nil_string(tm.CurrentState) + "\n" +
		"Transitions: " + fmt.Sprint(tm.Transitions) + "\n" +
		tape
}
