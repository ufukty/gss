package dimensional

type Unit string

const (
	Percent Unit = "percent"
)

// Length
const (
	Cm Unit = "cm"
	In Unit = "in" // 96px, 72pt
	Mm Unit = "mm"
	Pc Unit = "pc" // picas
	Pt Unit = "pt"
	Px Unit = "px"
	Q  Unit = "q" // quarter millimeter
)

// Length (relative)
const (
	Em  Unit = "em"
	Rem Unit = "rem"
	Vh  Unit = "vh"
	Vw  Unit = "vw"
)

// Angle
const (
	Deg  Unit = "deg"
	Grad Unit = "grad"
	Rad  Unit = "rad"
	Turn Unit = "turn"
)

// Duration
const (
	Ms Unit = "ms"
	S  Unit = "s"
)

const (
	Hz  Unit = "hz"
	KHz Unit = "khz"
)

const (
	Dpcm Unit = "dpcm" // dots per centimeter
	Dpi  Unit = "dpi"  // dots per inch
	Dppx Unit = "dppx" // dots per pixel
)

const (
	Fr Unit = "fr"
)
