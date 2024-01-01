package logwrapper

import (
	uuid "github.com/satori/go.uuid"
)

type Span struct {
	ID     string
	parent *Span
}

func createSpan(parent *Span) *Span {
	s := &Span{
		ID:     uuid.NewV4().String(),
		parent: parent,
	}
	return s
}
