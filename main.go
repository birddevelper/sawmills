package main

import "fmt"

// returns cut lenght, remained trunck length, remained cutter length
func cut(trunk int, cutterLength int) (int, int, int) {

	if cutterLength >= trunk {
		return trunk, 0, cutterLength - trunk
	} else {
		return cutterLength, trunk - cutterLength, 3
	}
}

func ProfitCalculator(trunks []int) {

	currentcutterLength := 3
	for i, trunk := range trunks {

		trunkRemainedLength := trunk
		for trunkRemainedLength > 0 {
			cutLength, trunkRemainedLength, currentcutterLength := cut(trunkRemainedLength, currentcutterLength)
			cutterLength = cutterLength - trunk
			fmt.Println(i, trunk)
		}

	}
}
