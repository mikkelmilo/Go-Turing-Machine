package Compiler

import (
	"bytes"
	"github.com/mikkelmilo/Go-Turing-Machine/TM"
)

/*
	This file contains the compiler.
	The compiler takes a source file and checks whether the program is valid.
	If so, it will return a Turing Machine which emulates this program when run.
*/

type TMLCompiler func(buffer bytes.Buffer) ([]TMLError, TM.TM)
