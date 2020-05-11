package wrap

import (
	"fmt"

	"github.com/superloach/defunct/functs"
)

type Args []*Wrap

func (a *Args) String() string {
	return fmt.Sprintf("%s", *a)
}

type Wrap struct {
	Funct  functs.Funct
	Argss  []*Args
	Parent *Wrap
	Debugf func(string, ...interface{})
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

func (w *Wrap) Run(under functs.Funct) (functs.Funct, error) {
	if w == nil {
		return functs.Zilch, nil
	}

	if _, ok := w.Funct.(*functs.Under); ok {
		w.Funct = under
	}

	w.Debugf("run %s\n", w)

	if w.Argss == nil {
		w.Debugf("nil args -> %s\n", w.Funct)
		return w.Funct, nil
	}

	if len(w.Argss) == 0 {
		w.Debugf("zero args -> %s\n", w.Funct)
		return w.Funct, nil
	}

	if w.Funct == nil {
		w.Debugf("nil funct -> Zilch\n")
		return functs.Zilch, nil
	}

	var funct functs.Funct = w.Funct
	for _, cargs := range w.Argss {
		args := make([]functs.Funct, len(*cargs))
		for i, carg := range *cargs {
			arg, err := carg.Run(under)
			if err != nil {
				return nil, err
			}
			args[i] = arg
		}

		w.Debugf("%s %s\n", funct, args)
		funct = funct.Call(under, args)

		w.Debugf(" -> %s\n", funct)

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
