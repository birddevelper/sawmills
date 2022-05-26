package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// returns cut lenght, remained trunck length, remained cutter length
func cut(trunk int, cutterLength int) (int, int, int) {
	// if trunk length is less than cutrer, return whole trunk length, otherwise retuen cutterLength
	if cutterLength >= trunk {
		return trunk, 0, cutterLength - trunk
	} else {
		return cutterLength, trunk - cutterLength, 3
	}
}

// calculate the profit of a permutation of trunks
func ProfitCalculator(trunks []int) int {
	totalProfit := 0
	//at the first the cutter has to cut 3 block
	currentcutterLength := 3
	for _, trunk := range trunks {
		// at the first, remained trunk length is whole the trunk
		trunkRemainedLength := trunk
		for trunkRemainedLength > 0 {
			var cutLength int
			// do cut the trunk
			cutLength, trunkRemainedLength, currentcutterLength = cut(trunkRemainedLength, currentcutterLength)

			// give bounce or penalty on cutted trunck length
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

// check if two slices of int are the same
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func isAlreadyExisted(permuts [][]int, permut []int) bool {

	for _, perm := range permuts {
		if Equal(perm, permut) {
			return true
		}
	}

	return false
}

// permutation of trunks thrown in river
func permutation(trunks []int) (permuts [][]int) {
	var rc func([]int, int)
	// calculate all permutations of the trunks sequence
	rc = func(a []int, k int) {
		if k == len(a) {
			if !isAlreadyExisted(permuts, a) {
				permuts = append(permuts, append([]int{}, a...))
			}
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

// find maximum profit between all permutation of a trunk sequence
func findMaxProfit(trunks []int) int {
	// calculate permutation
	trunksPerms := permutation(trunks)
	maxProfit := 0
	//find maximum profit
	for i := 0; i < len(trunksPerms); i++ {
		profit := ProfitCalculator(trunksPerms[i])
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}

// find permutation which has profit equal to maximum profit
func findPermutionWithMaxProfit(trunks []int) [][]int {
	var permutionWithMaxProfit [][]int
	// find maximum profit of this sequence
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
	// define cases
	type Case struct {
		CaseNumber int
		sawmills   [][]int
	}

	var Cases []Case
	// read from input
	reader_1 := bufio.NewReader(os.Stdin)
	fmt.Print("Give Z : ")
	Z, _ := reader_1.ReadString('\n')
	caseNo := 0
	// convert read string to int
	sawmillsCount, _ := strconv.Atoi(strings.Replace(Z, "\r\n", "", -1))

	// Take Input until Z = 0
	for sawmillsCount != 0 {

		caseNo += 1
		var sawmills [][]int
		for i := 0; i < sawmillsCount; i++ {
			// for each sawmill read the trunk sequence from input
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

		// get Z from input for next round, you can pass 0 to finish input
		fmt.Print("Give Z : ")
		Z, _ := reader_1.ReadString('\n')
		sawmillsCount, _ = strconv.Atoi(strings.Replace(Z, "\r\n", "", -1))
	}

	// calculating the requested output
	for i, cs := range Cases {
		perms_squence := ""
		fmt.Println("Case", i)

		// sort sawmills trunks based on maximum profit
		sort.Slice(cs.sawmills, func(i, j int) bool {

			return findMaxProfit(cs.sawmills[i]) > findMaxProfit(cs.sawmills[j])
		})

		// max profit of first sawmills trunks sequence is maximum profit between all sawmills
		maxPfrofit := findMaxProfit(cs.sawmills[0])
		for _, sawmill := range cs.sawmills {
			maxProfitPerms := findPermutionWithMaxProfit(sawmill)
			perms_squence += strings.Trim(strings.Join(strings.Fields(fmt.Sprint(maxProfitPerms)), " "), "") + ","
		}

		// print result
		fmt.Println("Max Profit :", maxPfrofit)
		fmt.Println("Order:", perms_squence)
	}
}
