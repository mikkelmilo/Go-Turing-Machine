// TML.g4
grammar TML;

// Tokens
DEFINE: 'define';
MACRO: 'macro';
STARTSTATE: 'hs';
ACCEPTSTATE: 'ha';
REJECTSTATE: 'hr';
UNDERSCORE: '_';
LPAREN: '(';
RPAREN: ')';
LBRACE: '{';
RBRACE:	'}';
COMMA: ',';
LT: '<';
GT: '>';
ID: [a-zA-Z_]([a-zA-Z_] | [0-9])*;
DECIMAL: [0-9]+;
BlockComment
    :   '/*' .*? '*/'
        -> skip
    ;
LineComment
    :   '//' ~[\r\n]*
        -> skip
;
WHITESPACE: [ \r\n\t]+ -> skip;


// Rules
start : program EOF;

// earlier expressions have higher precedence.
// The part after # is the type name to be used in the generated go code
program
	:	statement+
	;

statement
	:	command
	|	macroDef
	|	macroApp
	;

macroApp
	:	LPAREN
			stateLabel COMMA
			tapeSymbol
		RPAREN
		ID
		LPAREN
			stateLabel COMMA
			stateLabel
		RPAREN
	;

macroDef
	:	DEFINE MACRO ID LBRACE statement* RBRACE
	;


command
	:	LPAREN
			currentState=stateLabel COMMA
			newState=stateLabel COMMA
			currentSymbol=tapeSymbol COMMA
			newSymbol=tapeSymbol COMMA
			dir=direction
		RPAREN
	;

stateLabel
	:	ID
	|	STARTSTATE
	|	REJECTSTATE
	|	ACCEPTSTATE
	;

tapeSymbol
	:	ID
	|	DECIMAL
	|	UNDERSCORE
	;


direction
	:	LT
	|	GT
	|	UNDERSCORE
	;
