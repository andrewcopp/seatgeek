package record

import (
	"bytes"
	"encoding/csv"
	"io/ioutil"
	"log"
)

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
