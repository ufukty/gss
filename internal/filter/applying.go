package filter

import (
	gss "go.ufukty.com/gss/internal/gss/ast"
	html "go.ufukty.com/gss/internal/html/ast"
)

// returns if target ∋ selector(root)
func match(target html.Element, selector string) bool {
	return false
}

func Applying(target html.Element, rules []*gss.Rule) []*gss.Rule {
	as := []*gss.Rule{}
	for _, rule := range rules {
		if match(target, rule.Selector) {
			as = append(as, rule)
		}
	}
	return as
}
