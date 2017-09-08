package seatgeek

type Sample struct {
	Input    *Ticket   `json:"input"`
	Expected *Expected `json:"expected"`
}

func NewSample(in *Ticket, exp *Expected) *Sample {
	return &Sample{
		Input:    in,
		Expected: exp,
	}
}
