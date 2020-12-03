package run

import (
	"github.com/superloach/defunct/parse"
)

type Texter interface {
	Text() Text
}

type Text parse.Text

func TextOf(f Funct) Text {
	if t, ok := f.(Texter); ok {
		return t.Text()
	}

	return Text(f.String())
}

func (t Text) String() string {
	return "text(" + parse.Text(t).EscapeString() + ")"
}

func (t Text) Text() Text {
	return t
}

func (t Text) Call(_ *Under, _ Queue) (Funct, error) {
	panic("format stub")
}
