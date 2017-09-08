package seatgeek

type Output struct {
	*Result
}

func NewOutput(sec *int, row *int, val bool) *Output {
	return &Output{
		Result: NewResult(sec, row, val),
	}
}
