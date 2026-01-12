package core

import (
	"testing"

	"go.ufukty.com/gss/internal/parser/gss/csstokens"
	"go.ufukty.com/gss/internal/tokens"
)

func TestParseBorderStyle(t *testing.T) {
	ts, err := csstokens.Tokenize(`border-style: none`)
	if err != nil {
		t.Errorf("prep, unexpected error: %v", err)
	}
	tk := ts[0]
	bs, err := ParseBorderStyle(tk)
	if err != nil {
		t.Errorf("act, unexpected error: %v", err)
	}
	got, ok := bs.(tokens.BorderStyle)
	if !ok {
		t.Fatalf("assert, expected BorderStyle got %T", bs)
	}
	expected := tokens.BorderStyleNone
	if got != expected {
		t.Errorf("assert, expected %q, got %q", expected, got)
	}
}
