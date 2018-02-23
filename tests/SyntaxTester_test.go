package testing

import (
	"github.com/mikkelmilo/Go-Turing-Machine/TM-Language"
	"testing"
)

func TestSyntaxChecker(t *testing.T) {
	_, err := TML.CheckSyntax("test.txt")
	if err != nil {
		panic(err)
	}
}
