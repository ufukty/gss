package dom

import (
	"cmp"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

type Options struct {
	Width, Height, Density float64
}

type Sizer interface {
	Size(opts *Options) error
}

func (d *Div) Size(opts *Options) error

func (h *Html) Size(opts *Options) error

func (i *Img) Size(opts *Options) error {
	p := cmp.Or(i.Ast.SrcSet[opts.Density], i.Ast.Src)
	if p == "" {
		return fmt.Errorf("deciding correct src: no src or suitable srcset item")
	}

	f, err := os.Open(p)
	if err != nil {
		return fmt.Errorf("opening image: %w", err)
	}
	defer f.Close()

	i2, _, err := image.Decode(f)
	if err != nil {
		return fmt.Errorf("decoding image: %w", err)
	}

	s := Size{
		Width:  float64(i2.Bounds().Dx()),
		Height: float64(i2.Bounds().Dy()),
	}

	i.Min = s
	i.Max = s

	return nil
}

func (s *Span) Size(opts *Options) error

func (t *Text) Size(opts *Options) error

var (
	_ Sizer = (*Div)(nil)
	_ Sizer = (*Html)(nil)
	_ Sizer = (*Img)(nil)
	_ Sizer = (*Span)(nil)
	_ Sizer = (*Text)(nil)
)
