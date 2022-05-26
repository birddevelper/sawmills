package main

import "testing"

type ProfitTest struct {
	trunk    []int
	expected int
}

func TestCalcSeries(t *testing.T) {
	var ProfitTests = []ProfitTest{
		ProfitTest{[]int{2, 3, 1}, 4},
		ProfitTest{[]int{1, 3, 2}, 4},
		ProfitTest{[]int{1, 4}, 5},
	}
	for _, test := range ProfitTests {
		if output := ProfitCalculator(test.trunk); output != test.expected {
			t.Errorf("Output %d not equal to expected %d", output, test.expected)
		}
	}
}

func TestFindMaxProfit(t *testing.T) {
	var ProfitTests = []ProfitTest{
		ProfitTest{[]int{1, 2, 3}, 4},
		ProfitTest{[]int{4, 1}, 5},
	}
	for _, test := range ProfitTests {
		if output := findMaxProfit(test.trunk); output != test.expected {
			t.Errorf("Output %d not equal to expected %d", output, test.expected)
		}
	}
}

func TestFindPermutionWithMaxProfit(t *testing.T) {
	var ProfitTests = []ProfitTest{
		ProfitTest{[]int{1, 2, 3}, 2},
		ProfitTest{[]int{4, 1}, 1},
	}

	for _, test := range ProfitTests {
		if output := findPermutionWithMaxProfit(test.trunk); len(output) != test.expected {
			t.Errorf("Output %d not equal to expected %d", output, test.expected)
		}
	}
}
