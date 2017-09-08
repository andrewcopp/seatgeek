package seatgeek

import (
	"fmt"
	"strconv"

	"github.com/andrewcopp/seatgeek/record"
)

type Manifest struct {
	Sections map[string]*Section
}

func NewManifest() *Manifest {
	return &Manifest{
		Sections: map[string]*Section{},
	}
}

type Section struct {
	Identifier int
	Value      string
	Rows       map[string]*Row
}

func NewSection(id int, val string) *Section {
	return &Section{
		Identifier: id,
		Value:      val,
		Rows:       map[string]*Row{},
	}
}

type Row struct {
	Identifier int
	Value      string
}

func NewRow(id int, val string) *Row {
	return &Row{
		Identifier: id,
		Value:      val,
	}
}

func (m *Manifest) Load(path string) error {
	records := record.Read(path)
	fmt.Println(len(records))
	for _, record := range records[1:] {
		if _, ok := m.Sections[record[1]]; !ok {
			id, err := strconv.Atoi(record[0])
			if err != nil {
				return err
			}
			m.Sections[record[1]] = NewSection(id, record[1])
		}

		if record[3] != "" {
			id, err := strconv.Atoi(record[2])
			if err != nil {
				return err
			}
			m.Sections[record[1]].Rows[record[3]] = NewRow(id, record[3])
		}
	}

	return nil
}
