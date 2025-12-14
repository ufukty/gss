package html

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"go.ufukty.com/gss/internal/ast/html"
)

func parseSrcSet(s string) (map[float64]string, error) {
	set := map[float64]string{}
	for s := range strings.SplitSeq(s, ",") {
		s = strings.TrimSpace(s)
		ps := strings.Split(s, " ")
		if len(ps) == 2 {
			d, err := strconv.ParseFloat(ps[1], 64)
			if err != nil {
				return nil, fmt.Errorf("parsing density number: %w", err)
			}
			set[d] = strings.TrimSpace(ps[0])
		}
	}
	return set, nil
}

func Html(src io.Reader) *html.Html {
	h := &html.Html{}

	return h
}
