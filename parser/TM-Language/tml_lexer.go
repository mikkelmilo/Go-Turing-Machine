// Generated from TM-Language/TML.g4 by ANTLR 4.7.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 20, 121,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 7, 15, 80,
	10, 15, 12, 15, 14, 15, 83, 11, 15, 3, 16, 6, 16, 86, 10, 16, 13, 16, 14,
	16, 87, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17, 94, 10, 17, 12, 17, 14, 17,
	97, 11, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3,
	18, 7, 18, 108, 10, 18, 12, 18, 14, 18, 111, 11, 18, 3, 18, 3, 18, 3, 19,
	6, 19, 116, 10, 19, 13, 19, 14, 19, 117, 3, 19, 3, 19, 3, 95, 2, 20, 3,
	3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13,
	25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 3, 2, 7, 5, 2,
	67, 92, 97, 97, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 3, 2, 50,
	59, 4, 2, 12, 12, 15, 15, 5, 2, 11, 12, 15, 15, 34, 34, 2, 125, 2, 3, 3,
	2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3,
	2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19,
	3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2,
	27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2,
	2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 3, 39, 3, 2, 2, 2, 5, 46, 3, 2, 2,
	2, 7, 52, 3, 2, 2, 2, 9, 55, 3, 2, 2, 2, 11, 58, 3, 2, 2, 2, 13, 61, 3,
	2, 2, 2, 15, 63, 3, 2, 2, 2, 17, 65, 3, 2, 2, 2, 19, 67, 3, 2, 2, 2, 21,
	69, 3, 2, 2, 2, 23, 71, 3, 2, 2, 2, 25, 73, 3, 2, 2, 2, 27, 75, 3, 2, 2,
	2, 29, 77, 3, 2, 2, 2, 31, 85, 3, 2, 2, 2, 33, 89, 3, 2, 2, 2, 35, 103,
	3, 2, 2, 2, 37, 115, 3, 2, 2, 2, 39, 40, 7, 102, 2, 2, 40, 41, 7, 103,
	2, 2, 41, 42, 7, 104, 2, 2, 42, 43, 7, 107, 2, 2, 43, 44, 7, 112, 2, 2,
	44, 45, 7, 103, 2, 2, 45, 4, 3, 2, 2, 2, 46, 47, 7, 111, 2, 2, 47, 48,
	7, 99, 2, 2, 48, 49, 7, 101, 2, 2, 49, 50, 7, 116, 2, 2, 50, 51, 7, 113,
	2, 2, 51, 6, 3, 2, 2, 2, 52, 53, 7, 106, 2, 2, 53, 54, 7, 117, 2, 2, 54,
	8, 3, 2, 2, 2, 55, 56, 7, 106, 2, 2, 56, 57, 7, 99, 2, 2, 57, 10, 3, 2,
	2, 2, 58, 59, 7, 106, 2, 2, 59, 60, 7, 116, 2, 2, 60, 12, 3, 2, 2, 2, 61,
	62, 7, 97, 2, 2, 62, 14, 3, 2, 2, 2, 63, 64, 7, 42, 2, 2, 64, 16, 3, 2,
	2, 2, 65, 66, 7, 43, 2, 2, 66, 18, 3, 2, 2, 2, 67, 68, 7, 125, 2, 2, 68,
	20, 3, 2, 2, 2, 69, 70, 7, 127, 2, 2, 70, 22, 3, 2, 2, 2, 71, 72, 7, 46,
	2, 2, 72, 24, 3, 2, 2, 2, 73, 74, 7, 62, 2, 2, 74, 26, 3, 2, 2, 2, 75,
	76, 7, 64, 2, 2, 76, 28, 3, 2, 2, 2, 77, 81, 9, 2, 2, 2, 78, 80, 9, 3,
	2, 2, 79, 78, 3, 2, 2, 2, 80, 83, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 81, 82,
	3, 2, 2, 2, 82, 30, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 84, 86, 9, 4, 2, 2,
	85, 84, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 87, 88, 3,
	2, 2, 2, 88, 32, 3, 2, 2, 2, 89, 90, 7, 49, 2, 2, 90, 91, 7, 44, 2, 2,
	91, 95, 3, 2, 2, 2, 92, 94, 11, 2, 2, 2, 93, 92, 3, 2, 2, 2, 94, 97, 3,
	2, 2, 2, 95, 96, 3, 2, 2, 2, 95, 93, 3, 2, 2, 2, 96, 98, 3, 2, 2, 2, 97,
	95, 3, 2, 2, 2, 98, 99, 7, 44, 2, 2, 99, 100, 7, 49, 2, 2, 100, 101, 3,
	2, 2, 2, 101, 102, 8, 17, 2, 2, 102, 34, 3, 2, 2, 2, 103, 104, 7, 49, 2,
	2, 104, 105, 7, 49, 2, 2, 105, 109, 3, 2, 2, 2, 106, 108, 10, 5, 2, 2,
	107, 106, 3, 2, 2, 2, 108, 111, 3, 2, 2, 2, 109, 107, 3, 2, 2, 2, 109,
	110, 3, 2, 2, 2, 110, 112, 3, 2, 2, 2, 111, 109, 3, 2, 2, 2, 112, 113,
	8, 18, 2, 2, 113, 36, 3, 2, 2, 2, 114, 116, 9, 6, 2, 2, 115, 114, 3, 2,
	2, 2, 116, 117, 3, 2, 2, 2, 117, 115, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2,
	118, 119, 3, 2, 2, 2, 119, 120, 8, 19, 2, 2, 120, 38, 3, 2, 2, 2, 9, 2,
	79, 81, 87, 95, 109, 117, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'define'", "'macro'", "'hs'", "'ha'", "'hr'", "'_'", "'('", "')'",
	"'{'", "'}'", "','", "'<'", "'>'",
}

var lexerSymbolicNames = []string{
	"", "DEFINE", "MACRO", "STARTSTATE", "ACCEPTSTATE", "REJECTSTATE", "UNDERSCORE",
	"LPAREN", "RPAREN", "LBRACE", "RBRACE", "COMMA", "LT", "GT", "ID", "DECIMAL",
	"BlockComment", "LineComment", "WHITESPACE",
}

var lexerRuleNames = []string{
	"DEFINE", "MACRO", "STARTSTATE", "ACCEPTSTATE", "REJECTSTATE", "UNDERSCORE",
	"LPAREN", "RPAREN", "LBRACE", "RBRACE", "COMMA", "LT", "GT", "ID", "DECIMAL",
	"BlockComment", "LineComment", "WHITESPACE",
}

type TMLLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewTMLLexer(input antlr.CharStream) *TMLLexer {

	l := new(TMLLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "TML.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// TMLLexer tokens.
const (
	TMLLexerDEFINE       = 1
	TMLLexerMACRO        = 2
	TMLLexerSTARTSTATE   = 3
	TMLLexerACCEPTSTATE  = 4
	TMLLexerREJECTSTATE  = 5
	TMLLexerUNDERSCORE   = 6
	TMLLexerLPAREN       = 7
	TMLLexerRPAREN       = 8
	TMLLexerLBRACE       = 9
	TMLLexerRBRACE       = 10
	TMLLexerCOMMA        = 11
	TMLLexerLT           = 12
	TMLLexerGT           = 13
	TMLLexerID           = 14
	TMLLexerDECIMAL      = 15
	TMLLexerBlockComment = 16
	TMLLexerLineComment  = 17
	TMLLexerWHITESPACE   = 18
)
