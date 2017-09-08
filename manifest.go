package seatgeek

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
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
	records := Read(path)
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

func (m *Manifest) Normalize(section string, row *string) (*int, *int, bool) {

	r := ""
	if row != nil {
		r = *row
	}

	tkt := NewTicket(section, r)

	section, r = tkt.Normalize()

	perms := Shorten(section)
	for _, perm := range perms {
		if sec, ok := m.Sections[perm]; ok {
			result := Lock(sec, &r)
			return result.Section, result.Row, result.Valid
		}
	}

	re := regexp.MustCompile("[^0-9]")

	// Only Possibility
	sections := map[string][]*Section{}
	for _, section := range m.Sections {
		number := re.ReplaceAllString(section.Value, "")

		if number != "" {
			if _, ok := sections[number]; !ok {
				sections[number] = []*Section{}
			}

			sections[number] = append(sections[number], section)
		}
	}

	number := re.ReplaceAllString(section, "")
	if len(sections[number]) == 1 {
		sec := sections[number][0]
		result := Lock(sec, row)
		return result.Section, result.Row, result.Valid
	}

	// No Match
	return nil, nil, false

}

func Lock(sec *Section, row *string) *Output {
	if row != nil {
		if row, ok := sec.Rows[strings.TrimLeft(*row, "0")]; ok {
			return NewOutput(&sec.Identifier, &row.Identifier, true)
		}
	}

	// Extraneous Row
	if len(sec.Rows) != 0 {
		return NewOutput(nil, nil, false)
	}

	return NewOutput(&sec.Identifier, nil, true)
}

func Shorten(phrase string) []string {
	words := strings.Split(phrase, " ")
	if len(words) == 1 {
		return words
	}

	results := []string{}
	for idx := range words {
		new := make([]string, len(words[:idx]))
		for j, word := range words[:idx] {
			new[j] = word
		}
		new = append(new, words[idx+1:]...)
		result := Shorten(strings.Join(new, " "))
		results = append(results, result...)
	}

	results = append(results, strings.Join(words, " "))
	return results
}
