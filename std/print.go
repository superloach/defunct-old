package std

import (
	"github.com/pkg/errors"

	"github.com/superloach/defunct/run"
)

type Print struct{}

func (p Print) String() string {
	return "print"
}

func (p Print) Call(u *run.Under, in run.Queue) (run.Funct, error) {
	for i, f := range in {
		bs := []byte(string(run.TextOf(f)))

		if i+1 < len(in) {
			bs = append(bs, ' ')
		} else {
			bs = append(bs, '\n')
		}

		_, err := u.W.Write(bs)
		if err != nil {
			return nil, errors.Wrap(err, "write")
		}
	}

	return run.Zilch, nil
}

func (p Print) Praxis() []string {
	return []string{`[x...] -> prints x, returns Zilch`}
}
