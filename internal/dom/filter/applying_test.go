package filter

import (
	"reflect"
	"testing"

	"go.ufukty.com/gss/internal/ast/ast"
	"go.ufukty.com/gss/internal/ast/html"
)

func Test_Applying(t *testing.T) {
	var (
		title  = &html.Div{Classes: []string{"title"}}
		img    = &html.Img{}
		author = &html.Div{Classes: []string{"author"}}
		main   = &html.Div{Id: "main", Children: []html.Element{title, img, author}}
	)

	var ( // addressable to compare
		a = &ast.Styles{}
		b = &ast.Styles{}
		c = &ast.Styles{}
		d = &ast.Styles{}
		e = &ast.Styles{}
		f = &ast.Styles{}
	)

	g := &ast.Gss{
		Rules: []*ast.Rule{
			{"main", a},
			{".title", b},
			{"div.title", c},
			{"img", d},
			{"main .author", e},
			{"main > .author", f},
		},
	}

	tcs := map[html.Element][]*ast.Styles{
		main:   {a},
		title:  {b, c},
		img:    {d},
		author: {e, f},
	}

	for target, expected := range tcs {
		t.Run(reflect.ValueOf(target).String(), func(t *testing.T) {
			got := Applying(target, g.Rules)
			if len(expected) != len(got) {
				t.Errorf("expected %v got %v", expected, got)
			}
		})
	}
}
