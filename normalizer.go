package seatgeek

// Normalizer is the file the challenge specifically mentioned. Unforunately,
// -er names are reserved for interfaces. To compensate, I made sure to use
// the "Normalizer" name so that the main function would read similar to the
// Java, Python, and Ruby implementations.
type Normalizer interface {
	Normalize(section string, row *string) (*int, *int, bool)
}
