package module

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/superloach/defunct/functs"
)

type Module struct {
	Name   string
	Files  functs.Thing
	Debugf func(string, ...interface{})
}

func LoadModule(
	mod string,
	debugf func(string, ...interface{}),
) (*Module, error) {
	if len(mod) == 0 {
		return nil, fmt.Errorf("zero len mod")
	}

	split := strings.Split(mod, "/")
	pmod := path.Join(split...)
	mname := split[len(split)-1]

	m := &Module{
		Name:   mname,
		Files:  make(functs.Thing),
		Debugf: debugf,
	}

	if mod[0] == '@' {
		panic("pkger stub")
	} else {
		dir, err := os.Open(pmod)
		if err != nil {
			return m, err
		}

		fnames, err := dir.Readdirnames(-1)

		for _, fname := range fnames {
			if !strings.HasSuffix(fname, ".df") {
				continue
			}

			name := strings.ToLower(fname[:len(fname)-3])
			fpath := path.Join(pmod, fname)

			f, err := os.Open(fpath)
			if err != nil {
				return m, err
			}

			code, err := ioutil.ReadAll(f)
			if err != nil {
				return m, err
			}

			m.Files[name] = LoadFile(m, fpath, name, string(code), debugf)
		}
	}

	return m, nil
}

func (m *Module) Run(in functs.Funct) (functs.Funct, error) {
	under := &functs.Under{
		Builtin: functs.Thing{
			"in": in,
		},
		Debugf: m.Debugf,
	}
	m.Debugf("under %s\n", under)

	res := m.Call(under, make([]functs.Funct, 0))
	out := under.GetProp(under, "out")

	if err, ok := res.(error); ok {
		return out, err
	}

	if err, ok := out.(error); ok {
		return out, err
	}

	return out, nil
}

func (m *Module) Call(under functs.Funct, args []functs.Funct) functs.Funct {
	main := m.Files.GetProp(under, "_")

	under.SetProp(under, "mod", m)
	under.SetProp(under, "file", main)

	return main.Call(under, args)
}

func (m *Module) GetProp(under functs.Funct, name string) functs.Funct {
	return m.Files.GetProp(under, name)
}

func (m *Module) SetProp(under functs.Funct, name string, val functs.Funct) functs.Funct {
	return m.Files.SetProp(under, name, val)
}

func (m *Module) String() string {
	return "{Module " + m.Name + " " + m.Files.String() + "}"
}
