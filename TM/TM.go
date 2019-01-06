package TM

import (
	"errors"
	"fmt"
)

const (
	//empty underscore
	empty uint8 = 95

	// leftBracket representation
	leftBracket = 123

	// rightBracket representation
	rightBracket = 125

	// left arrow
	left = 60

	// right arrow
	right = 62
)

type TM interface {
	SetStartState(s *State)
	SetAcceptState(s *State)
	SetRejectState(s *State)
	GetTransitions() []Transition
	GetStartState() *State
	GetAcceptState() *State
	GetRejectState() *State
	GetCurrentState() *State
	GetTape() []uint8
	GetHead() int
	GetAlphabet() []string
	GetAlphabetMap() map[string]uint8
	SetTransitions(t []Transition) error
	AddListener(l TMListener)
	RemoveListener(l TMListener) error
	Run() error
	Step() error
	AddTransition(curState *State, newState *State, curSymbol string, newSymbol string, dir string) error
	String() string
}

// Transition : a representation of a transition
type Transition struct {
	CurState  *State
	NewState  *State
	curSymbol uint8
	newSymbol uint8
	dir       uint8
}

// State struct for representing a state in the tmImpl
type State struct {
	Name string
}

func (s *State) String() string {
	return s.Name
}

//tmImpl a Turing Machine where the input- and output alphabets consist of 1 and 0.
type tmImpl struct {
	StartState   *State
	AcceptState  *State
	RejectState  *State
	CurrentState *State
	Transitions  []Transition
	Tape         []uint8
	Head         int
	Alphabet     []string
	AlphabetMap  map[string]uint8
	listeners    []TMListener
}

func (tm *tmImpl) SetTransitions(t []Transition) error {
	tm.Transitions = t
	return nil
}

func (tm *tmImpl) GetAlphabet() []string {
	return tm.Alphabet
}

func (tm *tmImpl) GetAlphabetMap() map[string]uint8 {
	return tm.AlphabetMap
}

func (tm *tmImpl) GetTransitions() []Transition {
	return tm.Transitions
}

func (tm *tmImpl) GetStartState() *State {
	return tm.StartState
}

func (tm *tmImpl) GetAcceptState() *State {
	return tm.AcceptState
}

func (tm *tmImpl) GetRejectState() *State {
	return tm.RejectState
}

func (tm *tmImpl) GetCurrentState() *State {
	return tm.CurrentState
}

func (tm *tmImpl) GetTape() []uint8 {
	return tm.Tape
}

func (tm *tmImpl) GetHead() int {
	return tm.Head
}

/*
 * NewTM constructs a tmImpl from the specified alphabet, and optional initial tape.
 * @precondition: max alphabet size is 251
 */
func NewTM(alphabet []string, startTape []string) (error, TM) {
	tm := tmImpl{}

	if alphabet == nil {
		return errors.New("alphabet must be different from nil"), &tm
	}
	// report an error on too large alphabet (> 253)
	if len(alphabet) > 251 {
		return fmt.Errorf("alphabet size too large. Maximal size: 251. Got: %v", len(alphabet)), &tm
	}
	// build alphabet map to uint8 on which the tmImpl will operate
	// uint8 values 95, 60, 62, 123, and 125 are reserved.
	tm.Alphabet = alphabet
	tm.AlphabetMap = buildAlphabetMap(alphabet)

	tm.Head = 0
	// initially, set tape to an array of either two 'empty' elements,
	// or if a tape was given, set tape to [empty :: startTape]
	var len_s int
	if startTape == nil {
		len_s = 0
	} else {
		len_s = len(startTape)
	}
	if startTape != nil && len_s != 0 {
		// translate given tape to type []uint8 and append
		s_trans := make([]uint8, len_s+1)
		s_trans[0] = empty
		for i, elem := range startTape {
			s_trans[i+1] = tm.AlphabetMap[elem]
		}
		tm.Tape = s_trans
	} else {
		tm.Tape = []uint8{empty, empty}

	}
	//finally, initialize the list of listeners to be empty
	tm.listeners = make([]TMListener, 0)
	return nil, &tm
}

// SetStartState set the start state
func (tm *tmImpl) SetStartState(s *State) {
	tm.StartState = s
}

// SetAcceptState set the accept state
func (tm *tmImpl) SetAcceptState(s *State) {
	tm.AcceptState = s
}

// SetRejectState set the reject state
func (tm *tmImpl) SetRejectState(s *State) {
	tm.RejectState = s
}

func (tm *tmImpl) AddListener(l TMListener) {
	tm.listeners = append(tm.listeners, l)
}

