package is

import (
	"fmt"
	"testing"

	"github.com/tdewolff/parse/v2/css"
)

func TestTrimSpaces(t *testing.T) {
	input := []css.Token{
		{TokenType: css.WhitespaceToken},
		{TokenType: css.WhitespaceToken},
		{TokenType: css.IdentToken},
		{TokenType: css.IdentToken},
		{TokenType: css.IdentToken},
		{TokenType: css.IdentToken},
		{TokenType: css.WhitespaceToken},
		{TokenType: css.WhitespaceToken},
		{TokenType: css.WhitespaceToken},
	}
	trimmed := trimSpaces(input)
	got := len(trimmed)
	if got != 4 {
		t.Errorf("expected 4 got %d", got)
	}
	for _, tk := range trimmed {
		if tk.TokenType == css.WhitespaceToken {
			t.Errorf("unexpected whitespace token in trimmed array")
		}
	}
}

func TestColor_positive(t *testing.T) {
	tcs := []string{
		"#fff",
		"#ffff",
		"#ffffff",
		"#ffffffff",

		"red",
		"ReD",
		"transparent",
		"currentColor",

		"rgb(0 0 0)",
		"rgb(0, 0, 0)",
		"rgba(0,0,0,0.5)",
		"hsl(120 100% 50%)",
		"hwb(120 0% 0%)",
		"lab(29.2345% 39.3825 20.0664)",
		"oklch(62.2345% 0.12 120)",
		"color(display-p3 1 0 0)",
	}

	for _, tc := range tcs {
		t.Run(tc, func(t *testing.T) {
			toks, _ := tokenize(fmt.Sprintf("border-color: %s", tc))
			if !Color(toks[0]) {
				t.Fatal("unexpectedly false")
			}
		})
	}
}

func TestColor_negative(t *testing.T) {
	tcs := []string{
		"#ff",
		"#fffff",
		"#ggg",
		"#",

		"notacolor",

		"rgb()",
		"rgb(",
		"unknown(1 2 3)",

		"red solid 1px",
		"#fff !important",
		"var(--x)",
	}

	for _, tc := range tcs {
		t.Run(tc, func(t *testing.T) {
			toks, _ := tokenize(fmt.Sprintf("border-color: %s", tc))
			if Color(toks[0]) {
				t.Fatalf("unexpectedly true")
			}
		})
	}
}
