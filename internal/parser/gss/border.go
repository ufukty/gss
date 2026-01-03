package gss

import (
	"iter"
	"regexp"
	"slices"
	"strconv"

	"github.com/tdewolff/parse/v2/css"
	"go.ufukty.com/gss/internal/ast"
	"go.ufukty.com/gss/internal/dimensional"
)

var (
	regexBorderThickness = regexp.MustCompile(`([0-9]+)(|px|pt|em|rem|vw|vh)`)
	regexBorderColor     = regexp.MustCompile(`#(?:[0-9A-Fa-f][{3,4}|[0-9A-Fa-f]{6}|[0-9A-Fa-f]{8})`)
	regexBorderStyle     = regexp.MustCompile(`(solid|dashed|dotted)`)
)

func silent[T any](t T, _ error) T {
	return t
}

func ParseBorder(ts []css.Token) ast.Border {
	b := ast.Border{
		Color:     "#000000",
		Style:     "solid",
		Thickness: "none",
	}
	for _, t := range ts {
		if ms := regexBorderThickness.FindStringSubmatch(t); len(ms) > 0 {
			b.Thickness = dimensional.New(silent(strconv.ParseFloat(ms[1], 64)), dimensional.Unit(ms[2]))
		}
		if m := regexBorderColor.FindString(t); m != "" {
			b.Color = silent(ParseColor(m))
		}
		if m := regexBorderStyle.FindString(t); m != "" {
			b.Style = m
		}
	}
	return b
}

func split(ts []css.Token, sep css.TokenType) iter.Seq[[]css.Token] {
	return func(yield func([]css.Token) bool) {
		prev := 0
		for cur, t := range ts {
			if t.TokenType == sep {
				if !yield(ts[prev:cur]) {
					return
				}
				prev = cur + 1
			}
		}
	}
}

func ParseBorders(ts []css.Token) ast.Borders {
	ss := slices.Collect(split(ts, css.CommaToken))
	var left, right, top, bottom []css.Token
	switch len(ss) {
	case 1:
		top, right, bottom, left = ss[0], ss[0], ss[0], ss[0]
	case 2:
		top, bottom, left, right = ss[0], ss[0], ss[1], ss[1]
	case 3:
		top, left, right, bottom = ss[0], ss[1], ss[1], ss[2]
	case 4:
		top, right, bottom, left = ss[0], ss[1], ss[2], ss[3]
	}
	return ast.Borders{
		Top:    ParseBorder(top),
		Right:  ParseBorder(right),
		Bottom: ParseBorder(bottom),
		Left:   ParseBorder(left),
	}
}
