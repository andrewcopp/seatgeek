package seatgeek

// Result represents a prediction. It is utlized on by both sides of the
// equation (Output and Expected).
type Result struct {
	Section *int `json:"section_id"`
	Row     *int `json:"row_id"`
	Valid   bool `json:"valid"`
}

// NewResult takes the components of a guess and returns them in an
// encapsulated struct.
func NewResult(sec *int, row *int, val bool) *Result {
	return &Result{
		Section: sec,
		Row:     row,
		Valid:   val,
	}
}
