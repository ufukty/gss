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
	_ Occupier = (*Div)(nil)
	_ Occupier = (*Html)(nil)
	_ Occupier = (*Img)(nil)
	_ Occupier = (*Span)(nil)
	_ Occupier = (*Text)(nil)
)
