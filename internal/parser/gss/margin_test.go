package gss

import (
	"testing"

	"go.ufukty.com/gss/internal/ast"
	"go.ufukty.com/gss/internal/dimensional"
	"go.ufukty.com/gss/internal/parser/gss/csstokens"
)

func TestParseMargin_singleValue(t *testing.T) {
	// margin: 10px; (all sides)
	ts, err := csstokens.Tokenize("margin: 10px")
	if err != nil {
		t.Errorf("prep, unexpected error: %v", err)
	}

	got, err := ParseMargin(ts)
	if err != nil {
		t.Errorf("act, unexpected error: %v", err)
	}

	expected := ast.Margin{
		Top:    dimensional.New(10, dimensional.Px),
		Right:  dimensional.New(10, dimensional.Px),
		Bottom: dimensional.New(10, dimensional.Px),
		Left:   dimensional.New(10, dimensional.Px),
	}

	if !got.IsEqual(expected) {
		t.Errorf("assert, expected %#v, got %#v", expected, got)
	}
}

func TestParseMargin_twoValues(t *testing.T) {
	// margin: 10px 20px; (vertical horizontal)
	ts, err := csstokens.Tokenize("margin: 10px 20px")
	if err != nil {
		t.Errorf("prep, unexpected error: %v", err)
	}

	got, err := ParseMargin(ts)
	if err != nil {
		t.Errorf("act, unexpected error: %v", err)
	}

	expected := ast.Margin{
		Top:    dimensional.New(10, dimensional.Px),
		Right:  dimensional.New(20, dimensional.Px),
		Bottom: dimensional.New(10, dimensional.Px),
		Left:   dimensional.New(20, dimensional.Px),
	}

	if !got.IsEqual(expected) {
		t.Errorf("assert, expected %#v, got %#v", expected, got)
	}
}

func TestParseMargin_threeValues(t *testing.T) {
	// margin: 10px 20px 30px; (top horizontal bottom)
	ts, err := csstokens.Tokenize("margin: 10px 20px 30px")
	if err != nil {
		t.Errorf("prep, unexpected error: %v", err)
	}

	got, err := ParseMargin(ts)
	if err != nil {
		t.Errorf("act, unexpected error: %v", err)
	}

	expected := ast.Margin{
		Top:    dimensional.New(10, dimensional.Px),
		Right:  dimensional.New(20, dimensional.Px),
		Bottom: dimensional.New(30, dimensional.Px),
		Left:   dimensional.New(20, dimensional.Px),
	}

	if !got.IsEqual(expected) {
		t.Errorf("assert, expected %#v, got %#v", expected, got)
	}
}

func TestParseMargin_fourValues(t *testing.T) {
	// margin: 10px 20px 30px 40px; (top right bottom left)
	ts, err := csstokens.Tokenize("margin: 10px 20px 30px 40px")
	if err != nil {
		t.Errorf("prep, unexpected error: %v", err)
	}

	got, err := ParseMargin(ts)
	if err != nil {
		t.Errorf("act, unexpected error: %v", err)
	}

	expected := ast.Margin{
		Top:    dimensional.New(10, dimensional.Px),
		Right:  dimensional.New(20, dimensional.Px),
		Bottom: dimensional.New(30, dimensional.Px),
		Left:   dimensional.New(40, dimensional.Px),
	}

	if !got.IsEqual(expected) {
		t.Errorf("assert, expected %#v, got %#v", expected, got)
	}
}

func TestParseMargin_autoKeyword(t *testing.T) {
	// margin: auto;
	ts, err := csstokens.Tokenize("margin: auto")
	if err != nil {
		t.Errorf("prep, unexpected error: %v", err)
	}

	got, err := ParseMargin(ts)
	if err != nil {
		t.Errorf("act, unexpected error: %v", err)
	}

	expected := ast.Margin{
		Top:    "auto",
		Right:  "auto",
		Bottom: "auto",
		Left:   "auto",
	}

	if !got.IsEqual(expected) {
		t.Errorf("assert, expected %#v, got %#v", expected, got)
	}
}

func TestParseMargin_mixed(t *testing.T) {
	// margin: 10px auto;
	ts, err := csstokens.Tokenize("margin: 10px auto")
	if err != nil {
		t.Errorf("prep, unexpected error: %v", err)
	}

	got, err := ParseMargin(ts)
	if err != nil {
		t.Errorf("act, unexpected error: %v", err)
	}

	expected := ast.Margin{
		Top:    dimensional.New(10, dimensional.Px),
		Right:  "auto",
		Bottom: dimensional.New(10, dimensional.Px),
		Left:   "auto",
	}

	if !got.IsEqual(expected) {
		t.Errorf("assert, expected %#v, got %#v", expected, got)
	}
}

func TestParseMargin_zero(t *testing.T) {
	// margin: 0;
	ts, err := csstokens.Tokenize("margin: 0")
	if err != nil {
		t.Errorf("prep, unexpected error: %v", err)
	}

	got, err := ParseMargin(ts)
	if err != nil {
		t.Errorf("act, unexpected error: %v", err)
	}

	expected := ast.Margin{
		Top:    dimensional.New(0),
		Right:  dimensional.New(0),
		Bottom: dimensional.New(0),
		Left:   dimensional.New(0),
	}

	if !got.IsEqual(expected) {
		t.Errorf("assert, expected %#v, got %#v", expected, got)
	}
}

func TestParseMargin_globalKeywords(t *testing.T) {
	keywords := []string{"inherit", "initial", "unset"}

	for _, kw := range keywords {
		t.Run(kw, func(t *testing.T) {
			ts, err := csstokens.Tokenize("margin: " + kw)
			if err != nil {
				t.Errorf("prep, unexpected error: %v", err)
			}

			got, err := ParseMargin(ts)
			if err != nil {
				t.Errorf("act, unexpected error: %v", err)
			}

			expected := ast.Margin{
				Top:    kw,
				Right:  kw,
				Bottom: kw,
				Left:   kw,
			}

			if !got.IsEqual(expected) {
				t.Errorf("assert, expected %#v, got %#v", expected, got)
			}
		})
	}
}

func TestParseMargin_invalidValueCount(t *testing.T) {
	// margin: 1px 2px 3px 4px 5px; (too many values)
	ts, err := csstokens.Tokenize("margin: 1px 2px 3px 4px 5px")
	if err != nil {
		t.Errorf("prep, unexpected error: %v", err)
	}

	_, err = ParseMargin(ts)
	if err == nil {
		t.Error("expected error for too many values, got nil")
	}
}

func TestParseMargin_invalidValue(t *testing.T) {
	// margin: invalid;
	ts, err := csstokens.Tokenize("margin: invalid")
	if err != nil {
		t.Errorf("prep, unexpected error: %v", err)
	}

	_, err = ParseMargin(ts)
	if err == nil {
		t.Error("expected error for invalid value, got nil")
	}
}
