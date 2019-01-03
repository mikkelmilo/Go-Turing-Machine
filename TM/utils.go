package TM

import "fmt"

type TMListener interface {
	step(fromState *State, fromSymbol string, tm TM)
	haltedWithAccept(tm TM)
	haltedWithReject(tm TM)
	haltedWithError(tm TM, err error)
}

type TMPrintListener struct {
}

func (TMPrintListener) step(fromState *State, fromSymbol string, tm TM) {
	println("transitioned from state " + fromState.String() + " with symbol " + fromSymbol + " to state " + tm.GetCurrentState().String())
}

func (TMPrintListener) haltedWithAccept(tm TM) {
	println("halted with accept")
}

func (TMPrintListener) haltedWithReject(tm TM) {
	println("halted with reject")
}

func (TMPrintListener) haltedWithError(tm TM, err error) {
	println("halted with error:", err.Error())
}

/* Insert inserts the value into the slice at the specified index, which must be in range.
 * The slice must have room for the new element.
 */
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

// removes an element from a list. Does not guarantee order.
func remove(s []TMListener, i int) []TMListener {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

// ### Printing-related functions ####

func (t Transition) asString(alphabetMap map[string]uint8) string {
	invMap := getInverseAlphabetMapping(alphabetMap)
	return "(" +
		t.CurState.String() +
		"," + t.NewState.String() +
		"," + invMap[t.curSymbol] +
		"," + invMap[t.newSymbol] +
		"," + string(t.dir) + ")"
}

func (t Transition) GetCurSymbol(tm TM) string {
	invMap := getInverseAlphabetMapping(tm.GetAlphabetMap())
	return invMap[t.curSymbol]
}
func (t Transition) GetNewSymbol(tm TM) string {
	invMap := getInverseAlphabetMapping(tm.GetAlphabetMap())
	return invMap[t.newSymbol]
}
func (t Transition) GetDir() string {
	return string(t.dir)
}

func (tm *tmImpl) String() string {
	// convert tape into an array of characters from the alphabet
	tapeFormatted := make([]string, len(tm.GetTape()))
	invMap := getInverseAlphabetMapping(tm.GetAlphabetMap())
	for i, char := range tm.Tape {
		tapeFormatted[i] = invMap[char]
	}
	//insert { } around the current position on the tape
	t1 := Insert(tapeFormatted, tm.GetHead(), "{")
	t2 := Insert(t1, tm.GetHead()+2, "}")

	transitionsString := func(t []Transition) string {
		str := "["
		for _, trans := range t {
			str = str + trans.asString(tm.GetAlphabetMap()) + ","
		}
		str = str[0 : len(str)-1] // trim last ','
		str = str + "]"
		return str
	}

	nilString := func(s *State) string {
		if s == nil {
			return "None"
		}
		return s.String()
	}
	return "TM:\n" +
		"Alphabet: " + fmt.Sprintf("%v \n", tm.Alphabet) +
		"Reject state: " + nilString(tm.RejectState) + "\n" +
		"Current state: " + nilString(tm.CurrentState) + "\n" +
		"Transitions: " + transitionsString(tm.Transitions) + "\n" +
		"Tape:\n" + fmt.Sprintf("%v \n", t2)
}
