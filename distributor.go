package seatgeek

type Distributor struct {
	Tests []*Test
}

func NewDistributor(i *Input) *Distributor {
	tests := make([]*Test, len(i.Records))
	for idx, rec := range i.Records {
		tkt := NewTicket(rec.Section.Value, &rec.Row.Value)
		out := NewOutput(rec.Section.Identifier, rec.Row.Identifier, rec.Valid)
		tests[idx] = NewTest(tkt, out)
	}

	return &Distributor{
		Tests: tests,
	}
}
