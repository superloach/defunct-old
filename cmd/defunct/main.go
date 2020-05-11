package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/superloach/defunct/module"
	"github.com/superloach/defunct/types"
)

var debug = flag.Bool("debug", false, "print debug messages")

func debugf(f string, args ...interface{}) {
	if *debug {
		fmt.Printf(f, args...)
	}
}

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage: defunct [flags] (file or module)\n")
		flag.PrintDefaults()
	}

	flag.Parse()
	oargs := flag.Args()

	if len(oargs) == 0 {
		flag.Usage()
		return
	}

	fpath := oargs[0]

	args := make(types.Thing)
	for i, oarg := range oargs {
		args[strconv.Itoa(i)] = types.String(oarg)
	}
	debugf("%s\n", args)

	info, err := os.Stat(fpath)

	var mod *module.Module

	if err != nil || info.IsDir() {
		mod, err = module.LoadModule(fpath, debugf)
		if err != nil {
			panic(err)
		}
		debugf("%s\n", mod)
	} else {
		_, fname := filepath.Split(fpath)
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
			Files:  make(types.Thing),
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
