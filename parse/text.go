package parse

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

var (
	ErrNoText        = errors.New("no text")
	ErrIncompEscape  = errors.New("incomplete escape")
	ErrInvalidEscape = errors.New("invalid escape")
)

type Text []rune

func UnescapeString(s string) (Text, error) {
	return New(strings.NewReader(s)).Text()
}

func UnescapeRunes(rs []rune) (Text, error) {
	return UnescapeString(string(rs))
}

func (p *Parse) text() ([]rune, error) {
	t := []rune(nil)

	for {
		r, err := p.peek()
		if errors.Is(err, io.EOF) || strings.ContainsRune(StringEnder, r) {
			return t, nil
		}

		if err != nil {
			return nil, fmt.Errorf("peek: %w", err)
		}

		p.incr()

		if r == RuneEscape {
			n, err := p.peek()
			if err != nil {
				if errors.Is(err, io.EOF) {
					return nil, ErrIncompEscape
				}

				return nil, fmt.Errorf("peek: %w", err)
			}

			if !strings.ContainsRune(StringEscape, n) {
				return nil, ErrInvalidEscape
			}

			p.incr()

			r = n
		}

		t = append(t, r)
	}
}

func (p *Parse) Text() (Text, error) {
	t, err := p.text()
	if err != nil {
		return nil, fmt.Errorf("text: %w", err)
	}

	if len(t) == 0 {
		return nil, ErrNoText
	}

	return Text(t), nil
}

func (t Text) EscapeRunes() []rune {
	trs := []rune(t)
	rs := make([]rune, 0, len(trs))

	for _, r := range trs {
		if strings.ContainsRune(StringEscape, r) {
			rs = append(rs, RuneEscape)
		}

		rs = append(rs, r)
	}

	return rs
}

func (t Text) EscapeString() string {
	return string(t.EscapeRunes())
}

func (t Text) String() string {
	return t.EscapeString()
}
