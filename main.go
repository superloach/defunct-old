package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/superloach/defunct/parse"
	"github.com/superloach/defunct/run"
	"github.com/superloach/defunct/std"
)

const stdin = "-"

func main() {
	flag.Parse()

	file := stdin

	args := flag.Args()
	if len(args) > 0 {
		file = args[0]
		args = args[1:]
	}

	r := os.Stdin

	if file != stdin {
		f, err := os.Open(file)
		if err != nil {
			fmt.Println("open", err)
			os.Exit(1)
		}

		r = f
	}

	p := parse.New(r)

	m, err := p.Module(file)
	if err != nil {
		fmt.Println("parse error:", err)
		os.Exit(2)
	}

	in := run.StringsQueue(flag.Args()...)
	u := run.NewUnder(in, std.Std, os.Stdin, os.Stdout)

	ret, err := u.Run((*run.Module)(m))
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	if ret != run.Zilch {
		fmt.Println(ret)
	}
}
