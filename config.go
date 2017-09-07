package seatgeek

import "flag"

type Config struct {
	Manifest string
	Input    *string
	Section  *int
	Row      *int
}

func NewConfig() *Config {
	manifest := flag.String("manifest", "", "path to manifest file")
	input := flag.String("input", "", "path to input file")
	section := flag.Int("section", -1, "section input (for testing)")
	row := flag.Int("row", -1, "row input (for testing)")

	flag.Parse()

	config := &Config{}

	config.Manifest = *manifest

	if *input != "" {
		config.Input = input
	}

	if *section != -1 {
		config.Section = section
	}

	if *row != -1 {
		config.Row = row
	}

	return config
}
