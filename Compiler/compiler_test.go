package Compiler

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestCompileTMLProgram(t *testing.T) {
	b, err := ioutil.ReadFile("../examples/TML-programs/write_101.txt") // just pass the file name
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(b)
	errs, tm := CompileTMLProgram(*buf, ParseTMLProgram, CheckSemantic)
	assert.Nil(t, errs)
	println(tm.String())
	err = tm.Run(nil, nil)
	if err != nil {
		panic(err)
	}
	println(tm.String())

}
