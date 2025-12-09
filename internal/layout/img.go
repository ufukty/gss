package layout

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strconv"
	"strings"

	"go.ufukty.com/gss/internal/gss/ast"
)

func imgPath(img ast.Img, opts *opts) (string, error) {
	if has(img.Attributes, "srcset") {
		set, err := parseSrcSet(img.Attributes["srcset"])
		if err != nil {
			return "", fmt.Errorf("parsing srcset: %w", err)
		}
		if has(set, opts.Density) {
			return set[opts.Density], nil
		}
	}
	if has(img.Attributes, "src") {
		return img.Attributes["src"], nil
	}
	return "", fmt.Errorf("no src or suitable srcset item")
}

// TODO: per tag type
func imgSize(img ast.Img, opts *opts) (*size, error) {
	p, err := imgPath(img, opts)
	if err != nil {
		return nil, fmt.Errorf("picking src value: %w", err)
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
