package html

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var regexSrcSetItem = regexp.MustCompile(`\s*([^ ]+)\s+([0-9]*)x\s*`)

func parseSrcSet(s string) (map[float64]string, error) {
	set := map[float64]string{}
	for s := range strings.SplitSeq(s, ",") {
		ms := regexSrcSetItem.FindStringSubmatch(s)
		if len(ms) == 3 {
			d, err := strconv.ParseFloat(ms[2], 64)
			if err != nil {
				return nil, fmt.Errorf("parsing density number: %w", err)
			}
			set[d] = strings.TrimSpace(ms[1])
		}
	}
	return set, nil
}
