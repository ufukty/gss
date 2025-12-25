package filter

import (
	"go.ufukty.com/gss/internal/ast/ast"
)

// returns if target ∋ selector(root)
func match(target ast.Element, selector string) bool {
	return false
}

func Applying(target ast.Element, rules []*ast.Rule) []*ast.Rule {
	as := []*ast.Rule{}
	for _, rule := range rules {
		if match(target, rule.Selector) {
			as = append(as, rule)
		}
	}
	return as
}
