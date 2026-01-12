package ast

import (
	"maps"
	"slices"

	"go.ufukty.com/gommons/pkg/tree"
)

// Gss
type (
	Display struct {
		Outside string
		Inside  string
	}
	Border struct {
		Color any // "inherit", "transparent", color.NRGBA
		Style any // "inherit", gss.BorderStyle
		Width any // "none", dimensional.Dimension
	}
	BorderRadiuses struct {
		TopLeft, TopRight, BottomRight, BottomLeft any // "none", "inherit", dimensional.Dimension
	}
	Borders struct {
		Top, Right, Bottom, Left Border
	}
	Margin struct {
		Top, Right, Bottom, Left any // "inherit", dimensional.Dimension
	}
	Padding struct {
		Top, Right, Bottom, Left any // "inherit", dimensional.Dimension
	}
	Font struct {
		Family any // "inherit", []string
		Size   any // "inherit", dimensional.Dimension
		Weight any // "inherit", int
	}
	Text struct {
		Color         any // "inherit", "transparent", color.NRGBA
		LineHeight    any // "inherit", dimension.Dimensional
		TextAlignment any // "inherit", gss.TextAlignment
	}
	Dimensions struct {
		Height any // "auto", "min-content", "max-content", dimensional.Dimension
		Width  any // "auto", "min-content", "max-content", dimensional.Dimension
	}
	// TODO: handle shorthand syntaxes during parsing
	Styles struct {
		Dimensions      Dimensions
		Margin          Margin
		Padding         Padding
		Display         Display
		Text            Text
		Font            Font
		Border          Borders
		BorderRadiuses  BorderRadiuses
		BackgroundColor any // "inherit", "transparent", color.NRGBA
	}
	Rule struct {
		Selector string
		Styles   *Styles
	}
	Gss struct {
		Rules []*Rule
	}
)

func isZero[T comparable](t T) bool {
	var z T
	return t == z
}

func collect(s map[string]any) []string {
	ss := []string{}
	for _, k := range slices.Sorted(maps.Keys(s)) {
		if !isZero(s[k]) {
			switch v := s[k].(type) {
			case string:
				ss = append(ss, tree.List(k, []string{v}))
			case interface{ String() string }:
				ss = append(ss, tree.List(k, []string{v.String()}))
			case interface{ Strings() []string }:
				ss = append(ss, tree.List(k, v.Strings()))
			default:
				ss = append(ss, tree.List(k, []string{"value of unknown type"}))
			}
		}
	}
	if len(ss) > 0 {
		return ss
	}
	return nil
}

func (s Display) Strings() []string {
	return collect(map[string]any{
		"Outside": s.Outside,
		"Inside":  s.Inside,
	})
}

func (s Border) Strings() []string {
	return collect(map[string]any{
		"Color":     s.Color,
		"Style":     s.Style,
		"Thickness": s.Width,
	})
}

func (s BorderRadiuses) Strings() []string {
	return collect(map[string]any{
		"TopLeft":     s.TopLeft,
		"TopRight":    s.TopRight,
		"BottomRight": s.BottomRight,
		"BottomLeft":  s.BottomLeft,
	})
}

func (s Borders) Strings() []string {
	return collect(map[string]any{
		"Top":    s.Top,
		"Right":  s.Right,
		"Bottom": s.Bottom,
		"Left":   s.Left,
	})
}

func (s Margin) Strings() []string {
	return collect(map[string]any{
		"Top":    s.Top,
		"Right":  s.Right,
		"Bottom": s.Bottom,
		"Left":   s.Left,
	})
}

func (s Padding) Strings() []string {
	return collect(map[string]any{
		"Top":    s.Top,
		"Right":  s.Right,
		"Bottom": s.Bottom,
		"Left":   s.Left,
	})
}

func (s Font) Strings() []string {
	return collect(map[string]any{
		"Family": s.Family,
		"Size":   s.Size,
		"Weight": s.Weight,
	})
}

func (s Text) Strings() []string {
	return collect(map[string]any{
		"Color":         s.Color,
		"LineHeight":    s.LineHeight,
		"TextAlignment": s.TextAlignment,
	})
}

func (s Dimensions) Strings() []string {
	return collect(map[string]any{
		"Height": s.Height,
		"Width":  s.Width,
	})
}

func (s Styles) Strings() []string {
	return collect(map[string]any{
		"Dimensions":      s.Dimensions,
		"Margin":          s.Margin,
		"Padding":         s.Padding,
		"Display":         s.Display,
		"Text":            s.Text,
		"Font":            s.Font,
		"Border":          s.Border,
		"BorderRadiuses":  s.BorderRadiuses,
		"BackgroundColor": s.BackgroundColor,
	})
}

func (r Rule) String() string {
	return tree.List(r.Selector, r.Styles.Strings())
}
