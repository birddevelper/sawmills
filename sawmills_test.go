package main

import "testing"

type ProfitTest struct {
	trunk    []int
	expected int
}

type CutTest struct {
	trunk        int
	cutterLength int
	expected     []int
}

func TestCut(t *testing.T) {
	var CutTests = []CutTest{
		CutTest{1, 3, []int{1, 0, 2}},
		CutTest{2, 3, []int{2, 0, 1}},
		CutTest{3, 3, []int{3, 0, 3}},
		CutTest{4, 3, []int{3, 1, 3}},
		CutTest{3, 1, []int{1, 2, 3}},
		CutTest{3, 2, []int{2, 1, 3}},
		CutTest{2, 2, []int{2, 0, 3}}}

	for _, test := range CutTests {
		if cutLength, remainedTrunk, RemainedCutter := Cut(test.trunk, test.cutterLength); cutLength != test.expected[0] ||
			remainedTrunk != test.expected[1] || RemainedCutter != test.expected[2] {
			t.Errorf("Output %v not equal to expected %v", []int{cutLength, remainedTrunk, RemainedCutter}, test.expected)
		}
	}

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

func TestPermutation(t *testing.T) {
	trunks := []int{1, 2, 2}
	expected := 3
	perms := permutation(trunks)
	if len(perms) != expected {
		t.Errorf("Output %d not equal to expected %d", len(perms), expected)
	}

}
