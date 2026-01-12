package core

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"

	"github.com/tdewolff/parse/v2/css"
	"golang.org/x/image/colornames"
)

func has[K comparable, V any](m map[K]V, k K) bool {
	_, ok := m[k]
	return ok
}

func isHex(r rune) bool {
	return ('0' <= r && r <= '9') || ('a' <= r && r <= 'f') || ('A' <= r && r <= 'F')
}

func areAllHex(rs []rune) bool {
	for _, r := range rs {
		if !isHex(r) {
			return false
		}
	}
	return true
}

func isHexColor(b []byte) bool {
	switch len(b) {
	case 3, 4, 6, 8:
		return areAllHex([]rune(string(b)))
	}
	return false
}

var eColorValues = map[string]any{
	"transparent":  nil,
	"currentcolor": nil,
}

func IsColor(tok css.Token) bool {
	switch tok.TokenType {

	case css.HashToken:
		return isHexColor(tok.Data[1:])

	case css.IdentToken:
		s := strings.ToLower(string(tok.Data))
		return has(eColorValues, s) || has(colornames.Map, s)
	}
	return false
}

func parseHexColor(bs []byte) (color.RGBA, error) {
	s := strings.ToLower(string(bs))
	switch len(s) {
	case 3:
		r, g, b := s[0:1], s[1:2], s[2:3]
		s = r + r + g + g + b + b + "ff"
	case 4:
		r, g, b, a := s[0:1], s[1:2], s[2:3], s[3:4]
		s = r + r + g + g + b + b + a + a
	case 6:
		s = s + "ff"
	case 8:
		break
	default:
		return color.RGBA{}, fmt.Errorf("unexpected length")
	}
	u, err := strconv.ParseUint(s, 16, 32)
	if err != nil {
		return color.RGBA{}, err
	}
	nrgba := color.NRGBA{
		R: uint8(u >> 24),
		G: uint8(u >> 16),
		B: uint8(u >> 8),
		A: uint8(u),
	}
	rgba, ok := color.RGBAModel.Convert(nrgba).(color.RGBA)
	if !ok {
		return color.RGBA{}, fmt.Errorf("multiplying by alpha is failed")
	}
	return rgba, nil
}

// Parses the color to color.RGBA or string
func ParseColor(tok css.Token) (any, error) {
	switch tok.TokenType {

	case css.HashToken:
		c, err := parseHexColor(tok.Data[1:])
		if err != nil {
			return nil, fmt.Errorf("hex: %w", err)
		}
		return c, nil

	case css.IdentToken:
		s := strings.ToLower(string(tok.Data))
		switch {
		case has(colornames.Map, s):
			return colornames.Map[s], nil
		case s == "transparent":
			return color.RGBA{0, 0, 0, 0}, nil
		case s == "currentcolor":
			return "currentcolor", nil
		}
		return nil, fmt.Errorf("unknown value: %s", string(tok.Data))
	}

	return nil, fmt.Errorf("unknown representation: %T", tok)
}
