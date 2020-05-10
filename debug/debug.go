package debug

import "fmt"

var Debug bool = false

func Debugf(f string, args ...interface{}) {
	if Debug {
		fmt.Printf(f, args...)
	}
}
