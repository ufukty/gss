package gss

import (
	"fmt"
	"io"
	"iter"
	"strings"

	"github.com/tdewolff/parse/v2"
	"github.com/tdewolff/parse/v2/css"
	"go.ufukty.com/gss/internal/ast"
)

// FIXME:
func writeDecl(dst *ast.Rule, prop string, values []css.Token) error {
	switch prop {
	case "align-items":
	case "background-color":
		dst.Styles.BackgroundColor = values
	case "border-radius":
		dst.Styles.BorderRadiuses.TopLeft = values
	case "box-sizing":
	case "color":
		dst.Styles.Text.Color = values
	case "color-scheme":
	case "display":
	case "font-family":
	case "font-size":
	case "font-weight":
	case "grid-column":
	case "grid-row":
	case "grid-template-columns":
	case "height":
		dst.Styles.Dimensions.Height = serializeSelector(values)
	case "justify-content":
	case "margin":
		m, err := ParseMargin(values)
		if err != nil {
			return fmt.Errorf("parsing margin: %w", err)
		}
		dst.Styles.Margin = m
	case "padding":
		p, err := ParsePadding(values)
		if err != nil {
			return fmt.Errorf("parsing padding: %w", err)
		}
		dst.Styles.Padding = p
	case "position":
	case "text-align":
		dst.Styles.Text.TextAlignment = values
	case "text-decoration":
	case "width":
		dst.Styles.Dimensions.Width = values
	default:
		return fmt.Errorf("unknown property")
	}
	return nil
}

func serializeSelector(values []css.Token) string {
	ss := []string{}
	for _, v := range values {
		ss = append(ss, v.String())
	}
	return strings.Join(ss, "")
}

type event struct {
	GrammerType css.GrammarType
	TokenType   css.TokenType
	Data        string
	Values      []css.Token
}

func events(p *css.Parser) iter.Seq[event] {
	return func(yield func(event) bool) {
		for {
			g, t, d := p.Next()
			if g == css.ErrorGrammar {
				return
			}
			e := event{
				GrammerType: g,
				TokenType:   t,
				Data:        string(d),
				Values:      p.Values(),
			}
			if !yield(e) {
				return
			}
		}
	}
}

func Parse(r io.Reader) (*ast.Gss, error) {
	p := css.NewParser(parse.NewInput(r), false)
	a := &ast.Gss{
		Rules: []*ast.Rule{},
	}
	var cur *ast.Rule
	for e := range events(p) {
		switch e.GrammerType {

		case css.BeginRulesetGrammar:
			if cur != nil {
				return nil, fmt.Errorf("nesting is not supported")
			}
			cur = &ast.Rule{
				Selector: serializeSelector(e.Values),
				Styles:   &ast.Styles{},
			}

		case css.EndRulesetGrammar:
			a.Rules = append(a.Rules, cur)
			cur = nil

		case css.DeclarationGrammar:
			if cur != nil {
				err := writeDecl(cur, e.Data, e.Values)
				if err != nil {
					return nil, fmt.Errorf("writing property %q: %w", e.Data, err)
				}
			}
		}
	}
	if err := p.Err(); err != nil && err != io.EOF {
		return nil, fmt.Errorf("parser: %w", err)
	}
	return a, nil
}
