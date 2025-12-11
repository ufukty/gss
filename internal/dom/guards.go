package dom

var (
	_ Child = (*Div)(nil)
	_ Child = (*Img)(nil)
	_ Child = (*Span)(nil)
	_ Child = (*Text)(nil)
)

var (
	_ Parent = (*Div)(nil)
	_ Parent = (*Html)(nil)
	_ Parent = (*Span)(nil)
)

var (
	_ Sizer = (*Div)(nil)
	_ Sizer = (*Html)(nil)
	_ Sizer = (*Img)(nil)
	_ Sizer = (*Span)(nil)
	_ Sizer = (*Text)(nil)
)
