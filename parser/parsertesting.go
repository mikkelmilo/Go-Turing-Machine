package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"fmt"
	"github.com/mikkelmilo/Go-Turing-Machine/parser/TM-Language"
)

func main() {
	// Setup the input
	is := antlr.NewInputStream("(hs,a,_,1,>)(a,b,_,0,>)(b,ha,_,1,_)")

	// Create the Lexer
	lexer := parser.NewTMLLexer(is)

	// Read all tokens
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}
