// Defunct.g4
grammar Defunct;

// Tokens
UNDER: '_';
OPEN: '[';
CLOSE: ']';
STR: ('~'. | ~[_[\] \r\n\t])+;
NEWLINE: [\r\n]+;
WHITESPACE: [ \t]+ -> skip;

// Rules
start: EOF|(wrap EOF);
funct: UNDER|STR;
args: OPEN wrap* CLOSE;
wrap: funct args*;
