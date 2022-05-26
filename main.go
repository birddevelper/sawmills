package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func findPermutionWithMaxProfit(trunks []int) [][]int {
	var permutionWithMaxProfit [][]int
	maxProfit := findMaxProfit(trunks)
	trunksPerms := permutation(trunks)
	for i := 0; i < len(trunksPerms); i++ {
		profit := ProfitCalculator(trunksPerms[i])
		if profit == maxProfit {
			permutionWithMaxProfit = append(permutionWithMaxProfit, trunksPerms[i])
		}
	}

	return permutionWithMaxProfit
}

func main() {
	type Case struct {
		CaseNumber int
		sawmills   [][]int
	}

	var Cases []Case

	reader_1 := bufio.NewReader(os.Stdin)
	fmt.Print("Give Z : ")
	Z, _ := reader_1.ReadString('\n')
	caseNo := 0

	sawmillsCount, _ := strconv.Atoi(strings.Replace(Z, "\r\n", "", -1))

	for sawmillsCount != 0 {

		caseNo += 1
		var sawmills [][]int
		for i := 0; i < sawmillsCount; i++ {
			reader_2 := bufio.NewReader(os.Stdin)
			trunks, _ := reader_2.ReadString('\n')
			trunksData := strings.Fields(trunks)
			var trunkSeq []int
			for i := 1; i < len(trunksData); i++ {
				trunkLength, _ := strconv.Atoi(trunksData[i])
				trunkSeq = append(trunkSeq, trunkLength)
			}
			sawmills = append(sawmills, trunkSeq)
		}
		myCase := Case{caseNo, sawmills}
		Cases = append(Cases, myCase)

		fmt.Print("Give Z : ")
		Z, _ := reader_1.ReadString('\n')
		sawmillsCount, _ = strconv.Atoi(strings.Replace(Z, "\r\n", "", -1))
	}

	for i, cs := range Cases {
		perms_squence := ""
		fmt.Println("Case", i)
		maxPfrofit := 0
		for _, sawmill := range cs.sawmills {
			profit := findMaxProfit(sawmill)
			if profit > maxPfrofit {
				maxPfrofit = profit
			}

			maxProfitPerms := findPermutionWithMaxProfit(sawmill)
			perms_squence += strings.Trim(strings.Join(strings.Fields(fmt.Sprint(maxProfitPerms)), " "), "") + ","
		}
		fmt.Println("Max Profit :", maxPfrofit)
		fmt.Println("Order:", perms_squence)
	}
}
