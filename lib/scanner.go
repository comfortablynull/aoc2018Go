package lib

import (
	"bufio"
	"io"
	"strconv"
)

type AdventScanner struct {
	*bufio.Scanner
	rs io.ReadSeeker
}

func NewAdventScanner(r io.ReadSeeker) *AdventScanner {
	return &AdventScanner{Scanner: bufio.NewScanner(r), rs: r}
}

func (i *AdventScanner) Rewind() error {
	if _, err := i.rs.Seek(0, 0); err != nil {
		return err
	}
	i.Scanner = bufio.NewScanner(i.rs)
	return nil
}

func (i *AdventScanner) Int() (int64, error) {
	t := i.Text()
	if t[0] == '+' {
		t = t[1:]
	}
	return strconv.ParseInt(t, 10, 64)
}
