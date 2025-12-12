package dom

import (
	"go.ufukty.com/gss/internal/filter"
	"go.ufukty.com/gss/internal/gss/ast"
	"go.ufukty.com/gss/internal/gss/defaults"
	"go.ufukty.com/gss/internal/gss/tokens"
)

// reverse [cmp.Or]
func OrReverse[T comparable](ts []T) T {
	var zero T
	for i := len(ts) - 1; i >= 0; i-- {
		if ts[i] != zero {
			return ts[i]
		}
	}
	return zero
}

// reverse [cmp.Or] for slices (not compares elements comparison)
func OrSliceReverse[T any](ts [][]T) []T {
	for i := len(ts) - 1; i >= 0; i-- {
		if len(ts[i]) > 0 {
			return ts[i]
		}
	}
	return nil
}

func m[T, Y any](ts []T, f func(T) Y) []Y {
	ns := make([]Y, len(ts))
	for i := range ts {
		ns[i] = f(ts[i])
	}
	return ns
}

func pick[C comparable](defaults C, applying []*ast.Rule, mapper func(*ast.Styles) C) C {
	ss := make([]C, 0, len(applying)+1)
	ss[0] = defaults
	styles := m(applying, func(r *ast.Rule) *ast.Styles { return r.Styles })
	ss = append(ss, m(styles, mapper)...)
	return OrReverse(ss)
}

func picks[C any](defaults []C, applying []*ast.Rule, mapper func(*ast.Styles) []C) []C {
	ss := make([][]C, 0, len(applying)+1)
	ss[0] = defaults
	styles := m(applying, func(r *ast.Rule) *ast.Styles { return r.Styles })
	ss = append(ss, m(styles, mapper)...)
	return OrSliceReverse(ss)
}

// TODO: presedence
func Styles(e Element, rules []*ast.Rule) *ast.Styles {
	apl := filter.Applying(e.GetAst(), rules)
	def := defaults.For(e.GetAst())

	return &ast.Styles{
		BackgroundColor: pick(def.BackgroundColor, apl, func(s *ast.Styles) tokens.BackgroundColor { return s.BackgroundColor }),
		Color:           pick(def.Color, apl, func(s *ast.Styles) tokens.Color { return s.Color }),
		Display:         pick(def.Display, apl, func(s *ast.Styles) ast.Display { return s.Display }),
		FontFamily:      picks(def.FontFamily, apl, func(s *ast.Styles) []tokens.FontFamily { return s.FontFamily }),
		FontSize:        pick(def.FontSize, apl, func(s *ast.Styles) ast.Size { return s.FontSize }),
		Height:          pick(def.Height, apl, func(s *ast.Styles) tokens.Height { return s.Height }),
		Width:           pick(def.Width, apl, func(s *ast.Styles) tokens.Width { return s.Width }),
	}
}
