package layout

import (
	"cmp"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"go.ufukty.com/gss/internal/gss/ast"
)

// TODO: per tag type
func imgSize(img *ast.Img, opts *opts) (*size, error) {
	p := cmp.Or(img.SrcSet[opts.Density], img.Src)
	if p == "" {
		return nil, fmt.Errorf("deciding correct src: no src or suitable srcset item")
	}

	f, err := os.Open(p)
	if err != nil {
		return nil, fmt.Errorf("opening image: %w", err)
	}
	defer f.Close()

	i, _, err := image.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("decoding image: %w", err)
	}

	return &size{
		width:  float64(i.Bounds().Dx()),
		height: float64(i.Bounds().Dy()),
	}, nil
}
