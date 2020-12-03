```go
RuneUnder = '_'
RuneOpen = '['
RuneClose = ']'
RuneStart = '{'
RuneEnd = '}'
RuneEscape = '~'

StringSpace = " \t\n"

StringEnder =
	StringSpace
	RuneUnder
	RuneOpen
	RuneClose
	RuneStart
	RuneEnd

StringEscape =
	StringEnder
	RuneEscape

Module = []Expr

Expr = Literal []Arguments
	Literal = RuneUnder | Text | Block
		Text = [](Escape | !StringEnder)
			Escape = RuneEscape StringEscape

		Block = Section(RuneStart, RuneEnd)

	Arguments = Section(RuneOpen, RuneClose)

Section(Start, End) = Start []Expr End
```
