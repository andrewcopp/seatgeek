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

	segs := strings.Split(section, " ")

	words := []string{}
	nums := []string{}
	for _, seg := range segs {
		if strings.ContainsAny(seg, "0123456789") {
			nums = append(nums, seg)
		} else {
			words = append(words, seg)
		}
	}

	for idx, num := range nums {
		prefix := Prefix(num)
		suffix := Suffix(num)
		stripped := strings.TrimSuffix(strings.TrimPrefix(num, prefix), suffix)
		expanded := Expand(suffix)
		if expanded != "" {
			nums[idx] = expanded + " " + stripped
		} else {
			nums[idx] = stripped
		}
	}

	for idx, word := range words {
		expanded := Expand(word)
		words[idx] = expanded
	}

	for idx, word := range words {
		capitalized := Capitalize(word)
		words[idx] = capitalized
	}

	words = Exchange(words)
	words = Eliminate(words)

	// row := Capitalize(row)
	// t.Row = &row

	result := strings.Join(append(words, nums...), " ")

	result = Complete(result)

	perms := Shorten(result)
	for _, perm := range perms {
		if sec, ok := m.Sections[perm]; ok {
			result := Lock(sec, row)
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

func Suffix(num string) string {
	chars := strings.Split(num, "")
	for idx, char := range chars {
		if !strings.ContainsAny(char, "0123456789") {
			return num[idx:]
		}
	}

	return ""
}

func Prefix(num string) string {
	chars := strings.Split(num, "")
	for idx, char := range chars {
		if strings.ContainsAny(char, "0123456789") {
			return num[:idx]
		}
	}
	return ""
}

func Capitalize(phrase string) string {
	words := strings.Split(phrase, " ")
	for idx, word := range words {
		lower := strings.ToLower(word)
		first := strings.SplitN(lower, "", 2)[0]
		words[idx] = strings.ToUpper(first) + lower[1:]
	}

	return strings.Join(words, " ")
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
		}
	}

	return results
}

func Expand(suffix string) string {
	switch suffix {
	case "BL":
		return "Baseline Club"
	case "FD":
		return "Field Box"
	case "TD":
		return "Top Deck"
	case "RS":
		return "Reserve"
	case "LG":
		return "Loge Box"
	case "DG":
		return "Dugout Club"
	case "PR":
		return "Right Field Pavilion"
	}

	return suffix
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

func Exchange(words []string) []string {
	results := make([]string, len(words))
	for idx, word := range words {
		switch word {
		case "Infield":
			results[idx] = "Field"
		default:
			results[idx] = word
		}
	}
	return results
}
