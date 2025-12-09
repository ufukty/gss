package layout

import (
	"fmt"
	"slices"

	"go.ufukty.com/gss/internal/filter"
	"go.ufukty.com/gss/internal/gss/ast"
	"go.ufukty.com/gss/internal/gss/tokens"
)

type opts struct {
	Width, Height, Density float64
}

type size struct {
	width, height float64
}

type dimensional struct {
	element  *ast.Element
	min, max size
	parent   *dimensional
	children []*dimensional
}

var leaves = []tokens.Tag{tokens.Tag_Img}

// TODO: account cascading
// TODO: composite pattern?
func intrinsicSizes(elem *ast.Element, applying []*ast.Rule, opts *opts) (min, max size, err error) {
	if !slices.Contains(leaves, elem.Tag) {
		return
	}

	if elem.Tag == tokens.Tag_Img {
		s, err := imgSize(elem, opts)
		if err != nil {
			return min, max, fmt.Errorf("sizing image: %w", err)
		}
		return *s, *s, nil
	}

	return
}

func sizeUpRec(elem *ast.Element, rules []*ast.Rule, opts *opts) (*dimensional, error) {
	appl := filter.Applying(elem, rules)
	min, max, err := intrinsicSizes(elem, appl, opts)
	if err != nil {
		return nil, fmt.Errorf("sizing element: %w", err)
	}

	d := &dimensional{
		element:  elem,
		min:      min,
		max:      max,
		parent:   &dimensional{},
		children: []*dimensional{},
	}

	for _, child := range elem.Children {
		c, err := sizeUpRec(child, rules, opts)
		if err != nil {
			return nil, fmt.Errorf("re: %w", err)
		}
		c.parent = d
		d.children = append(d.children, c)
	}

	return d, nil
}

func sizeUp(html *ast.Html, gss *ast.Gss, opts *opts) (*dimensional, error) {
	return sizeUpRec(html.Root, gss.Rules, opts)
}
