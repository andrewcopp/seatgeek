package seatgeek

type Normalizer interface {
	Normalize(section string, row *string) (*int, *int, bool)
}
