package module

import (
	"strings"

	"github.com/superloach/defunct/liner"
	"github.com/superloach/defunct/types"
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

func (f *File) Call(under types.Funct, args []types.Funct) types.Funct {
	for _, line := range f.Lines {
		_, err := line.Wrap.Run(under)
		if err != nil {
			return types.Error(err.Error())
		}
	}

	return under.GetProp(under, "out")
}

func (f *File) GetProp(under types.Funct, name string) types.Funct {
	return under.GetProp(under, name)
}

func (f *File) SetProp(under types.Funct, name string, val types.Funct) types.Funct {
	return under.SetProp(under, name, val)
}

func (f *File) String() string {
	return "{File " + f.Name + " (" + f.Path + ")}"
}
