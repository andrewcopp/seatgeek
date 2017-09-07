package record

type Pair struct {
	Identifier *int
	Value      string
}

func NewPair(id *int, val string) *Pair {
	return &Pair{
		Identifier: id,
		Value:      val,
	}
}
