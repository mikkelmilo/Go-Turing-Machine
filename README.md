# Go-Turing-Machine

A Turing Machine implementation in Go as well as a declarative programming language (TML) for the Turing Machine.
The language supports macros, which are lightweight functions that take no arguments.

This project consists of the following components:
- TML: a declarative programming language for specifying Turing Machines, including support for macros
- ATML: anextension of TML which includes macros
- A Turing Machine representation, which may be executed.
- A compiler/interpreter that converts ATML/TML programs into a concrete TM instance in go  
- other features are on the way!

## !UPDATE!
I'm reviving this project! Expect lots of improvements, clean-ups, refactoring, and new stuff!
See the issues tab for my current TODO-list.

## An example
The following illustrates a program written in my Turing Machine Language (TML), and how to interpret this program using my TM implementation.

### TML program
The program below writes 1 on the initial tape position, then moves the tape head one spot to the right, writes 0, goes right again, and finally writes 1 and enters the accept state, thereby halting.
(hs,a,\_,1,>)
(a,b,\_,0,>)
(b,ha,\_,1,\_)

Save this in a file called 'test'.

```golang
func main() {
	// interprets the file and creates a TM ready to execute the program
	err, tm := TML.Interpret("test") 
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	go func() {
		tm.Run(c, quit) // run the TM in a separate goroutine
	}()
	// blocks until receive data from quit channel 
	// (which is sent from Run when it halts or finds an error)
	code := <-quit
	fmt.Println("Halted with error code: ", code)
	fmt.Println(tm) // prints a nicely formatted version of the TM
	return
}
```

This prints:
```
Halted with error code:  1
TM:
Alphabet: [0 1] 
Reject state: None
Current state: ha
Transitions: [(hs,a,_,1,1),(a,b,_,0,1),(b,ha,_,1,_)]
Tape:
[_ 1 0 { 1 }] 
```
The brackets around the 1 in the tape indicate that the tape head currently points at this element. 

If you want to see more involved programs, see the examples in examples/ExamplePrograms.go. There is for example a TM which emulates the increment function on binary numbers. Note that the TM "programs" in ExamplePrograms.go are not written in TML. They are implemented using the functions provided by the TM package; TM.AddTransition(...), etc. For a fun exercise, try to convert these programs into TML programs. It should be rather straight-forward.
