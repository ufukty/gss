package layout

import (
	"fmt"

	"go.ufukty.com/gss/internal/filter"
	"go.ufukty.com/gss/internal/gss/ast"
	"go.ufukty.com/gss/internal/layout/dom"
)

// TODO: account cascading
func Once(a ast.Element, rules []*ast.Rule, opts *dom.Options) error {
	d := dom.Wrap(a)

	appl := filter.Applying(a, rules)
	min, max, err := intrinsicSizes(a, appl, opts)
	if err != nil {
		return nil, fmt.Errorf("sizing element: %w", err)
	}

	d := &dimensional{
		element: a,
		min:     min,
		max:     max,
	}

	if s, ok := a.(Occupier); ok {
		if err := s.SizeUp(opts); err != nil {
			return nil, fmt.Errorf("size up: %w", err)
		}
	}

	if a, ok := a.(ast.Adopter); ok {
		for _, child := range a.GetChildren() {
			Once(child, rules, opts)
		}
	}

	return d
}
