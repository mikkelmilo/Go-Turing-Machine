package Tools

import (
	"bytes"
	"github.com/awalterschulze/gographviz"
	"github.com/mikkelmilo/Go-Turing-Machine/TM"
	"os"
	"strings"
)

/*
 This file contains a tool for exporting a Turing Machine into a DOT language file - a graph description language.
 This way, we can use external tools, such as graphviz, to visualize Turing Machines.
*/

func TMToDotFile(tm TM.TM, fileName string) error {
	transitions := tm.GetTransitions()
	var graphBuf bytes.Buffer
	transitionsString := make([]string, 0)
	hasSelfLoop := make(map[string]bool) // maps a state name to true if it has a self loop in the tmImpl

	graphBuf.WriteString("digraph tmImpl {\n")
	graphBuf.WriteString("node [nodesep=2.0, fontsize=11];\n")
	graphBuf.WriteString("graph [overlap = false];\n")

	transitionsString = makeTransitions(transitions, hasSelfLoop, tm, transitionsString)

	for _, tr := range transitionsString {
		graphBuf.WriteString(tr)
	}

	graphBuf.WriteString("}")
	graph, err := gographviz.Parse(graphBuf.Bytes())
	if err != nil {
		return err
	}

	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		return err
	}
	if _, err := f.WriteString(graph.String()); err != nil {
		return err
	}
	f.Sync()
	return nil
}

func makeTransitions(transitions []TM.Transition, hasSelfLoop map[string]bool, tm TM.TM, transitionsString []string) []string {
	for _, t := range transitions {
		// if this transition is a self loop, and it is the first self loop seen for this state
		// then draw a new transition. Otherwise just add text to the label.
		ok, val := hasSelfLoop[t.CurState.String()]
		if t.CurState.String() == t.NewState.String() && (!ok || !val) {
			hasSelfLoop[t.CurState.String()] = true
			transitionsString = addNewTransitionString(t, tm, transitionsString)
		} else if t.CurState.String() == t.NewState.String() && hasSelfLoop[t.CurState.String()] == true {
			// else find the graphviz transition and modify it
			str := t.CurState.String() + " -> " + t.NewState.String()
			findAndModifyTransition(transitionsString, str, t, tm)
		} else {
			transitionsString = addNewTransitionString(t, tm, transitionsString)
		}
	}
	return transitionsString
}

func addNewTransitionString(t TM.Transition, tm TM.TM, transitionsString []string) []string {
	label := "label=\" " + t.GetCurSymbol(tm) + "," + t.GetNewSymbol(tm) + "," + t.GetDir() + "\""
	minLength := "minLength=2"
	options := " [" + label + " " + minLength + " ]"
	transitionsString = append(transitionsString, t.CurState.String()+" -> "+t.NewState.String()+options+";\n")
	return transitionsString
}

func findAndModifyTransition(transitionsString []string, str string, t TM.Transition, tm TM.TM) {
	for j, tr := range transitionsString {
		if strings.HasPrefix(tr, str) {
			newStr := "label=\" " + t.GetCurSymbol(tm) + "," + t.GetNewSymbol(tm) + "," + t.GetDir() + "\\n"
			transitionsString[j] = strings.Replace(tr, "label=\" ", newStr, 1)
			break
		}
	}
}
