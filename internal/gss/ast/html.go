package ast

import "go.ufukty.com/gss/internal/gss/tokens"

type Element struct {
	Parent   *Element
	Children []*Element

	Tag tokens.Tag

	Id      string
	Classes []string
}

type HTML struct {
	Root *Element
}