/*
	Removes all instances of a listener from the tmImpl. Reports an error if listener is not found.
*/
func (tm *tmImpl) RemoveListener(l TMListener) error {
	found := false
	// remove all instances of this listener from the tmImpl
	for i, e := range tm.listeners {
		if e == l {
			found = true
			tm.listeners = remove(tm.listeners, i)
		}
	}
	// report error if listener does not exist.
	if !found {
		return errors.New("tried to remove a listener which doesn't exist on this object")
	}
	return nil
}

/*
	Removes all listeners from the tmImpl
*/
func (tm *tmImpl) RemoveListeners() {
	tm.listeners = make([]TMListener, 0)
}

// TODO: this operation may be very expensive and possibly redundant as well
// because the underlying array automatically expands when append() is called on a filled slice.
// doubles the size of the tape and places "_" symbols on the new slots
func expandTape(s []uint8) []uint8 {
	length := len(s)
	a := make([]uint8, length)
	for i := range a {
		a[i] = empty
	}
	return append(s, a...)
}

// Run the tmImpl until it halts (we assume it always halts)
func (tm *tmImpl) Run() error {
	var steps uint64
	steps = 0
	if tm.StartState == nil {
		return errors.New("no start state defined")
	}
	for tm.CurrentState != tm.AcceptState && tm.CurrentState != tm.RejectState {
		err := tm.Step()
		steps++
		if err != nil {
			return err

		}
	}
	return nil
}

// Step : takes one step
func (tm *tmImpl) Step() error {
	// first check if current state is nil; then set current state to the start state
	// else check if current state is either accept or reject. If so, report an error.
	if tm.CurrentState == nil {
		tm.Head = 0
		tm.CurrentState = tm.StartState

	} else if tm.CurrentState == tm.AcceptState {
		err := errors.New("tmImpl is already at the accept state and cannot make further transitions")
		for _, l := range tm.listeners {
			l.haltedWithAccept(tm)
		}
		return err
	} else if tm.CurrentState == tm.RejectState {
		for _, l := range tm.listeners {
			l.haltedWithReject(tm)
		}
		return nil
	}

	symbol := tm.Tape[tm.Head]
	// try to make transition. Report error if failed, and notify listeners
	err := tm.makeTransition(tm.CurrentState, symbol)
	if err != nil {
		for _, l := range tm.listeners {
			l.haltedWithError(tm, err)
		}
	} else {
		for _, l := range tm.listeners {
			inv_map := getInverseAlphabetMapping(tm.AlphabetMap)
			l.step(tm.CurrentState, inv_map[symbol], tm)
		}
	}
	return err
}

func (tm *tmImpl) makeTransition(s *State, symbol uint8) error {
	found := false
	for _, t := range tm.Transitions {
		if t.CurState == s && t.curSymbol == symbol {
			tm.CurrentState = t.NewState
			tm.Tape[tm.Head] = t.newSymbol
			if t.dir == right {
				tm.Head++
				if tm.Head >= len(tm.Tape) {
					tm.Tape = expandTape(tm.Tape)
				}
			} else if t.dir == left {
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

//AddTransition adds a transition to the tmImpl
func (tm *tmImpl) AddTransition(curState *State, newState *State, curSymbol string, newSymbol string, dir string) error {
	cSymbol, ok1 := tm.AlphabetMap[curSymbol]
	if !ok1 {
		return fmt.Errorf("symbol %v is not in the alphabet %v", curSymbol, tm.Alphabet)
	}
	nSymbol, ok2 := tm.AlphabetMap[newSymbol]
	if !ok2 {
		return fmt.Errorf("symbol %v is not in the alphabet %v", newSymbol, tm.Alphabet)
	}

	var cdir uint8
	if dir == "<" {
		cdir = left
	} else if dir == ">" {
		cdir = right
	} else if dir == "_" {
		cdir = empty
	} else {
		return fmt.Errorf("illegal argument: %s . Must be either <, > or _", dir)
	}

	if cSymbol == 3 {
		return fmt.Errorf("illegal argument: %s. Must be 0 or 1 or _", curSymbol)
	} else if nSymbol == 3 {
		return fmt.Errorf("illegal argument: %s . Must be 0 or 1 or _", newSymbol)
	}
	tm.Transitions = append(tm.Transitions, Transition{curState, newState, cSymbol, nSymbol, cdir})
	return nil
}

func buildAlphabetMap(alphabet []string) map[string]uint8 {
	alphabetMap := make(map[string]uint8)
	alphabetMap["_"] = empty
	alphabetMap["<"] = left
	alphabetMap[">"] = right
	alphabetMap["{"] = leftBracket
	alphabetMap["}"] = rightBracket

	var cur_uint uint8 = 0
	for _, char := range alphabet {
		// note: this suffices since we know neither of the contants follow each other in value.
		if cur_uint == empty || cur_uint == left || cur_uint == right ||
			cur_uint == leftBracket || cur_uint == rightBracket {
			cur_uint++
		}
		alphabetMap[char] = cur_uint
		cur_uint++
	}
	return alphabetMap
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
