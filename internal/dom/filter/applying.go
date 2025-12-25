package filter

import (
	"go.ufukty.com/gss/internal/ast/ast"
	"go.ufukty.com/gss/internal/ast/html"
)

// returns if target ∋ selector(root)
func match(target html.Element, selector string) bool {
	return false
}

func Applying(target html.Element, rules []*ast.Rule) []*ast.Rule {
	as := []*ast.Rule{}
	for _, rule := range rules {
		if match(target, rule.Selector) {
			as = append(as, rule)
		}
	}
	return as
}
