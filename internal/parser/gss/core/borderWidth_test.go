package core

import (
	"testing"

	"github.com/tdewolff/parse/v2/css"
	"go.ufukty.com/gss/internal/dimensional"
	"go.ufukty.com/gss/internal/parser/gss/csstokens"
)

func TestIsBorderWidth_positive(t *testing.T) {
	tcs := []string{
		"1px",
		"2.5em",
		"10rem",
		"0",
		"inherit",
		"initial",
	}

	for _, tc := range tcs {
		t.Run(tc, func(t *testing.T) {
			toks, err := csstokens.Tokenize("border-width: " + tc)
			if err != nil {
				t.Errorf("prep, unexpected error: %v", err)
			}
			if len(toks) == 0 {
				t.Fatal("prep, no tokens returned")
			}
			if !IsBorderWidth(toks[0]) {
				t.Fatal("unexpectedly false")
			}
		})
	}
}

func TestParseBorderWidth_dimensions(t *testing.T) {
	type tc struct {
		name     string
		input    string
		expected dimensional.Dimension
	}
	tcs := []tc{
		{"1px", "1px", dimensional.New(1, dimensional.Px)},
		{"2pt", "2pt", dimensional.New(2, dimensional.Pt)},
		{"3em", "3em", dimensional.New(3, dimensional.Em)},
		{"4rem", "4rem", dimensional.New(4, dimensional.Rem)},
		{"5vw", "5vw", dimensional.New(5, dimensional.Vw)},
		{"6vh", "6vh", dimensional.New(6, dimensional.Vh)},
		{"1.5px", "1.5px", dimensional.New(1.5, dimensional.Px)},
		{"2.25em", "2.25em", dimensional.New(2.25, dimensional.Em)},
		{"10.0px", "10.0px", dimensional.New(10.0, dimensional.Px)},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			toks, err := csstokens.Tokenize("border-width: " + tc.input)
			if err != nil {
				t.Errorf("prep, unexpected error: %v", err)
			}
			if len(toks) == 0 {
				t.Fatal("prep, no tokens returned")
			}

			got, err := ParseBorderWidth(toks[0])
			if err != nil {
				t.Fatalf("act, unexpected error: %v", err)
			}

			dim, ok := got.(dimensional.Dimension)
			if !ok {
				t.Fatalf("assert, expected dimensional.Dimension got %T", got)
			}

			if !dimensional.Compare(dim, tc.expected) {
				t.Errorf("assert, expected %v, got %v", tc.expected, dim)
			}
		})
	}
}

func TestParseBorderWidth_zero(t *testing.T) {
	toks, err := csstokens.Tokenize("border-width: 0")
	if err != nil {
		t.Errorf("prep, unexpected error: %v", err)
	}
	if len(toks) == 0 {
		t.Fatal("prep, no tokens returned")
	}

	got, err := ParseBorderWidth(toks[0])
	if err != nil {
		t.Fatalf("act, unexpected error: %v", err)
	}

	dim, ok := got.(dimensional.Dimension)
	if !ok {
		t.Fatalf("assert, expected dimensional.Dimension got %T", got)
	}

	expected := dimensional.New(0)
	if !dimensional.Compare(dim, expected) {
		t.Errorf("assert, expected %v, got %v", expected, dim)
	}
}

func TestParseBorderWidth_globalKeywords(t *testing.T) {
	tcs := []string{
		"inherit",
		"initial",
		"unset",
		"revert",
	}

	for _, tc := range tcs {
		t.Run(tc, func(t *testing.T) {
			toks, err := csstokens.Tokenize("border-width: " + tc)
			if err != nil {
				t.Errorf("prep, unexpected error: %v", err)
			}
			if len(toks) == 0 {
				t.Fatal("prep, no tokens returned")
			}

			got, err := ParseBorderWidth(toks[0])
			if err != nil {
				t.Fatalf("act, unexpected error: %v", err)
			}

			str, ok := got.(string)
			if !ok {
				t.Fatalf("assert, expected string got %T", got)
			}

			if str != tc {
				t.Errorf("assert, expected %q, got %q", tc, str)
			}
		})
	}
}

func TestParseBorderWidth_invalidNumber(t *testing.T) {
	// Non-zero numbers without units should fail
	tok := css.Token{
		TokenType: css.NumberToken,
		Data:      []byte("5"),
	}

	_, err := ParseBorderWidth(tok)
	if err == nil {
		t.Error("expected error for non-zero number without unit, got nil")
	}
}

func TestParseBorderWidth_invalidToken(t *testing.T) {
	// String token should fail
	tok := css.Token{
		TokenType: css.StringToken,
		Data:      []byte("invalid"),
	}

	_, err := ParseBorderWidth(tok)
	if err == nil {
		t.Error("expected error for invalid token type, got nil")
	}
}
