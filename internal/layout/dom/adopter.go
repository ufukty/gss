package dom

type Adopter interface {
	AppendChild(e Element)
}

func (d *Div) AppendChild(e Element)  { d.Children = append(d.Children, e) }
func (h *Html) AppendChild(e Element) { h.Children = append(h.Children, e) }
func (s *Span) AppendChild(e Element) { s.Children = append(s.Children, e) }
