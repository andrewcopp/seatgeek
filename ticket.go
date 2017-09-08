package seatgeek

import (
	"regexp"
	"strings"
)

// Ticket represents the object is attempted to be passed off for a real ticket.
// This item should be called "Input" but it collides with the naming for the
// object that encapulates the content of input CSV.
type Ticket struct {
	Section string `json:"section"`
	Row     string `json:"row"`
}

// NewTicket returns a Ticket object that represents a section and row.
func NewTicket(sec string, row string) *Ticket {
	return &Ticket{
		Section: sec,
		Row:     row,
	}
}

// Normalize processes the given name for a section and row and returns the best
// guess at what the ticket should actually be to map to a real ticket.
func (t *Ticket) Normalize() (string, string) {

	words := t.Words()
	words = t.Expand(words)
	words = t.Capitalize(words)
	words = t.Exchange(words)
	words = t.Eliminate(words)

	section := strings.Join(append(words, t.Number()), " ")

	return t.Complete(section), t.capitalize(t.Row)
}

// Words returns all of the components of the section that resemble words.
// Abbrieviations that are connected to a number are split out as their own
// word.
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

// Number extracts the number from the section name.
func (t *Ticket) Number() string {
	re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(t.Section, "")
}

// Expand takes any abbrievations in a slice of strings and maps them to a list
// of known results to be replaced.
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

// Capitalize takes a slice of strings and returns a slice of the same strings
// properly capitalized. The first character in each word is uppercase and all
// other characters are lowercase.
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

// Exchange maps words to a list of words that are known to be common
// substitutes for other words.
func (t *Ticket) Exchange(words []string) []string {
	for idx, word := range words {
		switch word {
		case "Infield":
			words[idx] = "Field"
		}
	}
	return words
}

// Eliminate deletes all words that aren't known to be on the manifest in some
// form or another.
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

// Complete takes known words and adds more words to reach a complete phrase
// if necessary.
func (t *Ticket) Complete(phrase string) string {
	if strings.Contains(phrase, "Loge") && !strings.Contains(phrase, "Loge Box") {
		phrase = strings.Replace(phrase, "Loge", "Loge Box", 1)
	}

	if strings.Contains(phrase, "Field") && !strings.Contains(phrase, "Field Box") {
		phrase = strings.Replace(phrase, "Field", "Field Box", 1)
	}

	return phrase
}
