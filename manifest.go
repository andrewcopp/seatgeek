package seatgeek

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Manifest represents the contents of a manifest CSV files. The contents are
// parsed into a two-deep map of Sections and Rows.
type Manifest struct {
	Sections map[string]*Section
}

// NewManifest returns an Manifest struct with an empty map that is ready to add
// values.
func NewManifest() *Manifest {
	return &Manifest{
		Sections: map[string]*Section{},
	}
}

// Section represents the id and name of a section and holds references to all
// valid rows.
type Section struct {
	Identifier int
	Value      string
	Rows       map[string]*Row
}

// NewSection returns a Section struct with correct values and an empty map
// ready to accept Row structs.
func NewSection(id int, val string) *Section {
	return &Section{
		Identifier: id,
		Value:      val,
		Rows:       map[string]*Row{},
	}
}

// Row represents the identifer and value of a row.
type Row struct {
	Identifier int
	Value      string
}

// NewRow returns a Row struct populated with the specified identifier and
// value.
func NewRow(id int, val string) *Row {
	return &Row{
		Identifier: id,
		Value:      val,
	}
}

// Load accepts a filepath to a manifest CSV file into a Manifest struct
// for O(1) lookups.
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

// Normalize satisfies the specifed interface of the challenge by taking a
// section and row and making the prediction of what those values should map to
// against the manifest.
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

// Lock is called when you are able to map the section from an Input to a
// section on the manifest. Calling this function means you are committed to
// making a prediction.
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

// Shorten returns all possible arrangements a word where zero to n-1 of the
// words are included.
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
