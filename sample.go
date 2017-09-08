package seatgeek

// Sample represents a test of what information is supplied and what the
// develoepr is expected to be able to translate it to.
type Sample struct {
	Input    *Ticket   `json:"input"`
	Expected *Expected `json:"expected"`
}

// NewSample returns a Sample struct with the given components.
func NewSample(in *Ticket, exp *Expected) *Sample {
	return &Sample{
		Input:    in,
		Expected: exp,
	}
}
