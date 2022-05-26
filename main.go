package main

// returns cut lenght, remained trunck length, remained cutter length
func cut(trunk int, cutterLength int) (int, int, int) {

	if cutterLength >= trunk {
		return trunk, 0, cutterLength - trunk
	} else {
		return cutterLength, trunk - cutterLength, 3
	}
}

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
