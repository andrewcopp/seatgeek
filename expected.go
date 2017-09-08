package seatgeek

type Expected struct {
	*Result
}

func NewExpected(sec *int, row *int, val bool) *Expected {
	return &Expected{
		Result: NewResult(sec, row, val),
	}
}
