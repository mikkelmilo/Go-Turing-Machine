package Compiler

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseTMLProgram_illegalStateNames(t *testing.T) {
	prog1 := bytes.NewBufferString("(1aa,a,_,1,>)") // states cannot start with a number, so 1aa is illegal
	prog2 := bytes.NewBufferString("(_,a,_,1,>)")   // states cannot be named _
	prog3 := bytes.NewBufferString("(a,<,_,1,>)")   // states cannot be named <
	prog4 := bytes.NewBufferString("(a,>,_,1,>)")   // states cannot be named >

	errs1, _ := ParseTMLProgram(*prog1)
	errs2, _ := ParseTMLProgram(*prog2)
	errs3, _ := ParseTMLProgram(*prog3)
	errs4, _ := ParseTMLProgram(*prog4)
	assert.Len(t, errs1, 1)
	assert.Len(t, errs2, 1)
	assert.Len(t, errs3, 1)
	assert.Len(t, errs4, 1)
}
