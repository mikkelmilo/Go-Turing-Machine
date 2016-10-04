/*
 * BNF for TML
 * example command: (s1, s2, 0, 1, >)
 * in text: in state s1 with symbol 0, replace with 1, move right, and go to state s2
 * <Program> ::= {<Command>}*
 * <Command> ::=  "(" <StateLabel>","<StateLabel>","<Symbol>","<Symbol>","<Direction>");"
 * <StateLabel> ::= <AcceptState> | <RejectState> | <StartState> <variableName>
 * <variableName ::= (<Letter>| "_"){<Letter> | <decimal digit> | "_"}*
 * <StartState> ::= "hs" ok
 * <AcceptState> ::= "ha" ok
 * <RejectState> ::= "hr" ok
 * <Letter> ::= "A" | "B" | Z | ... | "a" | "b" | ...   ok
 * <decimal digit> ::= <binary digit> | "2" | "3" | ... | 9   ok
 * <binary digit> ::= "0" | "1" ok
 * <Symbol> ::= <binary digit> | "_" ok
 * <Direction> ::= "<" | ">" | "_" ok
 * ------------------------------------
 * BNF for ATML (Name subject to change)
 * <Program> ::= {<Command>|<MacroDefinition>|<MacroApplication>}*
 * <MacroDefinition> ::= "define" <MacroLabel> "{"
 *							{<Command>|<MacroApplication>}*
 *						"}"
 * <MacroLabel> ::= "Macro("<VariableName>")"
 * <MacroApplication> ::= "(" <StateLabel> "," <Symbol> ")" <MacroLabel> "(" <StateLabel> "," <StateLabel> ")"
 *
 *
 */

package TML

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

// Interval representation
type Interval struct {
	low  byte
	high byte
}

var alphabet = [...]Interval{Interval{low: 'a', high: 'z'}, Interval{low: 'A', high: 'Z'}}
var startedMacroDef bool

//0-9 = 48-57
//A-Z = 65-90
//a-z = 97-122

//CheckSyntax checks syntax..
func CheckSyntax(fileLocation string) ([]string, error) {
	formatted, err := formatInput(readFile(fileLocation))
	fmt.Print("formatted input: ")
	fmt.Printf("%d %q\n", len(formatted), formatted)
	//str := SpaceMap(readFile("test.txt"))
	//a := strings.Count(str, ";")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(checkProgram(formatted))
	return formatted, nil
}

func inAlphabet(a byte) bool {
	for _, interval := range alphabet {
		if a >= interval.low && a <= interval.high {
			return true
		}
	}
	return false
}

func readFile(path string) string {
	dat, err := ioutil.ReadFile(path)
	check(err)
	return string(dat)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// formats input file and returns a slice of commands, represented as strings
// correctness of syntax not guaranteed. if input is "a;;b;c;;" it will return [a; ; b; c; ;]
func formatInput(input string) ([]string, error) {
	// insufficient: missing tabs and possibly more
	t := strings.Replace(input, "\r", "", -1)
	t = strings.Replace(t, "\t", "", -1)
	t = strings.Replace(t, " ", "", -1)

	/*prevSemicolon := -1
	occur := 0
	for i, s := range t {
		if s == ';' {
			result[occur] = t[prevSemicolon+1 : i+1]
			prevSemicolon = i
			occur++
		}
	}*/
	str := ""
	for i := 0; i < len(t); i++ {
		if t[i] != '\n' {
			str += t[i : i+1]
		} else if i > 0 && t[i] == '\n' && t[i-1] != '\n' {
			str += t[i : i+1]
		}
	}
	if str[len(str)-1] == '\n' {
		str = str[0 : len(str)-1]
	}
	count := strings.Count(str, "\n")
	result := strings.SplitN(str, "\n", count+1)
	return result, nil
}

// SpaceMap asd
func SpaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func checkSyntax(input []string) (bool, error) {

	return false, nil
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
	a := []byte(s)
	if len(a) == 1 {
		return inAlphabet(a[0])
	}
	return false
}

func isBinaryDigit(s string) bool {
	return s == "0" || s == "1"
}

func isDecimalDigit(s string) bool {
	i, err := strconv.Atoi(s)
	if err == nil {
		return i >= 0 && i <= 9
	}
	return false
}

func isSymbol(s string) bool {
	return isBinaryDigit(s) || s == "_"
}

func isDirection(s string) bool {
	return s == "<" || s == ">" || s == "_"
}

// checks for productions

func checkStateLabel(s string) bool {
	if isAccept(s) {
		return true
	} else if isReject(s) {
		return true
	} else if isStart(s) {
		return true
	} else {
		return checkVariableName(s)
	}
}

func checkVariableName(s string) bool {
	if isLetter(s[0:1]) || strings.HasPrefix(s, "_") {
		if len(s) == 1 {
			return true
		}
		for i, c := range s {
			if i > 0 {
				//convert c from a rune to a string
				var a string
				a = fmt.Sprintf("%c", c)
				if isLetter(a) || isDecimalDigit(a) || a == "_" {
					return true
				}
			}
		}
	}
	return false
}

func checkCommand(s string) bool {
	hasBoundary := strings.HasPrefix(s, "(") && strings.HasSuffix(s, ")")
	if hasBoundary && strings.Count(s, ",") == 4 {
		s1 := s[1 : len(s)-1]
		a := strings.SplitN(s1, ",", 5)
		for _, x := range a {
			if len(x) == 0 {
				return false
			}
		}
		if checkStateLabel(a[0]) {
			if checkStateLabel(a[1]) {
				if isSymbol(a[2]) {
					if isSymbol(a[3]) {
						if isDirection(a[4]) {
							return true
						} else {
							fmt.Println("failed direction of ", a[4])
						}
					} else {
						fmt.Println("failed symbol of ", a[3])
					}
				} else {
					fmt.Println("failed symbol of ", a[2])
				}
			} else {
				fmt.Println("failed label of ", a[1])
			}
		} else {
			fmt.Println("failed label of", a[0])
		}
	}
	return false
}

func checkProgram(s []string) bool {
	for _, a := range s {
		if checkCommand(a) == false {
			fmt.Println("incorrect command:", a)
			return false
		}
	}
	return true
}

//ATML part

func checkMacroLabel(s string) bool {
	//s does not include the "{"
	if String.HasPefix(s, "Macro(") && String.HasSuffix(s, ")") {
		return checkVariableName(s[6 : len(s)-1])
	}
	return false
}

func checkMacroDefinition(s string) {
	if String.HasPrefix(s, "define ") {
		s1 := SpaceMap(s[7:len(s)])
		length := len(s1)
		if s1[length-1:length] == "}" && checkMacroLabel(s1[0:length]) {

		}
	}
	return false
}
