package testing

import (
	"fmt"
	"github.com/mikkelmilo/Go-Turing-Machine/TM-Language"
	"log"
	"os"
	"testing"
)

func TestCompile(t *testing.T) {
	res, _ := TML.Compile("compileTest.txt")
	// open output file
	file, err := os.Create("compiledResult.txt")
	if err != nil {
		log.Fatal("Could create file", err)
	}
	defer file.Close()

	for _, line := range res {
		fmt.Fprintf(file, line+"\n")
	}

}
