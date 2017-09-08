package seatgeek

type Result struct {
	Section *int `json:"section_id"`
	Row     *int `json:"row_id"`
	Valid   bool `json:"valid"`
}

func NewResult(sec *int, row *int, val bool) *Result {
	return &Result{
		Section: sec,
		Row:     row,
		Valid:   val,
	}
}
