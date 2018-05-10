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

func TMToDotFile(tm TM.TM) error {
	transitions := tm.Transitions
	var graphBuf bytes.Buffer
	transitions_string := make([]string, 0)
	hasSelfLoop := make(map[string]bool) // maps a state name to true if it has a self loop in the TM
	graphBuf.WriteString("digraph TM {\n")
	graphBuf.WriteString("node [nodesep=2.0, fontsize=11];\n")
	graphBuf.WriteString("graph [overlap = false];\n")
	for _, t := range transitions {
		// if this transition is a self loop, and it is the first self loop seen for this state
		// then draw a new transition. Otherwise just add text to the label.
		ok, val := hasSelfLoop[t.CurState.String()]
		if t.CurState.String() == t.NewState.String() && (!ok || !val) {
			hasSelfLoop[t.CurState.String()] = true

			label := "label=\"" + t.GetCurSymbol(tm) + "," + t.GetNewSymbol(tm) + "," + t.GetDir() + "\""
			minlength := "minlength=2"
			options := " [" + label + " " + minlength + " ]"
			transitions_string = append(transitions_string, t.CurState.String()+" -> "+t.NewState.String()+options+";\n")
		} else if t.CurState.String() == t.NewState.String() && hasSelfLoop[t.CurState.String()] == true {
			// else find the graphviz transition and modify it
			str := t.CurState.String() + " -> " + t.NewState.String()
			for j, tr := range transitions_string {
				if strings.HasPrefix(tr, str) {
					//print("replaced ", tr, " with ")
					new_str := "label=\"" + t.GetCurSymbol(tm) + "," + t.GetNewSymbol(tm) + "," + t.GetDir() + "\\n"
					transitions_string[j] = strings.Replace(tr, "label=\"", new_str, 1)
					//println(transitions_string[j])
					break
				}
			}
		} else {
			label := "label=\"" + t.GetCurSymbol(tm) + "," + t.GetNewSymbol(tm) + "," + t.GetDir() + "\""
			minlength := "minlength=2"
			options := " [" + label + " " + minlength + " ]"
			transitions_string = append(transitions_string, t.CurState.String()+" -> "+t.NewState.String()+options+";\n")
		}
	}
	for _, tr := range transitions_string {
		graphBuf.WriteString(tr)
	}
	graphBuf.WriteString("}")
	println(graphBuf.String())
	graph, err := gographviz.Parse(graphBuf.Bytes())
	if err != nil {
		return err
	}

	f, err := os.Create("tm.dot")
	defer f.Close()
	if err != nil {
		return err
	}
	_, err = f.WriteString(graph.String())
	if err != nil {
		return err
	}
	f.Sync()

	return nil
}
