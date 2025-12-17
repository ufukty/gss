package gss

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

type DisplayOutside string

const (
	DisplayOutside_Block  DisplayOutside = "block"
	DisplayOutside_Inline DisplayOutside = "inline"
)

type DisplayInside string

const (
	DisplayInside_Flex DisplayInside = "flex"
	DisplayInside_Flow DisplayInside = "flow"
	DisplayInside_Grid DisplayInside = "grid"
)

type Unit string

const (
	Unit_Em      Unit = "em"
	Unit_Inherit Unit = "inherit"
	Unit_Pc      Unit = "%"
	Unit_Pt      Unit = "pt"
	Unit_Px      Unit = "px"
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

type TextAlignment string

const (
	TextAlignment_Inherit TextAlignment = "inherit"
	TextAlignment_Left    TextAlignment = "left"
	TextAlignment_Center  TextAlignment = "center"
	TextAlignment_Right   TextAlignment = "right"
)
