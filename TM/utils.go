package TM

type TMPrintListener struct {
}

func (TMPrintListener) step(tm *TM) {
	println("transitioned to state " + tm.CurrentState.String())
	println(tm.String())
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
