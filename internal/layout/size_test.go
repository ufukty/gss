package layout

import (
	"testing"

	"go.ufukty.com/gss/internal/gss/ast"
)

func Test_SizeUp(t *testing.T) {
	var (
		title  = &ast.Div{Classes: []string{"title"}}
		img    = &ast.Img{Src: "./profile.png"}
		author = &ast.Div{Classes: []string{"author"}}
		_      = &ast.Div{Id: "main", Children: []ast.Element{title, img, author}}
	)
}
