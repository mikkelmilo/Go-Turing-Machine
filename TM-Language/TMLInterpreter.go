package TML

import (
	"github.com/mikkelmilo/Go-Turing-Machine/TM"
	"strings"
)

//Interpret interprets the input
func Interpret(file string) (error, *TM.TM) {
	output, err := CheckSyntax(file)
	if err != nil {
		return err, nil
	}
	formatOut := [][]string{}
	for i := range output {
		// fmt.Println(output[i], len(output[i])-1)
		s1 := output[i][1 : len(output[i])-1]
		formatOut = append(formatOut, strings.SplitN(s1, ",", 5))
	}
	nameList := []string{}
	stateMap := make(map[string]*TM.State)
	err, tm := TM.NewTM([]string{"0", "1"}, nil)
	if err != nil {
		return err, nil
	}
	//add all states to map and create transitions
	for _, tuple := range formatOut {
		if contains(nameList, tuple[0]) == false {
			nameList = append(nameList, tuple[0])
			createAndAddState(tuple[0], nameList, stateMap, &tm)
		}
		if contains(nameList, tuple[1]) == false {
			nameList = append(nameList, tuple[1])
			createAndAddState(tuple[1], nameList, stateMap, &tm)
		}
		err := tm.AddTransition(stateMap[tuple[0]], stateMap[tuple[1]], tuple[2], tuple[3], tuple[4])
		if err != nil {
			return err, nil
		}
	}
	//fmt.Println("nameList: ", nameList)
	//fmt.Println("stateMap: ", stateMap)

	return nil, &tm
}

func createAndAddState(name string, nameList []string, stateMap map[string]*TM.State, tm *TM.TM) {
	state := TM.State{Name: name}
	if name == "ha" {
		tm.AcceptState = &state
	} else if name == "hr" {
		tm.RejectState = &state
	} else if name == "hs" {
		tm.StartState = &state
	}
	stateMap[name] = &state
}

/*
* for each tuple a:
* 	create new state with name a[0] if it doesn't already exist
* -||- for a[1]
* make a transition from a[0] to [1] with symbol [2] and replace by a[3] and move [4]
 */

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
