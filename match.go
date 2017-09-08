package seatgeek

type Match struct {
	*Sample
	Output *Output2 `json:"output"`
}

func NewMatch(smp *Sample, out *Output2) *Match {
	return &Match{
		Sample: smp,
		Output: out,
	}
}

type Sample struct {
	Input    *Input2   `json:"input"`
	Expected *Expected `json:"expected"`
}

func NewSample(in *Input2, exp *Expected) *Sample {
	return &Sample{
		Input:    in,
		Expected: exp,
	}
}

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

type Output2 struct {
	*Result
}

func NewOutput2(sec *int, row *int, val bool) *Output2 {
	return &Output2{
		Result: NewResult(sec, row, val),
	}
}

type Input2 struct {
	Section string `json:"section"`
	Row     string `json:"row"`
}

func NewInput2(sec string, row string) *Input2 {
	return &Input2{
		Section: sec,
		Row:     row,
	}
}

type Expected struct {
	*Result
}

func NewExpected(sec *int, row *int, val bool) *Expected {
	return &Expected{
		Result: NewResult(sec, row, val),
	}
}