package seatgeek

import (
	"strconv"
)

// Input represents all of the conents of an Input file specified through
// command line flags. These contents are stored in Sample structs.
type Input struct {
	Samples []*Sample
}

// NewInput returns an empty Input struct.
func NewInput() *Input {
	return &Input{
		Samples: []*Sample{},
	}
}

// Load takes a filepath to a CSV and parses the contents in the Input struct.
// A returned error means that the CSV file was not in the correct format.
func (i *Input) Load(path string) error {
	records := Read(path)

	samples := make([]*Sample, len(records)-1)
	for idx, record := range records[1:] {
		valid, err := strconv.ParseBool(record[4])
		if err != nil {
			return err
		}

		var sec *int
		var row *int

		if record[2] != "" {
			s, err := strconv.Atoi(record[2])
			if err != nil {
				return err
			}
			sec = &s
		}

		if record[3] != "" {
			r, err := strconv.Atoi(record[3])
			if err != nil {
				return err
			}
			row = &r
		}

		in := NewTicket(record[0], record[1])
		exp := NewExpected(sec, row, valid)
		samples[idx] = NewSample(in, exp)
	}

	i.Samples = samples

	return nil
}
