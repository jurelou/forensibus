package processors

import "strings"

type PError struct {
	Errors []error
}

func (e *PError) Add(err error) {
	e.Errors = append(e.Errors, err)
}

func (e *PError) Len() int {
	return len(e.Errors)
}

func (e *PError) Empty() bool {
	return len(e.Errors) == 0
}

func (e *PError) AsStrings() []string {
	var s []string

	for i := 0; i < 5 && i < len(e.Errors); i++ {
		s = append(s, e.Errors[i].Error())
	}
	return s
}

func (e *PError) AsJson() string {
	s := "["

	for i := 0; i < 5 && i < len(e.Errors); i++ {
		if i != 0 {
			s += ", "
		}
		s += "\""
		s += strings.Replace(e.Errors[i].Error(), "\"", "'", -1)
		s += "\""
	}
	s += "]"
	return s
}
