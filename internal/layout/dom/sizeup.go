package dom

import (
	"cmp"
	"fmt"
	"image"
	"os"
)

type Options struct {
	Width, Height, Density float64
}

type Occupier interface {
	SizeUp(opts *Options) error
}

func (img *Img) SizeUp(opts *Options) error {
	p := cmp.Or(img.Ast.SrcSet[opts.Density], img.Ast.Src)
	if p == "" {
		return fmt.Errorf("deciding correct src: no src or suitable srcset item")
	}

	f, err := os.Open(p)
	if err != nil {
		return fmt.Errorf("opening image: %w", err)
	}
	defer f.Close()

	i, _, err := image.Decode(f)
	if err != nil {
		return fmt.Errorf("decoding image: %w", err)
	}

	s := Size{
		Width:  float64(i.Bounds().Dx()),
		Height: float64(i.Bounds().Dy()),
	}

	img.Min = s
	img.Max = s

	return nil
}
