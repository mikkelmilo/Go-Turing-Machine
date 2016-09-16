/*
 * BNF for TML
 * example command: (s1, s2, 0, 1, >);
 * in text: in state s1 with symbol 0, replace with 1, move right, and go to state s2
 * <Program> ::= {<Command>}*
 * <Command> ::=  "(" <StateLabel>, <StateLabel>, <Symbol>, <Symbol>, <Direction> ");"
 * <StateLabel> ::= <AcceptState> | <RejectState> | <StartState> | (<Letter>| "_"){<Letter> | <decimal digit> | "_"}*
 * <StartState> ::= "hs"
 * <AcceptState> ::= "ha"
 * <RejectState> ::= "hr"
 * <Letter> ::= "A" | "B" | Z | ... | "a" | "b" | ...
 * <decimal digit> ::= <binary digit> | "2" | "3" | ... | 9
 * <binary digit> ::= "0" | "1"
 * <Symbol> ::= <binary digit> | "_"
 * <Direction> ::= "<" | ">" | "_"
 */

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

alphabet := [...]string{"A", "B", "C", "D"}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(path string) string {
	dat, err := ioutil.ReadFile(path)
	check(err)
	return string(dat)
}

// formats input file and returns a slice of commands, represented as strings
// correctness of syntax not guaranteed. if input is "a;;b;c;;" it will return [a; ; b; c; ;]
func formatInput(input string) []string {
	t := strings.Replace(input, "\n", "", -1)
	t = strings.Replace(t, "\r", "", -1)
	t = strings.Replace(t, "\t", "", -1)

	c := strings.Count(t, ";")

	result := make([]string, c)
	prevSemicolon := -1
	occur := 0
	for i, s := range t {
		if s == ';' {
			result[occur] = t[prevSemicolon+1 : i+1]
			prevSemicolon = i
			occur++
		}
	}
	return result
}

func main() {
	formatted := formatInput(readFile("test.txt"))
	fmt.Println("formatted input: \n", formatted)
    //res, err := checkSyntax(formatted)
    fmt.Println("A"=="\u0041")
}

func checkSyntax(input []string) bool, error {


}

// predicates

func isStart(s string) bool {
	return s == "hs"
}

func isReject(s string) bool {
	return s == "hr"
}

func isAccept(s string) bool {
	return s == "ha"
}

func isLetter(s string) bool {

}
