package gsse

import "go.ufukty.com/gss/internal/tokens/gss"

var (
	_ Expr[gss.Color] = (*LightDark[gss.Color])(nil)
)
