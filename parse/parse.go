package parse

import (
	"bufio"
	"io"
	"strconv"
)

type Parse struct {
	Buf *bufio.Reader
	Loc Loc
}

func New(r io.Reader) *Parse {
	return &Parse{
		Buf: bufio.NewReader(r),
		Loc: Loc{},
	}
}

type Loc struct {
	Line, Col, Char int
}

func (l Loc) String() string {
	return strconv.Itoa(l.Line+1) + "," + strconv.Itoa(l.Col+1)
}
