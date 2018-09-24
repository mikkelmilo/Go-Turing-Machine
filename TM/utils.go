package TM

type TMListener interface {
	step(fromState *State, fromSymbol string, tm *TM)
	haltedWithAccept(tm *TM)
	haltedWithReject(tm *TM)
	haltedWithError(tm *TM, err error)
}

type TMPrintListener struct {
}

func (TMPrintListener) step(fromState *State, fromSymbol string, tm *TM) {
	println("transitioned from state " + fromState.String() + " with symbol " + fromSymbol + " to state " + tm.CurrentState.String())
}

func (TMPrintListener) haltedWithAccept(tm *TM) {
	println("halted with accept")
}

func (TMPrintListener) haltedWithReject(tm *TM) {
	println("halted with reject")
}

func (TMPrintListener) haltedWithError(tm *TM, err error) {
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
	inv_map := getInverseAlphabetMapping(alphabetMap)
	return "(" +
		t.CurState.String() +
		"," + t.NewState.String() +
		"," + inv_map[t.curSymbol] +
		"," + inv_map[t.newSymbol] +
		"," + string(t.dir) + ")"
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
			str = str + trans.asString(tm.AlphabetMap) + ","
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
