package gss

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"
)

func ParseColor(s string) (any, error) {
	if s[0] != '#' {
		return nil, fmt.Errorf("expected leading # character")
	}
	s = strings.ToLower(s)
	s = s[1:] // hash
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
		return nil, fmt.Errorf("unexpected length")
	}
	u, err := strconv.ParseUint(s, 16, 32)
	if err != nil {
		return nil, err
	}
	return color.NRGBA{
		R: uint8(u >> 24),
		G: uint8(u >> 16),
		B: uint8(u >> 8),
		A: uint8(u),
	}, nil
}
