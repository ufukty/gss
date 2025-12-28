package gss

import (
	"fmt"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	f, err := os.Open("testdata/styles.css")
	if err != nil {
		t.Fatalf("prep, open test file: %v", err)
	}
	defer f.Close()
	r, err := Parse(f)
	if err != nil {
		t.Errorf("act, parse: %v", err)
	}

	for _, r := range r.Rules {
		fmt.Println(r)
	}
}
