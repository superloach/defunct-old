package parse_test

import (
	"testing"

	"github.com/superloach/defunct/parse"
)

func TestParser(t *testing.T) {
	for _, test := range []string{
		"_[foo bar]",
		"_[foo bar]\n_[baz qux]",
	} {
		prog := parse.Prog{}

		err := parse.Parser.ParseString(test, &prog)
		if err != nil {
			t.Errorf("parse string %q: %w", test, err)
		}

		if test != prog.String() {
			t.Errorf("mismatch %q %q", test, prog.String())
		}
	}
}
