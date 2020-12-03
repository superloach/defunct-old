package std

import "github.com/superloach/defunct/run"

type Repr struct{}

func (r Repr) String() string {
	return "repr"
}

func (r Repr) Call(_ *run.Under, in run.Queue) (run.Funct, error) {
	if len(in) != 1 {
		return nil, run.PraxisOf(r)
	}

	return run.Text(in[0].String()), nil
}

func (r Repr) Praxis() []string {
	return []string{
		`[x] -> repr of x`,
	}
}
