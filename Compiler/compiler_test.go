package Compiler

import (
	"bytes"
	"github.com/mikkelmilo/Go-Turing-Machine/TM"
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
	err = tm.Run(nil, nil)
	assert.Nil(t, err)
	println("Result of compiling and running write_101.txt:")
	println(tm.String())

}

func TestCompileTMLProgramOnBinary_increment(t *testing.T) {
	b, err := ioutil.ReadFile("../examples/TML-programs/binary_increment.txt") // just pass the file name
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(b)
	errs, tm := CompileTMLProgram(*buf, ParseTMLProgram, CheckSemantic)
	assert.Nil(t, errs)
	println(tm.String())
	var l TM.TMPrintListener
	tm.AddListener(&l)
	err = tm.Run(nil, nil)
	assert.Nil(t, err)
	println("Result of compiling and running write_101.txt:")
	println(tm.String())

}
