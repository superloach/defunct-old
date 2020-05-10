package main

import (
	"flag"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/superloach/defunct"
	"github.com/superloach/defunct/debug"
	"github.com/superloach/defunct/functs"
)

func init() {
	flag.BoolVar(&debug.Debug, "debug", false, "print debug messages")
}

func main() {
	flag.Parse()
	oargs := flag.Args()

	fn := oargs[0]

	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	code, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	debug.Debugf("%#v\n", string(code))

	args := make(functs.Thing)
	for i, oarg := range oargs {
		args[strconv.Itoa(i)] = functs.String(oarg)
	}

	env := &defunct.Env{}
	under, err := env.RunProg(string(code), args)
	if err != nil {
		panic(err)
	}
	debug.Debugf("%s\n", under)
}
