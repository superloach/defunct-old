package defunct

import (
	"fmt"

	"github.com/superloach/defunct/functs"
)

type Args []*Wrap

func (a *Args) String() string {
	if *a == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%s", *a)
}

type Wrap struct {
	Arena  *Arena
	Funct  functs.Funct
	Argss  []*Args
	Parent *Wrap
}

func (w *Wrap) Args() *Args {
	if w == nil {
		return nil
	}

	if len(w.Argss) == 0 {
		return nil
	}

	return w.Argss[len(w.Argss)-1]
}

func (w *Wrap) String() string {
	return fmt.Sprintf("%s%s", w.Funct, w.Argss)
}

func (w *Wrap) Run() (functs.Funct, error) {
	if w == nil {
		return functs.Zilch, nil
	}

	w.Arena.Debugf("run %s\n", w)

	if w.Argss == nil {
		w.Arena.Debugf("nil args -> %s\n", w.Funct)
		return w.Funct, nil
	}

	if len(w.Argss) == 0 {
		w.Arena.Debugf("zero args -> %s\n", w.Funct)
		return w.Funct, nil
	}

	if w.Funct == nil {
		w.Arena.Debugf("nil funct -> Zilch\n")
		return functs.Zilch, nil
	}

	funct := w.Funct
	for _, cargs := range w.Argss {
		args := make([]functs.Funct, len(*cargs))
		for i, carg := range *cargs {
			arg, err := carg.Run()
			if err != nil {
				return nil, err
			}
			args[i] = arg
		}

		w.Arena.Debugf("%s %s\n", funct, args)
		funct = funct.Call(args)

		w.Arena.Debugf(" -> %s\n", funct)

		switch funct.(type) {
		case nil:
			funct = functs.Zilch
		case functs.Error:
			return funct, funct.(functs.Error)
		default:
		}
	}

	return funct, nil
}
