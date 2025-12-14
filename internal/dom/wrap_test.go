package dom

import (
	"fmt"

	"go.ufukty.com/gss/internal/ast/html"
)

func ExampleWrap() {
	var (
		title           = &html.Text{Content: "Lorem ipsum"}
		titleContainer  = &html.Div{Classes: []string{"title"}, Children: []html.Element{title}}
		img             = &html.Img{Src: "./profile.png", SrcSet: map[float64]string{2.0: "./profile@2x.png"}}
		author          = &html.Text{Content: "Ufuktan Yıldırım"}
		authorContainer = &html.Div{Classes: []string{"author"}, Children: []html.Element{author}}
		main            = &html.Div{Id: "main", Children: []html.Element{titleContainer, img, authorContainer}}
		html            = &html.Html{Children: []html.Element{main}}
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
