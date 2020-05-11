package wrap

import (
	"fmt"

	"github.com/superloach/defunct/types"
)

type Args []*Wrap

func (a *Args) String() string {
	return fmt.Sprintf("%s", *a)
}

type Wrap struct {
	Funct  types.Funct
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

func (w *Wrap) Run(under types.Funct) (types.Funct, error) {
	if w == nil {
		return types.Zilch, nil
	}

	if _, ok := w.Funct.(*types.Under); ok {
		w.Funct = under
	}

	w.Debugf("run %s\n", w)

	/*	if w.Argss == nil {
			w.Debugf("nil args -> %s\n", w.Funct)
			return w.Funct, nil
		}

		if len(w.Argss) == 0 {
			w.Debugf("zero args -> %s\n", w.Funct)
			return w.Funct, nil
		}
	*/
	if w.Funct == nil {
		w.Debugf("nil funct -> Zilch\n")
		return types.Zilch, nil
	}

	var funct types.Funct = w.Funct
	for _, cargs := range w.Argss {
		args := make([]types.Funct, len(*cargs))
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
			funct = types.Zilch
		case types.Error:
			return funct, funct.(types.Error)
		default:
		}
	}

	return funct, nil
}
