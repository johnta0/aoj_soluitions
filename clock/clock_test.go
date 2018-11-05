package clock

import (
	"testing"
	"strconv"
)

func TestCalcFromShortHand(t *testing.T) {
	
	tests := []struct {
		input	int
		output  string
	}{
		{input: 0, output: "0 0"},
		{input: 45, output: "1 30"},
		{input: 100, output: "3 20"},
	}

	for _, test := range tests {
		h, m := CalcFromShortHand(test.input)
		ans := strconv.Itoa(h) + " " + strconv.Itoa(m)

		if ans != test.output {
			t.Errorf("Calculated output is %s want %s", ans, test.output)
		}
	}
}
