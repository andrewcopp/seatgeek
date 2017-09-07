package record

type Manifest struct {
	Section *Pair
	Row     *Pair
}

func NewManifest(sec, row *Pair) *Manifest {
	return &Manifest{
		Section: sec,
		Row:     row,
	}
}
