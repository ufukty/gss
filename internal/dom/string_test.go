package dom

import (
	"fmt"
	"strings"

	"go.ufukty.com/gommons/pkg/tree"
)

func selector(tag, id string, classes []string) string {
	s := tag
	if id != "" {
		s += "#" + id
	}
	if len(classes) > 0 {
		s += "." + strings.Join(classes, ".")
	}
	return s
}

func ss(children []Child) []string {
	cs := []string{}
	for _, c := range children {
		if s, ok := c.(fmt.Stringer); ok {
			cs = append(cs, s.String())
		}
	}
	return cs
}

func (e Div) String() string {
	return tree.List(selector("div", e.Ast.Id, e.Ast.Classes), ss(e.Children))
}

func (e Html) String() string {
	return tree.List("html", ss(e.Children))
}

func (e Img) String() string {
	s := selector("img", e.Ast.Id, e.Ast.Classes)
	if e.Ast.Src != "" {
		s = fmt.Sprintf("%s %q", s, e.Ast.Src)
	}
	for den, src := range e.Ast.SrcSet {
		s = fmt.Sprintf("%s %.0fx=>%q", s, den, src)
	}
	return s
}

func (e Span) String() string {
	return tree.List(selector("span", e.Ast.Id, e.Ast.Classes), ss(e.Children))
}

func (e TextNode) String() string {
	return fmt.Sprintf("text %q", e.Ast.Content)
}
