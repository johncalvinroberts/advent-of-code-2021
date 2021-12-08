package day07

import (
	"sort"

	"github.com/johncalvinroberts/advent-of-code-2021/utils"
)

func Part1(input string) int {
	crabs := utils.StrSliceToIntSlice(utils.StrToSlice(input, ","))
	sort.Slice(crabs, func(i, j int) bool {
		return crabs[i] < crabs[j]
	})
	// find median
	middle := len(crabs) / 2
	var median int
	if len(crabs)%2 == 0 {
		median = (crabs[middle-1] + crabs[middle]) / 2
	} else {
		median = crabs[middle]
	}
	fuelCost := 0
	for _, v := range crabs {
		cost := utils.Absolute(median - v)
		fuelCost += cost
	}
	return fuelCost
}
