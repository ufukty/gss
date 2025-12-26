package serialize

import (
	"bytes"
	"fmt"
	"strings"

	"go.ufukty.com/gommons/pkg/tree"
	"golang.org/x/net/html"
)

func node(n *html.Node) string {
	b := bytes.NewBufferString("")
	html.Render(b, &html.Node{
		Type:      n.Type,
		DataAtom:  n.DataAtom,
		Data:      n.Data,
		Namespace: n.Namespace,
		Attr:      n.Attr,
	})
	return b.String()
}

func TextNode(n *html.Node) string {
	return strings.TrimSpace(node(n))
}

func Node(n *html.Node) string {
	s := []string{}
	for c := range n.ChildNodes() {
		if c.Type == html.ElementNode {
			s = append(s, Node(c))
		} else if c.Type == html.TextNode {
			if f := TextNode(c); f != "" {
				s = append(s, fmt.Sprintf("%q", f))
			}
		}
	}
	return tree.List(node(n), s)
}
