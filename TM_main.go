package main

import (
	"fmt"
	"github.com/mikkelmilo/Go-Turing-Machine/TM"
	"github.com/mikkelmilo/Go-Turing-Machine/TM-Language"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("No file specified. Exiting")
		os.Exit(1)
	}
	fileName := os.Args[1]
	quit := make(chan int)
	c := make(chan string)

	err, tm := TML.Interpret(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var err1 error
	go func() {
		err1 = tm.Run(c, quit)
	}()
	/*scanner := bufio.NewScanner(os.Stdin)
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
	}()*/
	//blocks until receive data from quit channel (which is sent from Run when it halts or finds an error)
	code := <-quit
	fmt.Println("Halted with error code: ", code)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(tm)
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

func incBinaryTM(s []string) *TM.TM {
	err, tm := TM.NewTM([]string{"0", "1"}, s)
	if err != nil {
		print(err)
	}
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
