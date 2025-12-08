package ast

import "go.ufukty.com/gss/internal/gss/tokens"

type Element struct {
	Tag        tokens.Tag
	Id         string
	Classes    []string
	Attributes map[string]string // <img src> etc.
	Parent     *Element
	Children   []*Element
}

type Html struct {
	Root *Element
}
