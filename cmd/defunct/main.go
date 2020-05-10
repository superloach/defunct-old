package main

import (
	"flag"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/superloach/defunct"
	"github.com/superloach/defunct/functs"
)

var (
	debug = flag.Bool("debug", false, "print debug messages")
)

func main() {
	flag.Parse()
	oargs := flag.Args()

	args := make(functs.Thing)
	for i, oarg := range oargs {
		args[strconv.Itoa(i)] = functs.String(oarg)
	}

	arena := defunct.NewArena(args)
	arena.Debug = *debug

	f, err := os.Open(oargs[0])
	if err != nil {
		panic(err)
	}

	code, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	arena.Debugf("%#v\n", string(code))

	out, err := arena.RunString(string(code))
	if err != nil {
		panic(err)
	}
	arena.Debugf("%s\n", out)
}
