package seatgeek

import (
	"fmt"
	"regexp"
	"strings"
)

type Checker interface {
	Check(tkt *Ticket) *Output
}

type Attendant struct {
	Manifest *Manifest
	re       *regexp.Regexp
}

func NewAttendant(manifest *Manifest) *Attendant {
	return &Attendant{
		Manifest: manifest,
		re:       regexp.MustCompile("[^0-9]"),
	}
}

func (a *Attendant) Check(tkt *Ticket) *Output {

	fmt.Println(a.Manifest)

	perms := a.Shorten(tkt.Clean())
	for _, perm := range perms {
		if sec, ok := a.Manifest.Sections[perm]; ok {
			return a.Lock(tkt, sec)
		}
	}

	// Only Possibility
	sections := map[string][]*Section{}
	for _, section := range a.Manifest.Sections {
		number := a.re.ReplaceAllString(section.Value, "")

		if number != "" {
			if _, ok := sections[number]; !ok {
				sections[number] = []*Section{}
			}

			sections[number] = append(sections[number], section)
		}
	}

	number := a.re.ReplaceAllString(tkt.Section, "")
	if len(sections[number]) == 1 {
		sec := sections[number][0]
		return a.Lock(tkt, sec)
	}

	// No Match
	return NewOutput(nil, nil, false)
}

func (a *Attendant) Lock(tkt *Ticket, sec *Section) *Output {
	if tkt.Row != nil {
		if row, ok := sec.Rows[strings.TrimLeft(*tkt.Row, "0")]; ok {
			return NewOutput(&sec.Identifier, &row.Identifier, true)
		}
	}

	// Extraneous Row
	if len(sec.Rows) != 0 {
		return NewOutput(nil, nil, false)
	}

	return NewOutput(&sec.Identifier, nil, true)
}

func (a *Attendant) Shorten(phrase string) []string {
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
		result := a.Shorten(strings.Join(new, " "))
		results = append(results, result...)
	}

	results = append(results, strings.Join(words, " "))
	return results
}
