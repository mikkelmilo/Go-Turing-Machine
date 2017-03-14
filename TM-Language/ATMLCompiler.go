package TML

import (
	"strconv"
	"strings"
)

var replaceName string
var namenr int

//Compile syntaxchecks and generates a new TML file
func Compile(file string) ([]string, error) {
	replaceName = "0var"
	namenr = 0
	res, err := CheckSyntax(file)
	if err != nil {
		return nil, err
	}
	MacroDefMap := make(map[string][]string)
	cmdList := []string{}
	inMacro := ""
	// sets u the cmdList and MacroDefMap
	for _, line := range res {
		if CheckMacroDef(line) {
			inMacro = line[6 : len(line)-1]
			MacroDefMap[inMacro] = []string{}
		} else if inMacro != "" && (CheckCommand(line) || CheckMacroApp(line)) {
			MacroDefMap[inMacro] = append(MacroDefMap[inMacro], line)
		} else if line == "}" {
			inMacro = ""
		} else {
			cmdList = append(cmdList, line)
		}
	}

	//folds out macro-applications and replaces their local names with fresh names
	i := 0
	for {
		if i == len(cmdList) {
			break
		}
		if CheckMacroApp(cmdList[i]) {
			cmdList = foldOut(i, cmdList, MacroDefMap)
		}
		i++
	}
	return cmdList, nil

}

func foldOut(index int, cmds []string, macros map[string][]string) []string {
	s := getName(index, cmds)
	//replace names
	for _, line := range macros["Macro("+s+")"] {
		if CheckCommand(line) {
			//seperate by comma
		}
	}
	afterindex := cmds[index+1:]

	result := append(cmds[:index], append(macros["Macro("+s+")"], afterindex...)...)
	return result
}

func getName(index int, cmds []string) string {
	//(...)Macro(name)(...)
	a := strings.Index(cmds[index], "M")
	s := cmds[index][a:]
	// S = Macro(name)(...)
	s = s[6:]
	// S = name)(...)
	s = s[:strings.Index(s, ")")]
	// S = name
	return s
}

func nextName() string {
	namenr++
	return replaceName + strconv.Itoa(namenr)
}
