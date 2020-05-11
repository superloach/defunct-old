package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"

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
	fpath := oargs[0]

	args := make(functs.Thing)
	for i, oarg := range oargs {
		args[strconv.Itoa(i)] = functs.String(oarg)
	}
	debugf("%s\n", args)

	info, err := os.Stat(fpath)
	if err != nil {
		panic(err)
	}

	var mod *module.Module

	if info.IsDir() {
		mod, err = module.LoadModule(fpath, debugf)
		if err != nil {
			panic(err)
		}
		debugf("%s\n", mod)
	} else {
		_, fname := path.Split(fpath)
		name := strings.ToLower(fname[:len(fname)-3])

		f, err := os.Open(fpath)
		if err != nil {
			panic(err)
		}

		code, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}

		mod = &module.Module{
			Name:   "single-file",
			Files:  make(functs.Thing),
			Debugf: debugf,
		}

		mod.Files["_"] = module.LoadFile(mod, fpath, name, string(code), debugf)
	}

	out, err := mod.Run(args)

	if err != nil {
		panic(err)
	}

	debugf("%s\n", out)
}
