package parse

type Expr struct {
	Value *Value  `parser:"@@"`
	Calls []*Call `parser:"@@*"`
}

func (e *Expr) String() string {
	s := e.Value.String()

	for _, call := range e.Calls {
		s += call.String()
	}

	return s
}
