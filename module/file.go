package module

import (
	"github.com/superloach/defunct/liner"
	"github.com/superloach/defunct/types"
)

type File struct {
	Module *Module
	Path   string
	Name   string
	Liner  *liner.Liner
	Debugf func(string, ...interface{})
}

func LoadFile(
	mod *Module,
	path, name, code string,
	debugf func(string, ...interface{}),
) *File {
	debugf("load file %s %s %#v\n", path, name, code)
	file := &File{
		Module: mod,
		Path:   path,
		Name:   name,
		Liner:  liner.LoadLines(code, debugf),
		Debugf: debugf,
	}
	debugf("loaded %s\n", file.Liner)
	return file
}

func (f *File) Call(under types.Funct, args []types.Funct) types.Funct {
	f.Debugf("liner %#v\n", f.Liner)
	f.Debugf("lines %#v\n", f.Liner.Lines)

	for _, line := range f.Liner.Lines {
		f.Debugf("line %#v\n", line)
		_, err := line.Run(under)
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
