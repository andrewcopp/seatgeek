package seatgeek

import (
	"strconv"

	"github.com/andrewcopp/seatgeek/record"
)

type Input struct {
	Samples []*Sample
}

func NewInput() *Input {
	return &Input{
		Samples: []*Sample{},
	}
}

func (i *Input) Load(path string) error {
	records := record.Read(path)

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

		in := NewInput2(record[0], record[1])
		exp := NewExpected(sec, row, valid)
		samples[idx] = NewSample(in, exp)
	}

	i.Samples = samples

	return nil
}
