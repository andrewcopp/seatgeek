package main

import (
	"fmt"
	"log"

	"github.com/andrewcopp/seatgeek"
)

func main() {

	config := seatgeek.NewConfig()

	m := seatgeek.NewManifest()
	m.Load(config.Manifest)

	in := seatgeek.NewInput()
	err := in.Load(*config.Input)
	if err != nil {
		log.Fatalln(err.Error())
	}

	dist := seatgeek.NewDistributor(in)

	correct := 0

	stdm := seatgeek.NewStadium(m)
	var checker seatgeek.Checker
	checker = seatgeek.NewAttendant(stdm)
	norm := seatgeek.NewDefault(&checker)
	for _, test := range dist.Tests {
		section, row, valid := norm.Normalize(test.Ticket.Section, test.Ticket.Row)
		out := test.Output

		if section == nil {
			temp := -1
			section = &temp
		}

		if row == nil {
			temp := -1
			row = &temp
		}

		if out.Section == nil {
			temp := -1
			out.Section = &temp
		}

		if out.Row == nil {
			temp := -1
			out.Row = &temp
		}

		if *section == *out.Section && *row == *out.Row && valid == out.Valid {
			correct++
		} else {
			fmt.Println("Failure: ", *section, *out.Section, *row, *out.Row, valid, out.Valid)
		}
	}

	fmt.Printf("%.1f%% correct\n", float64(correct)/float64(len(dist.Tests))*100.0)
}
