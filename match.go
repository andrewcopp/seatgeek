package seatgeek

type Match struct {
	*Sample
	Output *Output `json:"output"`
}

func NewMatch(smp *Sample, out *Output) *Match {
	return &Match{
		Sample: smp,
		Output: out,
	}
}
