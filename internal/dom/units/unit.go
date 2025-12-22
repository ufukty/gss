package units

type Unit string

const (
	Percent Unit = "percent"
)

// Length
const (
	Cm  Unit = "cm"
	Em  Unit = "em"
	In  Unit = "in" // 96px
	Mm  Unit = "mm"
	Pc  Unit = "pc"
	Pt  Unit = "pt"
	Px  Unit = "px"
	Q   Unit = "q" // 0.25mm
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
	Dpcm Unit = "dpcm"
	Dpi  Unit = "dpi"
	Dppx Unit = "dppx"
)

const (
	Fr Unit = "fr"
)
