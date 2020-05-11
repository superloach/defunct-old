package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/superloach/defunct/functs"
	"github.com/superloach/defunct/module"
)

var (
	debug = flag.Bool("debug", false, "print debug messages")
)

func debugf(f string, args ...interface{}) {
	if *debug {
		fmt.Printf(f, args...)
	}
}

func main() {
	flag.Parse()
	oargs := flag.Args()

	args := make(functs.Thing)
	for i, oarg := range oargs {
		args[strconv.Itoa(i)] = functs.String(oarg)
	}
	debugf("%s\n", args)

	info, err := os.Stat(oargs[0])
	if err != nil {
		panic(err)
	}
	if info.IsDir() {
		mod, err := module.LoadModule(oargs[0], debugf)
		if err != nil {
			panic(err)
		}
		debugf("%s\n", mod)

		out, err := mod.Run(args)

		if err != nil {
			panic(err)
		}

		debugf("%s\n", out)
	} else {
		panic("single file stub")
	}
}
