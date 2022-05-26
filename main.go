package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// returns cut lenght, remained trunck length, remained cutter length
func cut(trunk int, cutterLength int) (int, int, int) {

	if cutterLength >= trunk {
		return trunk, 0, cutterLength - trunk
	} else {
		return cutterLength, trunk - cutterLength, 3
	}
}

// calculate the profit of a permutation of trunks
func ProfitCalculator(trunks []int) int {
	totalProfit := 0
	currentcutterLength := 3
	for _, trunk := range trunks {

		trunkRemainedLength := trunk
		for trunkRemainedLength > 0 {
			var cutLength int
			cutLength, trunkRemainedLength, currentcutterLength = cut(trunkRemainedLength, currentcutterLength)
			if cutLength == 1 {
				totalProfit += -1
			} else if cutLength == 2 {
				totalProfit += 3
			} else if cutLength == 3 {
				totalProfit += 1
			}

		}

	}

	return totalProfit
}

// permutation of trunks thrown in river
func permutation(trunks []int) (permuts [][]int) {
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			permuts = append(permuts, append([]int{}, a...))
		} else {
			for i := k; i < len(trunks); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(trunks, 0)

	return permuts
}

func findMaxProfit(trunks []int) int {
	trunksPerms := permutation(trunks)
	maxProfit := 0
	for i := 0; i < len(trunksPerms); i++ {
		profit := ProfitCalculator(trunksPerms[i])
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}

func main() {
	type Cases struct {
		CaseNumber int
		sawmills   [][]int
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Give Z : ")
	Z, _ := reader.ReadString('\n')
	sawmillsCount, _ := strconv.Atoi(Z)
	for i := 0; i < sawmillsCount; i++ {
		trunks, _ := reader.ReadString('\n')
		fmt.Printf(trunks)
	}

}
