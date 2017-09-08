package seatgeek

// Output represents the calculated guess of where a ticket belongs.
type Output struct {
	*Result
}

// NewOutput takes the components of a guess and returns them encapsulated in a
// Output struct.
func NewOutput(sec *int, row *int, val bool) *Output {
	return &Output{
		Result: NewResult(sec, row, val),
	}
}
