package gss

import (
	"fmt"
	"strconv"
	"strings"

	"go.ufukty.com/gss/internal/ast/gsse"
)

func ParseColor(s string) (gsse.Color, error) {
	if s[0] != '#' {
		return gsse.Color{}, fmt.Errorf("expected leading # character")
	}
	strings.ToLower(s)
	switch len(s) {
	case 3:
		s = strings.Repeat(s[0:1], 2) + strings.Repeat(s[1:2], 2) + strings.Repeat(s[2:3], 2) + "ff"
	case 4:
		s = strings.Repeat(s[0:1], 2) + strings.Repeat(s[1:2], 2) + strings.Repeat(s[2:3], 2) + strings.Repeat(s[3:4], 2)
	case 6:
		s = s + "ff"
	case 8:
		break
	default:
		return gsse.Color{}, fmt.Errorf("unexpected length")
	}
	u, err := strconv.ParseUint(s, 16, 32)
	if err != nil {
		return gsse.Color{}, err
	}
	return gsse.Color{
		R: uint8(u >> 24),
		G: uint8(u >> 16),
		B: uint8(u >> 8),
		A: uint8(u),
	}, nil
}
