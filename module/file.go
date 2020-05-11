package module

import (
	"strings"

	"github.com/superloach/defunct/functs"
	"github.com/superloach/defunct/liner"
)

type File struct {
	Module *Module
	Path   string
	Name   string
	Lines  []*liner.Liner
}

func LoadFile(
	mod *Module,
	path, name, code string,
	debugf func(string, ...interface{}),
) *File {
	lines := strings.Split(code, "\n")

	file := &File{
		Module: mod,
		Path:   path,
		Name:   name,
		Lines:  make([]*liner.Liner, len(lines)),
	}

	for i, line := range lines {
		file.Lines[i] = liner.LoadLine(line, debugf)
	}

	return file
}

func (f *File) Call(under functs.Funct, args []functs.Funct) functs.Funct {
	for _, line := range f.Lines {
		_, err := line.Wrap.Run(under)
		if err != nil {
			return functs.Error(err.Error())
		}
	}

	return under.GetProp(under, "out")
}

func (f *File) GetProp(under functs.Funct, name string) functs.Funct {
	return under.GetProp(under, name)
}

func (f *File) SetProp(under functs.Funct, name string, val functs.Funct) functs.Funct {
	return under.SetProp(under, name, val)
}

func (f *File) String() string {
	return "{File " + f.Name + " (" + f.Path + ")}"
}
