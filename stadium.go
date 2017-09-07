package seatgeek

import (
	"github.com/andrewcopp/seatgeek/record"
)

type Stadium struct {
	Sections map[string]*Section
}

func NewStadium(m *Manifest) *Stadium {
	sections := map[string]*Section{}
	for _, rec := range m.Records {
		if _, ok := sections[rec.Section.Value]; !ok {
			sections[rec.Section.Value] = NewSection(rec.Section)
		}

		if rec.Row != nil {
			sections[rec.Section.Value].Rows[rec.Row.Value] = NewRow(rec.Row)
		}
	}

	return &Stadium{
		Sections: sections,
	}
}

type Section struct {
	Identifier int
	Value      string
	Rows       map[string]*Row
}

func NewSection(pair *record.Pair) *Section {
	return &Section{
		Identifier: *pair.Identifier,
		Value:      pair.Value,
		Rows:       map[string]*Row{},
	}
}

type Row struct {
	Identifier int
	Value      string
}

func NewRow(pair *record.Pair) *Row {
	return &Row{
		Identifier: *pair.Identifier,
		Value:      pair.Value,
	}
}
