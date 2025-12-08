package tokens

type Tag string

const (
	Tag_Div  Tag = "div"
	Tag_Span Tag = "span"
	Tag_Img  Tag = "img"
)

type Width string

const (
	Width_Auto       Width = "auto"
	Width_MinContent Width = "min-content"
	Width_MaxContent Width = "max-content"
)

type Height string

const (
	Height_Auto       Height = "auto"
	Height_MinContent Height = "min-content"
	Height_MaxContent Height = "max-content"
)

type Display string

const (
	Display_Auto   Display = "auto"
	Display_Block  Display = "block"
	Display_Inline Display = "inline"
	Display_Grid   Display = "grid"
	Display_Flex   Display = "flex"
)

type Unit string

const (
	Unit_Pt      Unit = "pt"
	Unit_Px      Unit = "px"
	Unit_Em      Unit = "em"
	Unit_Inherit Unit = "inherit"
)

type FontFamily string

const (
	FontFamily_Inherit   FontFamily = "inherit"
	FontFamily_Serif     FontFamily = "serif"
	FontFamily_SansSerif FontFamily = "sans-serif"
)

type Color string

const (
	Color_Inherit Color = "inherit"
)

type BackgroundColor string

const (
	BackgroundColor_Inherit BackgroundColor = "inherit"
)
