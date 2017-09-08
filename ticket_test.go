package seatgeek

import "testing"

func TestNormalize(t *testing.T) {
	cases := []struct {
		in  *Ticket
		sec string
		row string
		err string
	}{
		{NewTicket("301", "A"), "301", "A", "Basic tickets should not be modified."},
		{NewTicket("BL 1", "A"), "Baseline Club 1", "A", "BL -> Baseline Clb"},
		{NewTicket("FD 1", "A"), "Field Box 1", "A", "FD -> Field Box"},
		{NewTicket("TD 1", "A"), "Top Deck 1", "A", "TD -> Top Deck"},
		{NewTicket("RS 1", "A"), "Reserve 1", "A", "RS -> Reserve"},
		{NewTicket("LG 1", "A"), "Loge Box 1", "A", "LG -> Loge Box"},
		{NewTicket("DG 1", "A"), "Dugout Club 1", "A", "DG -> Dugout Club"},
		{NewTicket("PR 1", "A"), "Right Field Box Pavilion 1", "A", "PR -> Right Field Pavilion -> Right Field Box Pavilion"},
		{NewTicket("RESERVE 1", "A"), "Reserve 1", "A", "Uppercased words should simply be capitalized."},
		{NewTicket("301", "HH"), "301", "Hh", "Rows that are more than ne letter should be capitalized."},
		{NewTicket("Infield 1", "A"), "Field Box 1", "A", "The word 'infield' should be mapped to 'field'."},
		{NewTicket("Left 1", "A"), "Left 1", "A", "Left should not be removed from possible guesses."},
		{NewTicket("Right 1", "A"), "Right 1", "A", "Right should not be removed from possible guesses."},
		{NewTicket("Pavilion 1", "A"), "Pavilion 1", "A", "Pavilion should not be removed from possible guesses."},
		{NewTicket("Loge 1", "A"), "Loge Box 1", "A", "People who enter 'Loge' really meant 'Loge Box'."},
		{NewTicket("Field 1", "A"), "Field Box 1", "A", "People who enter 'Field' really mean 'Field Box.'"},
	}
	for _, c := range cases {
		sec, row := c.in.Normalize()
		if sec != c.sec || row != c.row {
			t.Error(c.err)
		}
	}
}
