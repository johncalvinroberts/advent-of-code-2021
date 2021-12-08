package day06

import (
	"fmt"

	"github.com/johncalvinroberts/advent-of-code-2021/utils"
)

func Part1(input string) int {
	return calculateFish(input, 80)
}

func Part2(input string) int {
	return calculateFish(input, 256)
}

func calculateFish(input string, days int) int {
	fish := utils.StrSliceToIntSlice(utils.StrToSlice(input, ","))
	// instead of using a huge list, use a map of indexes
	// each index stores the number of fish at that age
	var counts [9]int
	for _, n := range fish {
		counts[n]++
	}
	for i := 0; i < days; i++ {

		var nextCounts [9]int
		for ttl, num := range counts {
			if ttl == 0 {
				nextCounts[6] += num
				nextCounts[8] += num
			} else {
				nextCounts[ttl-1] += num
			}
		}
		counts = nextCounts
	}
	fmt.Print("")
	sum := 0

	for _, count := range counts {
		sum += count
	}
	return sum
}
