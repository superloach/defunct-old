package module

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/superloach/defunct/builtins"
	"github.com/superloach/defunct/types"
)

type Module struct {
	Name   string
	Path   string
	Files  types.Thing
	Debugf func(string, ...interface{})
}

func LoadModule(
	mod string, // @foo/bar/baz
	debugf func(string, ...interface{}),
) (*Module, error) {
	if len(mod) == 0 {
		return nil, fmt.Errorf("zero len mod")
	}

	_, modn := path.Split(mod)

	m := &Module{
		Name:   modn, // baz
		Path:   mod,  // @foo/bar/baz
		Files:  make(types.Thing),
		Debugf: debugf,
	}

	isstd := mod[0] == '@'
	if isstd {
		debugf("isstd\n")
	}

	var dirn string
	var dir http.File
	var err error
	if isstd {
		dirn = "/" + mod[1:]
		dir, err = Stdlib.Open(dirn) // /foo/bar/baz
	} else {
		dirn = filepath.Join(strings.Split(mod, "/")...)
		dir, err = os.Open(dirn)
	}
	debugf("%s\n", dirn)
	if err != nil {
		return m, err
	}

	finfos, err := dir.Readdir(-1)
	if err != nil {
		return m, err
	}

	for _, finfo := range finfos {
		fname := finfo.Name()

		if !strings.HasSuffix(fname, ".df") {
			continue
		}

		name := fname[:len(fname)-3]
		fpath := path.Join(dirn, fname)

		var f http.File
		if isstd {
			f, err = Stdlib.Open(fpath)
		} else {
			f, err = os.Open(fpath)
		}
		if err != nil {
			return m, err
		}

		code, err := ioutil.ReadAll(f)
		if err != nil {
			return m, err
		}

		m.Files[name] = LoadFile(m, fpath, name, string(code), debugf)
	}

	return m, nil
}

func (m *Module) Run(in types.Funct) (types.Funct, error) {
	under := &types.Under{
		Arbitrary: types.Thing{
			"in": in,
		},
		Builtins: builtins.Builtins,
		Modules:  types.Zilch,
		Debugf:   m.Debugf,
	}

	m.Debugf("under %s\n", under)

	res := m.Call(under, make([]types.Funct, 0))
	out := under.GetProp(under, "out")

	if err, ok := res.(error); ok {
		return out, err
	}

	if err, ok := out.(error); ok {
		return out, err
	}

	return out, nil
}

func (m *Module) Call(under types.Funct, args []types.Funct) types.Funct {
	main := m.Files.GetProp(under, "_")
	m.Debugf("main %s\n", main)

	under.SetProp(under, "mod", m)
	under.SetProp(under, "file", main)

	return main.Call(under, args)
}

func (m *Module) GetProp(under types.Funct, name string) types.Funct {
	return m.Files.GetProp(under, name)
}

func (m *Module) SetProp(under types.Funct, name string, val types.Funct) types.Funct {
	return m.Files.SetProp(under, name, val)
}

func (m *Module) String() string {
	return "{Module " + m.Name + " (" + m.Path + ")}"
}
