package parse_test

import (
	"bytes"
	"testing"

	"github.com/alecthomas/participle/lexer"

	"github.com/superloach/defunct/parse"
)

func TestLexer(t *testing.T) {
	syms := lexer.SymbolsByRune(parse.Lexer)

	for _, test := range []struct {
		s    string
		toks [][2]string
	}{
		{"_", [][2]string{
			{"Under", "_"},
		}},

		{"hello world", [][2]string{
			{"Text", "hello"},
			{"Space", " "},
			{"Text", "world"},
		}},

		{"_[_ foo][bar]", [][2]string{
			{"Under", "_"},
			{"ArgsOpen", "["},
			{"Under", "_"},
			{"Space", " "},
			{"Text", "foo"},
			{"ArgsClose", "]"},
			{"ArgsOpen", "["},
			{"Text", "bar"},
			{"ArgsClose", "]"},
		}},

		{"~_~[~]~ ~\t~\n", [][2]string{
			{"Text", "~_~[~]~ ~\t~\n"},
		}},

		{"a b\tc\nd", [][2]string{
			{"Text", "a"},
			{"Space", " "},
			{"Text", "b"},
			{"Space", "\t"},
			{"Text", "c"},
			{"Space", "\n"},
			{"Text", "d"},
		}},
	} {
		r := bytes.NewBuffer([]byte(test.s))

		lex, err := parse.Lexer.Lex(r)
		if err != nil {
			t.Errorf("lex: %w", err)
			continue
		}

		toks, err := lexer.ConsumeAll(lex)
		if err != nil {
			t.Errorf("toks: %w", err)
			continue
		}

		for i, tok := range toks {
			if tok.EOF() {
				break
			}

			sym := syms[tok.Type]
			etok := test.toks[i]

			if etok[0] != sym {
				t.Errorf("expected %q, got %q", etok[0], sym)
			}

			if etok[1] != tok.Value {
				t.Errorf("expected %q, got %q", etok[1], tok.Value)
			}
		}
	}
}
