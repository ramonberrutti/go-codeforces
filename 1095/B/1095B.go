package main

import (
	"fmt"
)

func instability(in []uint) uint {
	if len(in) == 0 {
		return 0
	}

	var min, pmax, max uint = in[0], in[0], in[0]

	for i := range in {
		if max < in[i] {
			pmax = max
			max = in[i]
		} else if pmax < in[i] {
			pmax = in[i]
		} else if min > in[i] {
			min = in[i]
		}
	}

	return pmax - min
}

func main() {
	var arraysize int
	fmt.Scanf("%d\n", &arraysize)
	arrayNumbers := make([]uint, arraysize)

	for i := range arrayNumbers {
		fmt.Scanf("%d", &arrayNumbers[i])
	}

	fmt.Print(instability(arrayNumbers))
}
