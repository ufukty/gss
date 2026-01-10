package is

import (
	"strings"

	"github.com/tdewolff/parse/v2/css"
	"golang.org/x/image/colornames"
)

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

var colorFuncs = map[string]any{
	"rgb":   nil,
	"rgba":  nil,
	"hsl":   nil,
	"hsla":  nil,
	"hwb":   nil,
	"lab":   nil,
	"lch":   nil,
	"oklab": nil,
	"oklch": nil,
	"color": nil,
}

func Color(tok css.Token) bool {
	switch tok.TokenType {

	case css.HashToken:
		return isHexColor(tok.Data)

	case css.IdentToken:
		s := strings.ToLower(string(tok.Data))
		return has(eColorValues, s) || has(colornames.Map, s)

	case css.FunctionToken:
		return has(colorFuncs, strings.ToLower(string(tok.Data)))
	}
	return false
}
