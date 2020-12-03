package parse

const (
	RuneUnder  = '_'
	RuneOpen   = '['
	RuneClose  = ']'
	RuneStart  = '{'
	RuneEnd    = '}'
	RuneEscape = '~'
	RuneSpace  = ' '
	RuneTab    = '\t'
	RuneLine   = '\n'
)

const (
	StringSpace = string(RuneSpace) + string(RuneTab) + string(RuneLine)

	StringEnder = StringSpace + string(RuneUnder) + string(RuneOpen) +
		string(RuneClose) + string(RuneStart) + string(RuneEnd)

	StringEscape = StringEnder + string(RuneEscape)
)
