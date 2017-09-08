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

	words := Words(section)
	num := Number(section)

	words = Expand(words)
	words = Capitalize(words)

	words = Exchange(words)
	words = Eliminate(words)

	result := strings.Join(append(words, num), " ")

	result = Complete(result)

	perms := Shorten(result)
	for _, perm := range perms {
		if sec, ok := m.Sections[perm]; ok {
			row := capitalize(*row)
			result := Lock(sec, &row)
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

func Words(phrase string) []string {
	re := regexp.MustCompile("[^a-zA-Z ]")
	phrase = re.ReplaceAllString(phrase, " ")

	words := strings.Split(phrase, " ")
	for i := len(words) - 1; i >= 0; i-- {
		if words[i] == "" {
			words = append(words[:i], words[i+1:]...)
		}
	}

	return words
}

func Number(phrase string) string {
	re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(phrase, "")
}

func Expand(words []string) []string {
	results := []string{}
	for _, word := range words {
		switch word {
		case "BL":
			results = append(results, "Baseline")
			results = append(results, "Club")
		case "FD":
			results = append(results, "Field")
			results = append(results, "Box")
		case "TD":
			results = append(results, "Top")
			results = append(results, "Deck")
		case "RS":
			results = append(results, "Reserve")
		case "LG":
			results = append(results, "Loge")
			results = append(results, "Box")
		case "DG":
			results = append(results, "Dugout")
			results = append(results, "Club")
		case "PR":
			results = append(results, "Right")
			results = append(results, "Field")
			results = append(results, "Pavilion")
		default:
			results = append(results, word)
		}
	}

	return results
}

func Capitalize(words []string) []string {
	for idx, word := range words {
		words[idx] = capitalize(word)
	}

	return words
}

func capitalize(word string) string {
	lower := strings.ToLower(word)
	first := strings.SplitN(lower, "", 2)[0]
	return strings.ToUpper(first) + lower[1:]
}

func Exchange(words []string) []string {
	for idx, word := range words {
		switch word {
		case "Infield":
			words[idx] = "Field"
		}
	}
	return words
}

func Eliminate(words []string) []string {
	results := []string{}
	for _, word := range words {
		switch word {
		case "Reserve":
			results = append(results, word)
		case "Field":
			results = append(results, word)
		case "Box":
			results = append(results, word)
		case "Field Box":
			results = append(results, word)
		case "Top":
			results = append(results, word)
		case "Deck":
			results = append(results, word)
		case "Loge":
			results = append(results, word)
		case "Right":
			results = append(results, word)
		case "Left":
			results = append(results, word)
		case "Pavilion":
			results = append(results, word)
		case "Dugout":
			results = append(results, word)
		case "Club":
			results = append(results, word)
		case "Baseline":
			results = append(results, word)
		}
	}

	return results
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

func Complete(phrase string) string {
	if strings.Contains(phrase, "Loge") && !strings.Contains(phrase, "Loge Box") {
		phrase = strings.Replace(phrase, "Loge", "Loge Box", 1)
	}

	if strings.Contains(phrase, "Field") && !strings.Contains(phrase, "Field Box") {
		phrase = strings.Replace(phrase, "Field", "Field Box", 1)
	}

	return phrase
}
