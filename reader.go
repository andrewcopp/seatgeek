package seatgeek

import (
	"bytes"
	"encoding/csv"
	"io/ioutil"
	"log"
)

// Read takes a filepath to a CSV and returns a slice of records.
func Read(path string) [][]string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err.Error())
	}

	reader := csv.NewReader(bytes.NewReader(data))
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalln(err.Error())
	}

	return records
}
