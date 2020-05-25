package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

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

	args := make([]types.Funct, len(oargs))
	for i, oarg := range oargs {
		args[i] = types.String(oarg)
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
		debugf("%s\n", fname)

		f, err := os.Open(fpath)
		if err != nil {
			panic(err)
		}
		debugf("opened %s\n", fpath)

		code, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}
		debugf("%#v\n", string(code))

		mod = &module.Module{
			Name:   "",
			Path:   fpath,
			Files:  make(types.Thing),
			Debugf: debugf,
		}
		debugf("%s\n", mod)

		mod.Files["_"] = module.LoadFile(mod, fpath, fname, string(code), debugf)
		debugf("%s\n", mod.Files)
	}

	debugf("mod %s\n", mod)

	out, err := mod.Run(args)

	if err != nil {
		panic(err)
	}

	debugf("%s\n", out)
}
