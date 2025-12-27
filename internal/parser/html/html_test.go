package html

import (
	"fmt"
	"os"
	"testing"

	"go.ufukty.com/gss/internal/parser/html/serialize"
	"golang.org/x/net/html"
)

func TestFindBody(t *testing.T) {
	f, err := os.Open("testdata/a.html")
	if err != nil {
		t.Errorf("prep, open test file: %v", err)
	}
	r, err := html.Parse(f)
	if err != nil {
		t.Errorf("prep, parse: %v", err)
	}
	b := findBody(r)
	fmt.Println(serialize.Node(b))
}
