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

func calculateFish(input string, count int) int {
	fish := utils.StrSliceToIntSlice(utils.StrToSlice(input, ","))
	for i := 0; i < count; i++ {
		fmt.Printf("Day: %d, len%d\n", i, len(fish))
		for j, v := range fish {
			if v == 0 {
				fish[j] = 6
				fish = append(fish, 8)
			} else {
				fish[j]--
			}
		}
	}
	return len(fish)
}
