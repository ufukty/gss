package dom

// TODO: update when Go allows interface defining on common fields
type Parent interface {
	Element
	AppendChild(e Child)
	GetChildren() []Child
}

func (d *Div) AppendChild(e Child)  { d.Children = append(d.Children, e) }
func (h *Html) AppendChild(e Child) { h.Children = append(h.Children, e) }
func (s *Span) AppendChild(e Child) { s.Children = append(s.Children, e) }

func (d Div) GetChildren() []Child  { return d.Children }
func (h Html) GetChildren() []Child { return h.Children }
func (s Span) GetChildren() []Child { return s.Children }
