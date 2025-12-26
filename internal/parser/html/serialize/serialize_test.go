package serialize

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func ExampleNode() {
	i, err := os.Open("testdata/a.html")
	if err != nil {
		panic(fmt.Errorf("prep, open test file: %v", err))
	}
	defer i.Close()
	z, err := html.Parse(i)
	if err != nil {
		panic(fmt.Errorf("act, parse: %v", err))
	}
	fmt.Println(Node(z.FirstChild))
	// Output:
	// <html></html>
	// ├─ <head></head>
	// ╰─ <body></body>
	//    ╰─ <main></main>
	//       ├─ <div id="title"></div>
	//       │  ╰─ "Lorem ipsum"
	//       ├─ <img src="a.jpg"/>
	//       ╰─ <div class="author"></div>
	//          ╰─ "Dolor sit amet"
}
