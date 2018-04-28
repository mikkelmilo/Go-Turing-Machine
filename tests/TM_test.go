package testing

import (
	"github.com/mikkelmilo/Go-Turing-Machine/TM"
	"github.com/mikkelmilo/Go-Turing-Machine/examples"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
 * This file contains tests for the robustness of the Turing Machine implementation.
 * This mainly includes test for wrong/inappropriate usage, and other quirky usages of TM.
 */

var BIN_ALPHABET = []string{"0", "1"}

const OK_CODE = 1
const ERROR_CODE = -1

// Tests that when TM.run() is called with a nil channel, it will handle it appropriately
func TestRunTM_nilchans(t *testing.T) {
	//first test on dummy TM
	err, tm := TM.NewTM(BIN_ALPHABET, nil)
	AssertOk(t, err)
	err = tm.Run(nil, nil)

	AssertOk(t, err)

	// now test on TM generated by a TML program

	err1, tm1 := TML.Interpret("../examples/TML-programs/write_101.txt")
	assert.Nil(t, err1)
	err1 = tm1.Run(nil, nil)
	assert.Nil(t, err1)
}

func TestRunTMChanBehaviour(t *testing.T) {
	// first, construct tm from the program in write_101.txt, and check that there were no errors during construction.
	err1, tm1 := TML.Interpret("../examples/TML-programs/write_101.txt")
	assert.Nil(t, err1)

	//make channels and test behaviour
	quit := make(chan int)
	c := make(chan string)
	var err error
	go func() {
		err = tm1.Run(c, quit)
	}()
	// get current "state" of the TM, as a string representation
	c <- ""          // send request for state
	tm_string := <-c // wait for answer
	assert.NotNil(t, tm_string)
	res_err_code := <-quit                 // wait for the TM to halt
	assert.Equal(t, OK_CODE, res_err_code) // assert no error returned
}

func TestBinaryIncTM(t *testing.T) {
	// the binary number we wish to increment is 011, which is expected to become 100
	err, tm := examples.IncBinaryTM([]string{"0", "1", "1"})
	assert.Nil(t, err)
	err = tm.Run(nil, nil)
	assert.Nil(t, err)
	// note that tm.Tape[0] is "reserved", so the binary number starts at tm.Tape[1]
	assert.Equal(t, tm.AlphabetMap["1"], tm.Tape[1])
	assert.Equal(t, tm.AlphabetMap["0"], tm.Tape[2])
	assert.Equal(t, tm.AlphabetMap["0"], tm.Tape[3])
}

func TestInfiniteBinaryIncTM(t *testing.T) {
	//TODO: implement this test
}
