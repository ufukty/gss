package serialize

import (
	"bytes"
	"fmt"
	"slices"
	"strings"

	"go.ufukty.com/gommons/pkg/tree"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
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

var tags = []atom.Atom{atom.Html, atom.Body, atom.Div, atom.Span, atom.Img, atom.Main}

func Node(n *html.Node) string {
	s := []string{}
	for c := range n.ChildNodes() {
		if c.Type == html.ElementNode && slices.Contains(tags, c.DataAtom) {
			s = append(s, Node(c))
		} else if c.Type == html.TextNode {
			if f := TextNode(c); f != "" {
				s = append(s, fmt.Sprintf("%q", f))
			}
		}
	}
	return tree.List(node(n), s)
}
