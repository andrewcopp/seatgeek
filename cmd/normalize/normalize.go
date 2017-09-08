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
	man := seatgeek.NewManifest()
	man.Load(config.Manifest)

	if config.Input != nil {
		smps := Read(*config.Input)
		matches := Normalize(man, smps, false)
		Output(matches)
	} else if config.Section != nil && config.Row != nil {
		section, row, valid := man.Normalize(*config.Section, config.Row)

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

// Read takes a filepath to an input CSV and translates the data into Sample
// objects.
func Read(path string) []*seatgeek.Sample {
	in := seatgeek.NewInput()
	err := in.Load(path)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return in.Samples
}

// Normalize takes an object that can normalize an Input and a slice of Samples
// and maps it to the correct ticket in an Output struct. The Output is combined
// with the Sample to create a slice of Match objects to return.
func Normalize(norm seatgeek.Normalizer, smps []*seatgeek.Sample, verb bool) []*seatgeek.Match {
	matches := make([]*seatgeek.Match, len(smps))
	for idx, sample := range smps {
		section, row, valid := norm.Normalize(sample.Input.Section, &sample.Input.Row)
		out := seatgeek.NewOutput(section, row, valid)
		matches[idx] = seatgeek.NewMatch(sample, out)
	}
	return matches
}

// Output prints the calculated Match objects in JSON format to the console for
// grading.
func Output(mtcs []*seatgeek.Match) {
	for _, mtc := range mtcs {
		result, err := json.Marshal(mtc)
		if err != nil {
			log.Fatalln(err.Error())
		}

		fmt.Println(string(result))
	}
}
