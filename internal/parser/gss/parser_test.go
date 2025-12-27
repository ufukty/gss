package gss

import (
	"fmt"
	"os"
	"testing"

	"github.com/tdewolff/parse/v2"
	"github.com/tdewolff/parse/v2/css"
)

func TestParser(t *testing.T) {
	f, err := os.Open("testdata/simple.css")
	if err != nil {
		t.Fatalf("prep, open test file: %v", err)
	}
	p := css.NewParser(parse.NewInput(f), false)
	for ; err == nil; err = p.Err() {
		g, t, b := p.Next()
		if g == css.DeclarationGrammar {
			fmt.Println(g, "|", t, "|", string(b), "|", p.Values())
		} else {
			fmt.Println(g, "|", t, "|", string(b))
		}
	}
}
