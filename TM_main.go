package main

import (
	"TuringMachine/TM"
	"TuringMachine/TML"
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	quit := make(chan int)
	c := make(chan int)

	tm := TML.Interpret("tests/compiledResult.txt")
	go func() { tm.Run(c, quit) }()
	go func() {
		for scanner.Scan() {
			if scanner.Text() == "state" {
				c <- 0
			} else if scanner.Text() == "stop" {
				quit <- 0
			} else if scanner.Text() == "help" {
				fmt.Println("Keywords: ")
				fmt.Println("  states")
				fmt.Println("  stop")
				fmt.Println("Explanation of error codes:")
				fmt.Println("  0 means the TM has been manually stopped with no error.")
				fmt.Println("  1 means the TM has halted with no error.")
				fmt.Println(" -1 means the TM has halted with an error.")
			} else {
				fmt.Println("incorrect expression")
			}
		}
	}()
	//blocks until receive data from quit channel (which is sent from Run when it halts or finds an error)
	code := <-quit
	fmt.Println("Halted with error code: ", code)
	TM.PrintTM(tm)

	return
}

/*func scanForInput(quit, c chan int) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if scanner.Text() == "state" {
			c <- 0
		} else if scanner.Text() == "stop" {
			quit <- 0
			if err == nil {
				fmt.Println("Stopped TM.")
				TM.PrintTM(tm)
			}
		} else {
			fmt.Println("incorrect expression")
		}
	}
}*/

// doubles the size of the tape and places "_" symbols on the new slots
func expandTape(s []string) []string {
	a := make([]string, len(s))
	for i := range a {
		a[i] = "_"
	}
	return append(s, a...)
}

func incBinaryTM(s []uint8) *TM.TM {
	tm := TM.NewTM(s)
	s1 := TM.State{Name: "r"}
	s2 := TM.State{Name: "b"}
	s3 := TM.State{Name: "c"}
	s4 := TM.State{Name: "d"}
	s5 := TM.State{Name: "e"}
	ha := TM.State{Name: "ha"}

	tm.StartState = &s1
	tm.AddTransition(&s1, &s1, "0", "0", ">")
	tm.AddTransition(&s1, &s1, "1", "1", ">")
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
