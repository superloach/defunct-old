package std

import "github.com/superloach/defunct/run"

var Std = run.Map{
	"print":  Print{},
	"doad":   Doad{},
	"parent": Parent{},
	"repr":   Repr{},
	"text":   Text{},
}
