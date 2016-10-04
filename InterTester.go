package main

import (
	"TuringMachine/TML"
	"fmt"
)

func main() {
	tm := TML.Interpret()
	fmt.Println(tm.Transitions)
	fmt.Println(tm.Tape)
	er := tm.Run(nil, nil)
	fmt.Println(tm.Tape)

	fmt.Println(er)
}
