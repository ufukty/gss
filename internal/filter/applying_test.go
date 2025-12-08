package filter

import (
	"reflect"
	"testing"

	"go.ufukty.com/gss/internal/gss/ast"
	"go.ufukty.com/gss/internal/gss/tokens"
)

func Test_Applying(t *testing.T) {
	var (
		title = &ast.Element{
			Tag:     tokens.Tag_Div,
			Classes: []string{"title"},
		}
		img = &ast.Element{
			Tag: tokens.Tag_Img,
			Attributes: map[string]string{
				"src": "./profile.png",
			},
		}
		author = &ast.Element{
			Tag:     tokens.Tag_Div,
			Classes: []string{"author"},
		}
		main = &ast.Element{
			Tag:      tokens.Tag_Div,
			Id:       "main",
			Children: []*ast.Element{title, img, author},
		}
	)

	var ( // addressable to compare
		a = &ast.Styles{}
		b = &ast.Styles{}
		c = &ast.Styles{}
		d = &ast.Styles{}
		e = &ast.Styles{}
		f = &ast.Styles{}
	)

	gss := &ast.Gss{
		Rules: []*ast.Rule{
			&ast.Rule{"main", a},
			&ast.Rule{".title", b},
			&ast.Rule{"div.title", c},
			&ast.Rule{"img", d},
			&ast.Rule{"main .author", e},
			&ast.Rule{"main > .author", f},
		},
	}

	tcs := map[*ast.Element][]*ast.Styles{
		main:   []*ast.Styles{a},
		title:  []*ast.Styles{b, c},
		img:    []*ast.Styles{d},
		author: []*ast.Styles{e, f},
	}

	for target, expected := range tcs {
		t.Run(reflect.ValueOf(target).String(), func(t *testing.T) {
			got := Applying(target, gss.Rules)
			if len(expected) != len(got) {
				t.Error("expected %v got %v", expected, got)
			}
		})
	}
}
