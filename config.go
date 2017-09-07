package seatgeek

import "flag"

type Config struct {
	Manifest string
	Input    *string
	Section  *string
	Row      *string
}

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
