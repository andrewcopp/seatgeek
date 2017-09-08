package seatgeek

// Match represents the pairing of a given sample and the calulcated output for
// comparison.
type Match struct {
	*Sample
	Output *Output `json:"output"`
}

// NewMatch returns the sample and output tied together in a Match struct.
func NewMatch(smp *Sample, out *Output) *Match {
	return &Match{
		Sample: smp,
		Output: out,
	}
}
