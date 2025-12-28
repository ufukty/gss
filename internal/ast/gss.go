package ast

import (
	"go.ufukty.com/gommons/pkg/tree"
)

// Gss
type (
	Display struct {
		Outside string
		Inside  string
	}
	Border struct {
		Color     any // "inherit", "transparent", color.NRGBA
		Style     any // "inherit", gss.BorderStyle
		Thickness any // "none", dimensional.Dimension
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

func appendString(ss []string, prop string, value any) []string {
	if !isZero(value) {
		switch s := value.(type) {
		case string:
			ss = append(ss, tree.List(prop, []string{s}))
		case interface{ String() string }:
			ss = append(ss, tree.List(prop, []string{s.String()}))
		case interface{ Strings() []string }:
			ss = append(ss, tree.List(prop, s.Strings()))
		default:
			ss = append(ss, tree.List(prop, []string{"value of unknown type"}))
		}
	}
	return ss
}

func (s Display) Strings() []string {
	ss := []string{}
	ss = appendString(ss, "Outside", s.Outside)
	ss = appendString(ss, "Inside", s.Inside)
	return ss
}

func (s Border) Strings() []string {
	ss := []string{}
	ss = appendString(ss, "Color", s.Color)
	ss = appendString(ss, "Style", s.Style)
	ss = appendString(ss, "Thickness", s.Thickness)
	return ss
}

func (s BorderRadiuses) Strings() []string {
	ss := []string{}
	ss = appendString(ss, "TopLeft", s.TopLeft)
	ss = appendString(ss, "TopRight", s.TopRight)
	ss = appendString(ss, "BottomRight", s.BottomRight)
	ss = appendString(ss, "BottomLeft", s.BottomLeft)
	return ss
}

func (s Borders) Strings() []string {
	ss := []string{}
	ss = appendString(ss, "Top", s.Top)
	ss = appendString(ss, "Right", s.Right)
	ss = appendString(ss, "Bottom", s.Bottom)
	ss = appendString(ss, "Left", s.Left)
	return ss
}

func (s Margin) Strings() []string {
	ss := []string{}
	ss = appendString(ss, "Top", s.Top)
	ss = appendString(ss, "Right", s.Right)
	ss = appendString(ss, "Bottom", s.Bottom)
	ss = appendString(ss, "Left", s.Left)
	return ss
}

func (s Padding) Strings() []string {
	ss := []string{}
	ss = appendString(ss, "Top", s.Top)
	ss = appendString(ss, "Right", s.Right)
	ss = appendString(ss, "Bottom", s.Bottom)
	ss = appendString(ss, "Left", s.Left)
	return ss
}

func (s Font) Strings() []string {
	ss := []string{}
	ss = appendString(ss, "Family", s.Family)
	ss = appendString(ss, "Size", s.Size)
	ss = appendString(ss, "Weight", s.Weight)
	return ss
}

func (s Text) Strings() []string {
	ss := []string{}
	ss = appendString(ss, "Color", s.Color)
	ss = appendString(ss, "LineHeight", s.LineHeight)
	ss = appendString(ss, "TextAlignment", s.TextAlignment)
	return ss
}

func (s Dimensions) Strings() []string {
	ss := []string{}
	ss = appendString(ss, "Height", s.Height)
	ss = appendString(ss, "Width", s.Width)
	return ss
}

func (s Styles) Strings() []string {
	ss := []string{}
	ss = appendString(ss, "Dimensions", s.Dimensions)
	ss = appendString(ss, "Margin", s.Margin)
	ss = appendString(ss, "Padding", s.Padding)
	ss = appendString(ss, "Display", s.Display)
	ss = appendString(ss, "Text", s.Text)
	ss = appendString(ss, "Font", s.Font)
	ss = appendString(ss, "Border", s.Border)
	ss = appendString(ss, "BorderRadiuses", s.BorderRadiuses)
	ss = appendString(ss, "BackgroundColor", s.BackgroundColor)
	return ss
}

func (r Rule) String() string {
	return tree.List(r.Selector, r.Styles.Strings())
}
