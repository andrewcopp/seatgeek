package seatgeek

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/andrewcopp/seatgeek/record"
)

type Input struct {
	Records []*record.Input
}

func NewInput() *Input {
	return &Input{
		Records: []*record.Input{},
	}
}

func (i *Input) Load(path string) error {
	// "~/Developer/sectionnorm/samples/metstest.csv"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err.Error())
	}

	var reader io.Reader
	reader = bytes.NewReader(data)
	reader2 := csv.NewReader(reader)
	recs, err := reader2.ReadAll()

	records := make([]*record.Input, len(recs)-1)
	for idx, rec := range recs[1:] {
		valid, err := strconv.ParseBool(rec[4])
		if err != nil {
			return err
		}

		if rec[2] != "" {
			section, err := strconv.Atoi(rec[2])
			if err != nil {
				fmt.Println(rec)
				return err
			}

			if rec[3] != "" {
				row, err := strconv.Atoi(rec[3])
				if err != nil {
					return err
				}
				records[idx] = record.NewInput(record.NewPair(&section, rec[0]), record.NewPair(&row, rec[1]), valid)
			} else {
				records[idx] = record.NewInput(record.NewPair(&section, rec[0]), record.NewPair(nil, rec[1]), valid)
			}
		} else {
			records[idx] = record.NewInput(record.NewPair(nil, rec[0]), record.NewPair(nil, rec[1]), valid)
		}

	}

	i.Records = records

	return nil
}
