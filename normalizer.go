package seatgeek

type Normalizer interface {
	Normalize(section string, row *string) (*int, *int, bool)
}

type Default struct {
	Checker *Checker
}

func NewDefault(checker *Checker) *Default {
	return &Default{
		Checker: checker,
	}
}

func (d *Default) Normalize(section string, row *string) (*int, *int, bool) {
	tkt := NewTicket(section, row)
	out := (*d.Checker).Check(tkt)
	return out.Section, out.Row, out.Valid
}
