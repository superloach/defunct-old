package std

import "github.com/superloach/defunct/run"

type Text struct{}

func (t Text) String() string {
	return "text"
}

func (t Text) Call(_ *run.Under, in run.Queue) (run.Funct, error) {
	if len(in) != 1 {
		return run.Text{}, nil
	}

	return run.Text(in[0].String()), nil
}

func (t Text) Praxis() []string {
	return []string{
		`[x] -> x as text`,
		`[] -> empty text`,
	}
}
