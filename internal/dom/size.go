package dom

import (
	"fmt"
)

// orders self-sizing on dfs post-order
func SizeUp(e Element, opts *Options) error {
	if p, ok := e.(Parent); ok {
		for _, c := range p.GetChildren() {
			if err := SizeUp(c, opts); err != nil {
				return fmt.Errorf("size")
			}
		}
	}

	o, ok := e.(Sizer)
	if !ok {
		return fmt.Errorf("can't size the element type: %T", e)
	}
	if err := o.Size(opts); err != nil {
		return fmt.Errorf("tag sizing: %w", err)
	}

	return nil
}
