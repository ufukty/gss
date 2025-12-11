package dom

import (
	"fmt"

	"go.ufukty.com/gss/internal/gss/ast"
)

func ExampleWrap() {
	var (
		title           = &ast.Text{Content: "Lorem ipsum"}
		titleContainer  = &ast.Div{Classes: []string{"title"}, Children: []ast.Element{title}}
		img             = &ast.Img{Src: "./profile.png", SrcSet: map[float64]string{2.0: "./profile@2x.png"}}
		author          = &ast.Text{Content: "Ufuktan Yıldırım"}
		authorContainer = &ast.Div{Classes: []string{"author"}, Children: []ast.Element{author}}
		main            = &ast.Div{Id: "main", Children: []ast.Element{titleContainer, img, authorContainer}}
		html            = &ast.Html{Children: []ast.Element{main}}
	)

	d := Wrap(html)

	fmt.Println(d)
	// Output:
	// html
	// ╰─ div#main
	//    ├─ div.title
	//    │  ╰─ text "Lorem ipsum"
	//    ├─ img "./profile.png" 2x=>"./profile@2x.png"
	//    ╰─ div.author
	//       ╰─ text "Ufuktan Yıldırım"
}
