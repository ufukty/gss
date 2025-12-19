package primitives

import (
	"image/color"
)

// Value types are to be used in instantiating [Expr] types.
type (
	Color        color.RGBA // eg. #FF0000
	Pixels       float64    // eg. 10px
	Angle        float64    // eg. 360deg
	Milliseconds int        // eg. 1000ms
	Image        string     // eg. url()
	FontSet      []string   // eg. "Helvetica", "Helvetica Neue", sans-serif
)
