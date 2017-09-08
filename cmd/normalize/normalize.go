package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/andrewcopp/seatgeek"
)

func main() {

	config := seatgeek.NewConfig()

	m := seatgeek.NewManifest()
	m.Load(config.Manifest)

	if config.Input != nil {
		in := seatgeek.NewInput()
		err := in.Load(*config.Input)
		if err != nil {
			log.Fatalln(err.Error())
		}

		for _, sample := range in.Samples {
			section, row, valid := m.Normalize(sample.Input.Section, &sample.Input.Row)
			out := seatgeek.NewOutput(section, row, valid)
			match := seatgeek.NewMatch(sample, out)

			result, err := json.Marshal(match)
			if err != nil {
				log.Fatalln(err.Error())
			}

			fmt.Println(string(result))
		}

	} else if config.Section != nil && config.Row != nil {
		section, row, valid := m.Normalize(*config.Section, config.Row)

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

// Read
// Normalize
// Output
