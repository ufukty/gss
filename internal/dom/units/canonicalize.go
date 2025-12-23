package units

import (
	"fmt"
	"math"
)

const (
	PtPerIn = 72.0
	PxPerIn = 96.0
	CmPerIn = 2.54
	PxPerCm = PxPerIn / CmPerIn
	PxPerPt = PxPerIn / PtPerIn
)

var factors = map[Unit]float64{
	// to Pixels
	Px: 1.0,
	In: PxPerIn,
	Cm: PxPerCm,
	Mm: PxPerCm / 10.0,
	Q:  PxPerCm / 40.0,
	Pt: PxPerPt,
	Pc: PxPerPt * 12.0,

	// to Degrees
	Deg:  1.0,
	Rad:  180.0 / math.Pi,
	Grad: 0.9,
	Turn: 360.0,

	// to Seconds
	S:  1.0,
	Ms: 0.001,

	// to Hz
	Hz:  1.0,
	KHz: 1000,

	// to Dppx
	Dppx: 1.0,
	Dpi:  1.0 / PxPerIn,
	Dpcm: 1.0 / PxPerCm,
}

var contextual = map[Unit]any{
	Fr:      nil,
	Percent: nil,
	Em:      nil,
	Rem:     nil,
	Vw:      nil,
	Vh:      nil,
}

func factor(u Unit) (float64, error) {
	if f, ok := factors[u]; ok {
		return f, nil
	}
	if _, ok := contextual[u]; ok {
		return 0.0, ErrCannonicalizingContextualUnit
	}
	return 0.0, ErrUnknownUnit
}

var toCore = map[Unit]Unit{
	Cm: Px,
	In: Px,
	Mm: Px,
	Pc: Px,
	Pt: Px,
	Px: Px,
	Q:  Px,

	Deg:  Deg,
	Grad: Deg,
	Rad:  Deg,
	Turn: Deg,

	Ms: S,
	S:  S,

	Hz:  Hz,
	KHz: Hz,

	Dpcm: Dppx,
	Dpi:  Dppx,
	Dppx: Dppx,
}

func Canonicalize(d Dimension) (Dimension, error) {
	n := Dimension{Value: d.Value, Unit: make(Complex)}
	for u, p := range d.Unit {
		if p == 0 {
			continue
		}
		u2, ok := toCore[u]
		if !ok {
			return n, fmt.Errorf("checking the core unit for %s: %w", u, ErrUnknownUnit)
		}
		n.Unit = n.Unit.Multiply(map[Unit]int{u2: p})
		f, err := factor(u)
		if err != nil {
			return n, fmt.Errorf("canonicalizing %s to core unit: %w", u, err)
		}
		n.Value *= math.Pow(f, float64(p))
	}
	return n, nil
}
