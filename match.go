package seatgeek

import (
	"regexp"
	"strings"
)

type Match struct {
	*Sample
	Output *Output `json:"output"`
}

func NewMatch(smp *Sample, out *Output) *Match {
	return &Match{
		Sample: smp,
		Output: out,
	}
}

type Sample struct {
	Input    *Ticket   `json:"input"`
	Expected *Expected `json:"expected"`
}

func NewSample(in *Ticket, exp *Expected) *Sample {
	return &Sample{
		Input:    in,
		Expected: exp,
	}
}

type Result struct {
	Section *int `json:"section_id"`
	Row     *int `json:"row_id"`
	Valid   bool `json:"valid"`
}

func NewResult(sec *int, row *int, val bool) *Result {
	return &Result{
		Section: sec,
		Row:     row,
		Valid:   val,
	}
}

type Output struct {
	*Result
}

func NewOutput(sec *int, row *int, val bool) *Output {
	return &Output{
		Result: NewResult(sec, row, val),
	}
}

type Ticket struct {
	Section string `json:"section"`
	Row     string `json:"row"`
}

func NewTicket(sec string, row string) *Ticket {
	return &Ticket{
		Section: sec,
		Row:     row,
	}
}

func (t *Ticket) Normalize() (string, string) {

	words := t.Words()
	words = t.Expand(words)
	words = t.Capitalize(words)
	words = t.Exchange(words)
	words = t.Eliminate(words)

	section := strings.Join(append(words, t.Number()), " ")

	return t.Complete(section), t.capitalize(t.Row)
}

func (t *Ticket) Words() []string {
	re := regexp.MustCompile("[^a-zA-Z ]")
	phrase := re.ReplaceAllString(t.Section, " ")

	words := strings.Split(phrase, " ")
	for i := len(words) - 1; i >= 0; i-- {
		if words[i] == "" {
			words = append(words[:i], words[i+1:]...)
		}
	}

	return words
}

func (t *Ticket) Number() string {
	re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(t.Section, "")
}

func (t *Ticket) Expand(words []string) []string {
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

func (t *Ticket) Capitalize(words []string) []string {
	for idx, word := range words {
		words[idx] = t.capitalize(word)
	}

	return words
}

func (t *Ticket) capitalize(word string) string {
	lower := strings.ToLower(word)
	first := strings.SplitN(lower, "", 2)[0]
	return strings.ToUpper(first) + lower[1:]
}

func (t *Ticket) Exchange(words []string) []string {
	for idx, word := range words {
		switch word {
		case "Infield":
			words[idx] = "Field"
		}
	}
	return words
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

func (t *Ticket) Complete(phrase string) string {
	if strings.Contains(phrase, "Loge") && !strings.Contains(phrase, "Loge Box") {
		phrase = strings.Replace(phrase, "Loge", "Loge Box", 1)
	}

	if strings.Contains(phrase, "Field") && !strings.Contains(phrase, "Field Box") {
		phrase = strings.Replace(phrase, "Field", "Field Box", 1)
	}

	return phrase
}

type Expected struct {
	*Result
}

func NewExpected(sec *int, row *int, val bool) *Expected {
	return &Expected{
		Result: NewResult(sec, row, val),
	}
}
