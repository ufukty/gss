package gsse

import (
	"reflect"
	"testing"

	"go.ufukty.com/gss/internal/ast/ast"
	"go.ufukty.com/gss/internal/dimensional"
)

func compare(a, b any) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}

func TestParse(t *testing.T) {
	type tc struct {
		name     string
		input    string
		expected any
	}

	oneEm := dimensional.New(1, "em")
	oneVw := dimensional.New(1, "vw")
	oneVh := dimensional.New(1, "vh")

	tcs := []tc{
		{"direct", "1", dimensional.New(1.0)},
		{"unit", "1em", dimensional.New(1.0, "em")},
		{"calc", "calc(1em)", ast.Calc{Expr: oneEm}},
		{"addition", "calc(1em + 1em)", ast.Calc{Expr: ast.Addition{Left: oneEm, Right: oneEm}}},
		{"nested", "calc(1em + min(1vw, 1vh))", ast.Calc{Expr: ast.Addition{Left: oneEm, Right: ast.Min{Exprs: []any{oneVw, oneVh}}}}},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Parse(tc.input)
			if err != nil {
				t.Errorf("act, unexpected error: %v", err)
			}
			if !compare(tc.expected, got) {
				t.Errorf("assert, expected %v got %v", tc.expected, got)
			}
		})
	}
}
