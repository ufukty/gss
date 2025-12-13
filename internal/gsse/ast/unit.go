package ast

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"go.ufukty.com/gss/internal/gss/tokens"
)

type Unit map[tokens.Unit]int // eg. px^2/em

func (a Unit) Compare(b Unit) bool {
	if len(a) != len(b) {
		return false
	}
	for unit := range a {
		if a[unit] != b[unit] {
			return false
		}
	}
	return true
}

func (a Unit) Multiply(b Unit) Unit {
	c := maps.Clone(a)
	for u, p := range b {
		c[u] += p
	}
	return c
}

func (a *Unit) clean() {
	for u, p := range *a {
		if p == 0 {
			delete(*a, u)
		}
	}
}

func (a Unit) Divide(b Unit) Unit {
	c := maps.Clone(a)
	for u, p := range b {
		c[u] -= p
	}
	c.clean()
	return c
}

var superscript = []string{"⁰", "¹", "²", "³", "⁴", "⁵", "⁶", "⁷", "⁸", "⁹"}

func positiveDigits(i int) []int {
	if i == 0 {
		return []int{0}
	}
	ds := []int{}
	for ; i > 0; i /= 10 {
		ds = append(ds, i%10)
	}
	slices.Reverse(ds)
	return ds
}

func super(i int) string {
	s := []string{}
	if i < 0 {
		i *= -1
		s = append(s, "⁻")
	}
	for _, d := range positiveDigits(i) {
		s = append(s, superscript[d])
	}
	return strings.Join(s, "")
}

// power in superscript unless ^1
func power(p int) string {
	if p == 1 {
		return ""
	}
	return super(p)
}

func (u Unit) String() string {
	us := []string{}
	for _, c := range slices.Sorted(maps.Keys(u)) {
		us = append(us, fmt.Sprintf("%s%s", c, power(u[c])))
	}
	return strings.Join(us, "·")
}

func Units(us ...tokens.Unit) Unit {
	m := map[tokens.Unit]int{}
	for _, u := range us {
		m[u] += 1
	}
	return m
}
