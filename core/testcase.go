package core

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"github.com/camypaper/libra"
)

/*
Testcase shows Testcase
*/
type Testcase struct {
	Name string
	In   *string
	Out  *string
}

type input struct {
	name   string
	reader io.Reader
}

func (i input) Name() string {
	return i.name
}
func (i input) Reader() io.Reader {
	return i.reader
}

/*
ToInput converts to input.
*/
func (t Testcase) ToInput() (libra.Input, error) {
	if t.In == nil {
		return nil, errors.New("input is nil")
	}
	ret := input{}
	ret.name = t.Name
	ret.reader = bytes.NewBufferString(*t.In)
	return ret, nil
}
func (t Testcase) String() string {
	trim := func(s *string) string {
		if s == nil {
			return "<nil>"
		}
		if len(*s) <= 20 {
			return *s
		}
		return (*s)[0:20] + "..."
	}
	return fmt.Sprintf("[Name:%v, In:%v, Out:%v]", t.Name, trim(t.In), trim(t.Out))
}
