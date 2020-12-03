package run

var Zilch = Funct(zilch{})

type zilch struct{}

func (z zilch) String() string {
	return "zilch"
}

func (z zilch) Text() Text {
	return Text("")
}

func (z zilch) Call(_ *Under, _ Queue) (Funct, error) {
	return z, nil
}
