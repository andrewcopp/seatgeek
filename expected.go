package seatgeek

// Expected represents the information necessary to check the output of the
// normalizer.
type Expected struct {
	*Result
}

// NewExpected takes a optional section, optional row, and boolean to create
// a Expected struct.
func NewExpected(sec *int, row *int, val bool) *Expected {
	return &Expected{
		Result: NewResult(sec, row, val),
	}
}
