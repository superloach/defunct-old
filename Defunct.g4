// Defunct.g4
grammar Defunct;

// Tokens
UNDER: '_';
OPEN: '[';
CLOSE: ']';
ESC: '~';
WS: [ \t\n];
ESCD: ESC ('_' | '[' | ']' | '~' | [ \t\n]);
CHAR: ~ ('_' | '[' | ']' | '~' | [ \t\n]);
STR: (ESCD | CHAR)+;

// Rules
start: (line WS*)* EOF;
line: wrap;
wrap: funct args*;
funct: UNDER|STR;
args: OPEN (wrap (WS+ wrap)*)? CLOSE;
