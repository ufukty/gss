package dom

type Child interface {
	Element
	GetParent() Parent
	SetParent(e Parent)
}

func (d Div) GetParent() Parent  { return d.Parent }
func (i Img) GetParent() Parent  { return i.Parent }
func (s Span) GetParent() Parent { return s.Parent }
func (t Text) GetParent() Parent { return t.Parent }

func (d *Div) SetParent(p Parent)  { d.Parent = p }
func (i *Img) SetParent(p Parent)  { i.Parent = p }
func (s *Span) SetParent(p Parent) { s.Parent = p }
func (t *Text) SetParent(p Parent) { t.Parent = p }
