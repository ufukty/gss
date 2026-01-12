package core

import (
	"fmt"
	"image/color"
	"testing"

	"github.com/tdewolff/parse/v2/css"
	"go.ufukty.com/gss/internal/parser/gss/csstokens"
)

func TestIsColor_positive(t *testing.T) {
	tcs := []string{
		"#fff",
		"#ffff",
		"#ffffff",
		"#ffffffff",

		"red",
		"ReD",
		"transparent",
		"currentColor",
	}

	for _, tc := range tcs {
		t.Run(tc, func(t *testing.T) {
			toks, err := csstokens.Tokenize(fmt.Sprintf("border-color: %s", tc))
			if err != nil {
				t.Errorf("prep, unexpected error: %v", err)
			}
			if !IsColor(toks[0]) {
				t.Fatal("unexpectedly false")
			}
		})
	}
}

func hashToken(data string) css.Token {
	return css.Token{
		TokenType: css.HashToken,
		Data:      []byte(data),
	}
}

func Test_ParseColor(t *testing.T) {
	type tc struct {
		name     string
		input    css.Token
		expected color.RGBA
	}
	redOpaque := color.RGBA{255, 0, 0, 255}
	redHalf := color.RGBA{128, 0, 0, 128}
	tcs := []tc{
		{"rgb", hashToken("#f00"), redOpaque},
		{"RGB", hashToken("#F00"), redOpaque},
		{"rgba", hashToken("#f00f"), redOpaque},
		{"RGBA", hashToken("#F00F"), redOpaque},
		{"rrggbb", hashToken("#ff0000"), redOpaque},
		{"RRGGBB", hashToken("#FF0000"), redOpaque},
		{"rrggbbaa", hashToken("#ff0000ff"), redOpaque},
		{"RRGGBBAA", hashToken("#FF0000FF"), redOpaque},

		{"rrggbbaa:half", hashToken("#ff000080"), redHalf},
		{"RRGGBBAA:half", hashToken("#FF000080"), redHalf},

		{"red", hashToken("#FF0000FF"), redOpaque},
		{"RED", hashToken("#FF0000FF"), redOpaque},
		{"Red", hashToken("#FF0000FF"), redOpaque},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseColor(tc.input)
			if err != nil {
				t.Fatalf("act, unexpected error: %v", err)
			}
			if tc.expected != got {
				t.Errorf("assert, expected %v, got %v", tc.expected, got)
			}
			c, ok := got.(color.RGBA) // since all testcases resolves to a color
			if !ok {
				t.Fatalf("assert, expected color, got %T", got)
			}
			if c != tc.expected {
				t.Error()
			}
		})
	}
}
