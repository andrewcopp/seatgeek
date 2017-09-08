package seatgeek

import (
	"strings"
)

type Ticket struct {
	Section string
	Row     *string
}

func NewTicket(sec string, row *string) *Ticket {
	return &Ticket{
		Section: sec,
		Row:     row,
	}
}

func (t *Ticket) Clean() string {
	segs := strings.Split(t.Section, " ")

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
		prefix := t.Prefix(num)
		suffix := t.Suffix(num)
		stripped := strings.TrimSuffix(strings.TrimPrefix(num, prefix), suffix)
		expanded := t.Expand(suffix)
		if expanded != "" {
			nums[idx] = expanded + " " + stripped
		} else {
			nums[idx] = stripped
		}
	}

	for idx, word := range words {
		expanded := t.Expand(word)
		words[idx] = expanded
	}

	for idx, word := range words {
		capitalized := t.Capitalize(word)
		words[idx] = capitalized
	}

	words = t.Exchange(words)
	words = t.Eliminate(words)

	row := t.Capitalize(*t.Row)
	t.Row = &row

	result := strings.Join(append(words, nums...), " ")

	result = t.Complete(result)

	return result

}

func (t *Ticket) Suffix(num string) string {
	chars := strings.Split(num, "")
	for idx, char := range chars {
		if !strings.ContainsAny(char, "0123456789") {
			return num[idx:]
		}
	}

	return ""
}

func (t *Ticket) Prefix(num string) string {
	chars := strings.Split(num, "")
	for idx, char := range chars {
		if strings.ContainsAny(char, "0123456789") {
			return num[:idx]
		}
	}
	return ""
}

func (t *Ticket) Capitalize(phrase string) string {
	words := strings.Split(phrase, " ")
	for idx, word := range words {
		lower := strings.ToLower(word)
		first := strings.SplitN(lower, "", 2)[0]
		words[idx] = strings.ToUpper(first) + lower[1:]
	}

	return strings.Join(words, " ")
}

func (t *Ticket) Eliminate(words []string) []string {
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

func (t *Ticket) Expand(suffix string) string {
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

func (t *Ticket) Complete(phrase string) string {
	if strings.Contains(phrase, "Loge") && !strings.Contains(phrase, "Loge Box") {
		phrase = strings.Replace(phrase, "Loge", "Loge Box", 1)
	}

	if strings.Contains(phrase, "Field") && !strings.Contains(phrase, "Field Box") {
		phrase = strings.Replace(phrase, "Field", "Field Box", 1)
	}

	return phrase
}

func (t *Ticket) Exchange(words []string) []string {
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

type Test struct {
	Ticket *Ticket
	Output *Output
}

func NewTest(tkt *Ticket, out *Output) *Test {
	return &Test{
		Ticket: tkt,
		Output: out,
	}
}
