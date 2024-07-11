package main

import (
	"math/rand"
	"slices"
)

type number []int

func getNumbers(count int) number {
	// num := [] int {1, 2, 3, 4, 5, 6, 7, 8, 9}

	var randomNumber []int

	for i := 0; i < count; i++ {
		newNum := rangeIn(0, 9)
		if(slices.Contains(randomNumber, newNum)) {
			continue
		} else {
			randomNumber =  append(randomNumber, rangeIn(0, 9))
		}
	}

	return randomNumber
}

func rangeIn(low, hi int) int {
    return rand.Intn(hi-low)
}