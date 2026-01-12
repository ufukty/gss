package core

import "github.com/tdewolff/parse/v2/css"

// Trim the leading and trailing spaces.
func trimSpaces(ts []css.Token) []css.Token {
	i, j := 0, len(ts)-1
	for ; ts[i].TokenType == css.WhitespaceToken; i++ {
	}
	for ; ts[j].TokenType == css.WhitespaceToken; j-- {
	}
	return ts[i : j+1]
}

func has[K comparable, V any](m map[K]V, k K) bool {
	_, ok := m[k]
	return ok
}
