package testing

import (
	"github.com/mikkelmilo/Go-Turing-Machine/examples"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
 * This file contains tests for the robustness of the Turing Machine implementation.
 * This mainly includes test for wrong/inappropriate usage, and other quirky usages of TM.
 */

var BIN_ALPHABET = []string{"0", "1"}

func TestBinaryIncTM(t *testing.T) {
	// the binary number we wish to increment is 011, which is expected to become 100
	err, tm := examples.IncBinaryTM([]string{"0", "1", "1"})
	assert.Nil(t, err)
	err = tm.Run()
	assert.Nil(t, err)
	// note that tm.Tape[0] is "reserved", so the binary number starts at tm.Tape[1]
	assert.Equal(t, tm.GetAlphabetMap()["1"], tm.GetTape()[1])
	assert.Equal(t, tm.GetAlphabetMap()["0"], tm.GetTape()[2])
	assert.Equal(t, tm.GetAlphabetMap()["0"], tm.GetTape()[3])
}

func TestInfiniteBinaryIncTM(t *testing.T) {
	//TODO: implement this test
}
