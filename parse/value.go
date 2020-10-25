package parse

type Value struct {
	Under *string `parser:"  @Under"`
	Text  *Text   `parser:"| @Text"`
}

func (v *Value) String() string {
	switch {
	case v.Under != nil:
		return *v.Under
	case v.Text != nil:
		return v.Text.String()
	default:
		panic("unknown literal")
	}
}
