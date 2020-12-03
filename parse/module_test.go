package parse_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/superloach/defunct/parse"
	"github.com/kr/pretty"
)

func TestParse(t *testing.T) {
	for i, src := range []string{
		`_[@print][hello world]`,
		`_[greet {_[@print][hello~ @s![_[@in][0]]]}]`,
		`_[greet][world]`,
	} {
		t.Run("src-"+strconv.Itoa(i+1), func(t *testing.T) {
			t.Logf("%q", src)

			p := parse.New(strings.NewReader(src))

			m, err := p.Module(t.Name())
			if err != nil {
				t.Errorf("parse %q: %#v", src, err)
				t.Errorf("%q %q %q", src[:p.Loc.Char], src[p.Loc.Char], src[p.Loc.Char+1:])
			}

			pretty.Println(m)

			t.Logf("%q", m.String())
		})
	}
}
