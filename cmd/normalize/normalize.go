package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/andrewcopp/seatgeek"
)

type Match struct {
	Expected match `json:"expected"`
	Output   match `json:"output"`
}

type match struct {
	Section string `json:"section_id"`
	Row     string `json:"row_id"`
	Valid   bool   `json:"valid"`
}

func NewMatch(sec, row *int, valid bool) match {
	s := ""
	if sec != nil {
		s = strconv.Itoa(*sec)
	}

	r := ""
	if row != nil {
		r = strconv.Itoa(*row)
	}

	return match{
		Section: s,
		Row:     r,
		Valid:   valid,
	}
}

func main() {

	config := seatgeek.NewConfig()

	m := seatgeek.NewManifest()
	m.Load(config.Manifest)

	stdm := seatgeek.NewStadium(m)
	var checker seatgeek.Checker
	checker = seatgeek.NewAttendant(stdm)
	norm := seatgeek.NewDefault(&checker)

	if config.Input != nil {
		in := seatgeek.NewInput()
		err := in.Load(*config.Input)
		if err != nil {
			log.Fatalln(err.Error())
		}

		dist := seatgeek.NewDistributor(in)

		matches := make([]map[string]map[string]interface{}, len(dist.Tests))
		for idx, test := range dist.Tests {
			section, row, valid := norm.Normalize(test.Ticket.Section, test.Ticket.Row)
			out := test.Output

			matches[idx] = map[string]map[string]interface{}{
				"input": map[string]interface{}{
					"section": test.Ticket.Section,
					"row":     test.Ticket.Row,
				},
				"expected": map[string]interface{}{
					"section_id": out.Section,
					"row_id":     out.Row,
					"valid":      out.Valid,
				},
				"output": map[string]interface{}{
					"section_id": section,
					"row_id":     row,
					"valid":      valid,
				},
			}
		}

		for _, match := range matches {
			result, err := json.Marshal(match)
			if err != nil {
				log.Fatalln(err.Error())
			}

			fmt.Println(string(result))
		}

	} else if config.Section != nil && config.Row != nil {
		section, row, valid := norm.Normalize(*config.Section, config.Row)

		out1 := ""
		if section != nil {
			out1 = strconv.Itoa(*section)
		}

		out2 := ""
		if row != nil {
			out2 = strconv.Itoa(*row)
		}

		fmt.Printf("Input:\n\t[section] %s\t[row] %s\nOutput:\n\t[section_id] %s\t[row_id] %s\nValid?:\n\t%t\n", *config.Section, *config.Row, out1, out2, valid)
	}
}
