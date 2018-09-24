package examples

import "github.com/mikkelmilo/Go-Turing-Machine/TM"

/*
 * This packages contains example Turing-Machine programs written using the functions provided
 * by the TM package (ie. NOT by generating them using the compiler+interpreter).
 *
 * Each function returns a TM struct which will emulate the specified program when executed
 * using the .Run() method.
 */

/*
 * This function returns a TM which emulates the increment function on a binary number (left is most significant bit).
 * The binary number is pre-specified to be TODO: complete description
 *
 * The logic of this TM is as follows:
 * First move tape head to the end of the given binary number (by going right until a _ is met)
 *
 */
func IncBinaryTM(s []string) (error, *TM.TM) {
	err, tm := TM.NewTM([]string{"0", "1"}, s)
	if err != nil {
		return err, nil
	}
	start_state := TM.State{Name: "r"}
	s0 := TM.State{Name: "x"}
	s1 := TM.State{Name: "a"}
	s2 := TM.State{Name: "b"}
	s3 := TM.State{Name: "c"}
	s4 := TM.State{Name: "d"}
	s5 := TM.State{Name: "e"}
	ha := TM.State{Name: "ha"}
	hr := TM.State{Name: "hr"}
	tm.AcceptState = &ha
	tm.StartState = &start_state
	tm.RejectState = &hr
	//first go to the end of the binary number.

	err1 := tm.AddTransition(&start_state, &s0, "_", "_", ">")
	if err1 != nil {
		return err1, nil
	}

	// go to reject state if the first symbol is not "_"
	err1 = tm.AddTransition(&start_state, &hr, "1", "1", "_")
	if err1 != nil {
		return err1, nil
	}
	err1 = tm.AddTransition(&start_state, &hr, "0", "0", "_")
	if err1 != nil {
		return err1, nil
	}

	err1 = tm.AddTransition(&s0, &s1, "1", "1", ">")
	if err1 != nil {
		return err1, nil
	}
	err1 = tm.AddTransition(&s0, &s1, "0", "0", ">")
	if err1 != nil {
		return err1, nil
	}
	// if the first bit in the number is "_", then go to reject state since this is an error.
	err1 = tm.AddTransition(&s0, &hr, "_", "_", "<")
	if err1 != nil {
		return err1, nil
	}
	// keep going right
	err1 = tm.AddTransition(&s1, &s1, "1", "1", ">")
	if err1 != nil {
		return err1, nil
	}
	// keep going right
	err1 = tm.AddTransition(&s1, &s1, "0", "0", ">")
	if err1 != nil {
		return err1, nil
	}
	// here we arrive at the right-end of the number.
	err1 = tm.AddTransition(&s1, &s2, "_", "_", "<")
	if err1 != nil {
		return err1, nil
	}
	// on our way left-wards, replace occurrences of 1 with 0
	err1 = tm.AddTransition(&s2, &s2, "1", "0", "<")
	if err1 != nil {
		return err1, nil
	}
	// if we're at the left-most position, go back to the end (to the right) of the binary number
	err1 = tm.AddTransition(&s2, &s3, "_", "_", ">")
	if err1 != nil {
		return err1, nil
	}
	// if we see a 0 on our way left-wards, then replace with 1. Then we're done with this increment
	// because by construction we can assume all previously "visited" bits have been 1s
	// (which have already been replaced by 0s by construction of state s2)
	err1 = tm.AddTransition(&s2, &s5, "0", "1", "<")
	if err1 != nil {
		return err1, nil
	}
	err1 = tm.AddTransition(&s3, &s4, "0", "1", ">")
	if err1 != nil {
		return err1, nil
	}
	err1 = tm.AddTransition(&s4, &s4, "0", "0", ">")
	if err1 != nil {
		return err1, nil
	}
	err1 = tm.AddTransition(&s4, &s4, "1", "1", ">")
	if err1 != nil {
		return err1, nil
	}
	// add a new 0 at the end. This is the situation where 1111...111 is incremented to 1000...0000
	err1 = tm.AddTransition(&s4, &s5, "_", "0", "<")
	if err1 != nil {
		return err1, nil
	}

	//now move back to the "start" ie. all the way to the left.
	err1 = tm.AddTransition(&s5, &s5, "1", "1", "<")
	if err1 != nil {
		return err1, nil
	}
	err1 = tm.AddTransition(&s5, &s5, "0", "0", "<")
	if err1 != nil {
		return err1, nil
	}

	// when we're at the beginning, goto accept state and halt
	err1 = tm.AddTransition(&s5, &ha, "_", "_", "_")
	return err1, &tm
}

func InfinitelyIncBinaryTM(s []string) (error, *TM.TM) {
	err, tm := IncBinaryTM(s)
	if err != nil {
		return err, nil
	}

	// we exploit the fact that we know from the implementation of IncBinaryTM that the
	// last transition added is from s5 to ha, so we can get the last element in tm.Transitions
	// to get the pointer to s5. Then we delete the last transition, and add a new transition
	// which causes the TM to loop.
	s5 := tm.Transitions[len(tm.Transitions)-1].CurState
	tm.Transitions = tm.Transitions[:len(tm.Transitions)-1] // remove last item from tm.Transitions
	// add a transition where we start from the beginning whenever the number as been incremented once.
	err = tm.AddTransition(s5, tm.StartState, "_", "_", ">")
	return err, tm
}
