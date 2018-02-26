package TM

import (
	"errors"
	"fmt"
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
	Alphabet     []string
	AlphabetMap  map[string]uint8
}

// Transition : a representation of a transition
type Transition struct {
	curState  *State
	newState  *State
	curSymbol uint8
	newSymbol uint8
	dir       uint8
}

func (t Transition) String(alphabetMap map[string]uint8) string {
	inv_map := getInverseAlphabetMapping(alphabetMap)
	return "(" +
		t.curState.String() +
		"," + t.newState.String() +
		"," + inv_map[t.curSymbol] +
		"," + inv_map[t.newSymbol] +
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

/*
 * NewTM constructs a TM from the specified alphabet, and optional initial tape.
 * @precondition: max alphabet size is 251
 */
func NewTM(alphabet []string, s []string) (error, TM) {
	tm := TM{}

	if alphabet == nil {
		return errors.New("Alphabet must be different from nil"), tm
	}
	// report an error on too large alphabet (> 253)
	if len(alphabet) > 251 {
		return fmt.Errorf("Alphabet size too large. Maximal size: 251. Got: %v", len(alphabet)), tm
	}
	// build alphabet map to uint8 on which the TM will operate
	// uint8 values 95, 60, 62, 123, and 125 are reserved.
	tm.Alphabet = alphabet
	alphabetMap := make(map[string]uint8)
	alphabetMap["_"] = Empty
	alphabetMap["<"] = Left
	alphabetMap[">"] = Right
	alphabetMap["{"] = LeftBracket
	alphabetMap["}"] = RightBracket

	var cur_uint uint8 = 0
	for _, char := range alphabet {
		// note: this suffices since we know neither of the contants follow each other in value.
		if cur_uint == Empty || cur_uint == Left || cur_uint == Right ||
			cur_uint == LeftBracket || cur_uint == RightBracket {
			cur_uint++
		}
		alphabetMap[char] = cur_uint
		cur_uint++
	}
	tm.AlphabetMap = alphabetMap

	tm.Head = 0
	// initially, set tape to an array of either two 'Empty' elements,
	// or if a tape was given, set tape to [Empty :: s]
	var len_s int
	if s == nil {
		len_s = 0
	} else {
		len_s = len(s)
	}
	if s != nil && len_s != 0 {
		// translate given tape to type []uint8 and append
		s_trans := make([]uint8, len_s+1)
		s_trans[0] = Empty
		for i, elem := range s {
			s_trans[i+1] = alphabetMap[elem]
		}
		tm.Tape = s_trans
	} else {
		tm.Tape = []uint8{Empty, Empty}

	}
	return nil, tm
}

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
func (tm *TM) Run(state chan string, quit chan int) error {
	var steps uint64
	steps = 0
	if tm.StartState == nil {
		return nil
	}
	for tm.CurrentState != tm.AcceptState {
		select {
		case <-state:
			state <- tm.String()
		case <-quit:
			quit <- 1
			return nil
		default:
			err := tm.Step()
			steps++
			if err != nil {
				if quit != nil {
					quit <- -1
				}
				return err
			}
		}
	}
	if quit != nil {
		quit <- 1
	}
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
		return errors.New("TM is already at the accept state and cannot make further transitions")
	} else if tm.CurrentState == tm.RejectState {
		return errors.New("TM is already at the reject state and cannot make further transitions")
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

//AddTransition adds a transition to the TM
func (tm *TM) AddTransition(curState *State, newState *State, curSymbol string, newSymbol string, dir string) error {
	cSymbol, ok1 := tm.AlphabetMap[curSymbol]
	if !ok1 {
		return fmt.Errorf("Symbol %v is not in the alphabet %v", curSymbol, tm.Alphabet)
	}
	nSymbol, ok2 := tm.AlphabetMap[newSymbol]
	if !ok2 {
		return fmt.Errorf("Symbol %v is not in the alphabet %v", newSymbol, tm.Alphabet)
	}

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

// SetStartState set the start state
func (tm *TM) SetStartState(s *State) {
	tm.StartState = s
}

// SetAcceptState set the accept state
func (tm *TM) SetAcceptState(s *State) {
	tm.AcceptState = s
}

func (tm *TM) String() string {
	// convert tape into an array of characters from the alphabet
	tape_formatted := make([]string, len(tm.Tape))
	inv_map := getInverseAlphabetMapping(tm.AlphabetMap)
	for i, char := range tm.Tape {
		tape_formatted[i] = inv_map[char]
	}
	//insert { } around the current position on the tape

	t1 := Insert(tape_formatted, tm.Head, "{")
	t2 := Insert(t1, tm.Head+2, "}")

	transitions_string := func(t []Transition) string {
		str := "["
		for _, trans := range t {
			str = str + trans.String(tm.AlphabetMap) + ","
		}
		str = str[0 : len(str)-1] // trim last ','
		str = str + "]"
		return str
	}

	nil_string := func(s *State) string {
		if s == nil {
			return "None"
		}
		return s.String()
	}
	return "TM:\n" +
		"Alphabet: " + fmt.Sprintf("%v \n", tm.Alphabet) +
		"Reject state: " + nil_string(tm.RejectState) + "\n" +
		"Current state: " + nil_string(tm.CurrentState) + "\n" +
		"Transitions: " + transitions_string(tm.Transitions) + "\n" +
		"Tape:\n" + fmt.Sprintf("%v \n", t2)
}

/*
 * Reverses the alphabet mapping. Assumes alphabetMap is reversible.
 */
func getInverseAlphabetMapping(alphabetMap map[string]uint8) map[uint8]string {
	res := make(map[uint8]string)
	for key, value := range alphabetMap {
		res[value] = key
	}
	return res
}

// Insert inserts the value into the slice at the specified index,
// which must be in range.
// The slice must have room for the new element.
func Insert(slice []string, index int, value string) []string {
	// Grow the slice by one element.
	r := append(slice, "")
	// Use copy to move the upper part of the slice out of the way and open a hole.
	copy(r[index+1:], r[index:])
	// Store the new value.
	r[index] = value
	// Return the result.
	return r
}
