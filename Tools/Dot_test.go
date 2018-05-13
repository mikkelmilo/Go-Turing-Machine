package Tools

import (
	"bytes"
	C "github.com/mikkelmilo/Go-Turing-Machine/Compiler"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestTMToDotFile(t *testing.T) {
	b, err := ioutil.ReadFile("../examples/TML-programs/binary_increment.txt") // just pass the file name
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(b)
	_, tm := C.CompileTMLProgram(*buf, C.ParseTMLProgram, C.CheckSemantic)
	assert.NoError(t, TMToDotFile(tm, "tm.dot"))
}
