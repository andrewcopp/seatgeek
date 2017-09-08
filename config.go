package seatgeek

import "flag"

// Config contains all possible values that could be passed to the executable
// through command line flags.
type Config struct {
	Manifest string
	Input    *string
	Section  *string
	Row      *string
}

// NewConfig creates a Config struct that is populated by the values returned
// from parsing the executable flags. Flags that return zero values are treated
// as nil.
func NewConfig() *Config {
	manifest := flag.String("manifest", "", "path to manifest file")
	input := flag.String("input", "", "path to input file")
	section := flag.String("section", "", "section input (for testing)")
	row := flag.String("row", "", "row input (for testing)")

	flag.Parse()

	config := &Config{}

	config.Manifest = *manifest

	if *input != "" {
		config.Input = input
	}

	if *section != "" {
		config.Section = section
	}

	config.Row = row

	return config
}
