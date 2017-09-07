package seatgeek

import (
	"bytes"
	"encoding/csv"
	"io"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/andrewcopp/seatgeek/record"
)

type Pair struct {
	Identifier int
	Value      string
}

func NewPair(id int, val string) *Pair {
	return &Pair{
		Identifier: id,
		Value:      val,
	}
}

type Line struct {
	Section *Pair
	Row     *Pair
}

func NewLine(sec, row *Pair) *Line {
	return &Line{
		Section: sec,
		Row:     row,
	}
}

type Manifest struct {
	Records []*record.Manifest
}

func NewManifest() *Manifest {
	return &Manifest{
		Records: []*record.Manifest{},
	}
}

func (m *Manifest) Load(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err.Error())
	}

	var reader io.Reader
	reader = bytes.NewReader(data)
	reader2 := csv.NewReader(reader)
	recs, err := reader2.ReadAll()
	if err != nil {
		log.Fatalln(err.Error())
	}

	records := make([]*record.Manifest, len(recs)-1)
	for idx, rec := range recs[1:] {
		section, err := strconv.Atoi(rec[0])
		if err != nil {
			return err
		}

		if rec[2] != "" {
			row, err := strconv.Atoi(rec[2])
			if err != nil {
				return err
			}
			records[idx] = record.NewManifest(record.NewPair(&section, rec[1]), record.NewPair(&row, rec[3]))
		} else {
			records[idx] = record.NewManifest(record.NewPair(&section, rec[1]), nil)
		}
	}

	m.Records = records

	return nil
}
