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
start : primaryExpression EOF;

// earlier expressions have higher precedence.
// The part after # is the type name to be used in the generated go code
primaryExpression
	:	subExpression+
	;

subExpression
	:	command
	|	macroDef
	|	macroApp
	;

macroApp
	:	LPAREN
			stateLabel COMMA
			tapeSymbol
		RPAREN
		macroLabel
		LPAREN
			stateLabel COMMA
			stateLabel COMMA
		RPAREN
	;

macroDef
	:	DEFINE macroLabel LBRACE subExpression* RBRACE
	;

macroLabel
	:	MACRO LPAREN ID RPAREN
	;

command
	:	LPAREN
			stateLabel COMMA
			stateLabel COMMA
			tapeSymbol COMMA
			tapeSymbol COMMA
			direction
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
