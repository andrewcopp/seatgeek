package record

type Input struct {
	Section *Pair
	Row     *Pair
	Valid   bool
}

func NewInput(sec, row *Pair, valid bool) *Input {
	return &Input{
		Section: sec,
		Row:     row,
		Valid:   valid,
	}
}
