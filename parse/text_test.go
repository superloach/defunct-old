package parse_test

import (
	"testing"

	"github.com/superloach/defunct/parse"
)

func TestText(t *testing.T) {
	for _, test := range [][2]string{
		{"foo", "foo"},
		{"foo~ bar", "foo bar"},
	} {
		text := parse.Text("")

		err := text.Capture([]string{test[0]})
		if err != nil {
			t.Errorf("text capture %q: %w", test[0], err)
		}

		if string(text) != test[1] {
			t.Errorf("mismatch1 %q %q", string(text), test[1])
		}

		if text.String() != test[0] {
			t.Errorf("mismatch0 %q %q", text.String(), test[0])
		}
	}
}
