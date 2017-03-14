package testing

import (
	"TuringMachine/TML"
	"testing"
)

func TestSyntaxChecker(t *testing.T) {
	_, err := TML.CheckSyntax("test.txt")
	if err != nil {
		panic(err)
	}
}
