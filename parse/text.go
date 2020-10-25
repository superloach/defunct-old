package parse

import "errors"

var ErrTextOneVal = errors.New("text should only capture one value")

type Text string

func (t *Text) Capture(values []string) error {
	if len(values) != 1 {
		return ErrTextOneVal
	}

	value := []rune(values[0])
	rs := make([]rune, 0, len(value))
	esc := false

	for _, r := range value {
		if !esc && r == '~' {
			esc = true
			continue
		}

		if esc {
			switch r {
			case '_', '[', ']', '{', '}', '~', ' ', '\n', '\r', '\t':
			default:
				rs = append(rs, '~')
			}
		}

		rs = append(rs, r)
		esc = false
	}

	*t = Text(rs)

	return nil
}

func (t *Text) String() string {
	value := []rune(*t)
	rs := make([]rune, 0, len(value))

	for _, r := range value {
		switch r {
		case '_', '[', ']', '{', '}', '~', ' ', '\n', '\r', '\t':
			rs = append(rs, '~')
		}

		rs = append(rs, r)
	}

	return string(rs)
}
