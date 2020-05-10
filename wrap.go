package defunct

import (
	"fmt"

	"github.com/superloach/defunct/debug"
	"github.com/superloach/defunct/functs"
)

type Args []*Wrap

func (a *Args) String() string {
	return fmt.Sprint(*a)
}

type Wrap struct {
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
	debug.Debugf("run %s\n", w)

	if w == nil {
		debug.Debugf("nil wrap -> Zilch\n")
		return functs.Zilch, nil
	}

	if w.Argss == nil {
		debug.Debugf("nil args -> %s\n", w.Funct)
		return w.Funct, nil
	}

	if len(w.Argss) == 0 {
		debug.Debugf("zero args -> %s\n", w.Funct)
		return w.Funct, nil
	}

	if w.Funct == nil {
		debug.Debugf("nil funct -> Zilch\n")
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

		debug.Debugf("%s %s\n", funct, args)
		funct = funct.Call(args)

		debug.Debugf(" -> %s\n", funct)

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
