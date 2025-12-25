package gss

import (
	"regexp"
	"strconv"
	"strings"

	"go.ufukty.com/gss/internal/ast/gss"
	"go.ufukty.com/gss/internal/dimensional"
)

var (
	regexBorderThickness = regexp.MustCompile(`([0-9]+)(|px|pt|em|rem|vw|vh)`)
	regexBorderColor     = regexp.MustCompile(`#(?:[0-9A-Fa-f][{3,4}|[0-9A-Fa-f]{6}|[0-9A-Fa-f]{8})`)
	regexBorderStyle     = regexp.MustCompile(`(solid|dashed|dotted)`)
)

func ParseBorder(s string) gss.Border {
	b := gss.Border{
		Color:     "#000000",
		Style:     "solid",
		Thickness: "none",
	}
	for seq := range strings.SplitSeq(s, " ") {
		if ms := regexBorderThickness.FindStringSubmatch(seq); len(ms) > 0 {
			b.Thickness = dimensional.New(silent(strconv.ParseFloat(ms[1], 64)), dimensional.Unit(ms[2]))
		}
		if m := regexBorderColor.FindString(seq); m != "" {
			b.Color = silent(ParseColor(m))
		}
		if m := regexBorderStyle.FindString(seq); m != "" {
			b.Style = m
		}
	}
	return b
}

func ParseBorders(s string) gss.Borders {
	left, right, top, bottom := "", "", "", ""
	switch ss := strings.Split(s, ","); len(ss) {
	case 1:
		top, right, bottom, left = ss[0], ss[0], ss[0], ss[0]
	case 2:
		top, bottom, left, right = ss[0], ss[0], ss[1], ss[1]
	case 3:
		top, left, right, bottom = ss[0], ss[1], ss[1], ss[2]
	case 4:
		top, right, bottom, left = ss[0], ss[1], ss[2], ss[3]
	}
	return gss.Borders{
		Top:    ParseBorder(top),
		Right:  ParseBorder(right),
		Bottom: ParseBorder(bottom),
		Left:   ParseBorder(left),
	}
}
