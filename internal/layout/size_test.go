package layout

import (
	"testing"

	"go.ufukty.com/gss/internal/gss/ast"
	"go.ufukty.com/gss/internal/gss/tokens"
)

func Test_SizeUp(t *testing.T) {
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

}
