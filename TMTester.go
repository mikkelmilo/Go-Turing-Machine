package main

import (
	"Turingmachine/TM"
	"fmt"
)

func main() {
	tm := incBinaryTM("0")
	err := tm.Run()
	if err == nil {
		fmt.Println("Succesfully ran TM")
	} else {
		fmt.Println(err)
	}

	//tm1 := incBinaryTM("111")
	//fmt.Println(tm1.Run())
	/*
		tm := TM.NewTM("010")
		s1 := TM.State{"r"}
		s2 := TM.State{"b"}
		s3 := TM.State{"c"}
		s4 := TM.State{"d"}
		s5 := TM.State{"e"}
		ha := TM.State{"ha"}

		tm.StartState = &s1
		tm.AddTransition(&s1, &s1, "1", "1", ">")
		tm.AddTransition(&s1, &s1, "0", "0", ">")
		tm.AddTransition(&s1, &s2, "_", "_", "<")
		tm.AddTransition(&s2, &s2, "1", "0", "<")
		tm.AddTransition(&s2, &s3, "_", "_", ">")
		tm.AddTransition(&s2, &s5, "0", "1", "<")
		tm.AddTransition(&s3, &s4, "0", "1", ">")
		tm.AddTransition(&s4, &s4, "0", "0", ">")
		tm.AddTransition(&s4, &s4, "1", "1", ">")
		tm.AddTransition(&s4, &s5, "_", "0", "<")
		tm.AddTransition(&s5, &s5, "1", "1", "<")
		tm.AddTransition(&s5, &s5, "0", "0", "<")
		tm.AddTransition(&s5, &s1, "_", "_", ">")
		tm.AcceptState = &ha

		fmt.Println(tm.Run())*/
}

// doubles the size of the tape and places "_" symbols on the new slots
func expandTape(s []string) []string {
	a := make([]string, len(s))
	for i := range a {
		a[i] = "_"
	}
	return append(s, a...)
}

func incBinaryTM(s string) *TM.TM {
	tm := TM.NewTM(s)
	s1 := TM.State{"r"}
	s2 := TM.State{"b"}
	s3 := TM.State{"c"}
	s4 := TM.State{"d"}
	s5 := TM.State{"e"}
	ha := TM.State{"ha"}

	tm.StartState = &s1
	tm.AddTransition(&s1, &s1, "1", "1", ">")
	tm.AddTransition(&s1, &s1, "0", "0", ">")
	tm.AddTransition(&s1, &s2, "_", "_", "<")
	tm.AddTransition(&s2, &s2, "1", "0", "<")
	tm.AddTransition(&s2, &s3, "_", "_", ">")
	tm.AddTransition(&s2, &s5, "0", "1", "<")
	tm.AddTransition(&s3, &s4, "0", "1", ">")
	tm.AddTransition(&s4, &s4, "0", "0", ">")
	tm.AddTransition(&s4, &s4, "1", "1", ">")
	tm.AddTransition(&s4, &s5, "_", "0", "<")
	tm.AddTransition(&s5, &s5, "1", "1", "<")
	tm.AddTransition(&s5, &s5, "0", "0", "<")
	tm.AddTransition(&s5, &s1, "_", "_", ">")

	tm.AcceptState = &ha
	return &tm
}

func printTM(tm TM.TM) {
	fmt.Println("Tape:")
	a := tm.Tape[tm.Head]
	tm.Tape[tm.Head] = "{" + a + "}"
	fmt.Println(tm.Tape)
	tm.Tape[tm.Head] = a
}
