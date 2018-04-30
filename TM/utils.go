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
