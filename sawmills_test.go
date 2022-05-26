package main

import "testing"

type ProfitTest struct {
	trunk    []int
	expected int
}

var ProfitTests = []ProfitTest{
	ProfitTest{[]int{2, 3, 1}, 4},
	ProfitTest{[]int{1, 3, 2}, 4},
	ProfitTest{[]int{1, 4}, 8},
}

func TestCalcSeries(t *testing.T) {

	for _, test := range ProfitTests {
		if output := ProfitCalculator(test.trunk); output != test.expected {
			t.Errorf("Output %d not equal to expected %d", output, test.expected)
		}
	}
}
