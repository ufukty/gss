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

func TextNode(n *html.Node) string {
	b := bytes.NewBufferString("")
	html.Render(b, &html.Node{
		Type:      n.Type,
		DataAtom:  n.DataAtom,
		Data:      n.Data,
		Namespace: n.Namespace,
		Attr:      n.Attr,
	})
	return strings.TrimSpace(b.String())
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
	b := bytes.NewBufferString("")
	html.Render(b, &html.Node{
		Type:      n.Type,
		DataAtom:  n.DataAtom,
		Data:      n.Data,
		Namespace: n.Namespace,
		Attr:      n.Attr,
	})
	return tree.List(b.String(), s)
}
