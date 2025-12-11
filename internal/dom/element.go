package dom

type Element interface {
	element()
}

func (*Div) element()  {}
func (*Html) element() {}
func (*Img) element()  {}
func (*Span) element() {}
func (*Text) element() {}
