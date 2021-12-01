package day01

import (
	"github.com/johncalvinroberts/advent-of-code-2021/utils"
)

// find the number of items which are an increase over previous item
func Part1(input string) int {
	readings1 := utils.StrToSlice(input, "\n")
	readings := utils.StrSliceToIntSlice(readings1)
	count := 0
	for i := 1; i < len(readings); i++ {
		if readings[i-1] < readings[i] {
			count++
		}
	}
	return count
}

// find the number of 3-measurement-windows which are an increase over the previous
func Part2(input string) int {
	readings := utils.StrSliceToIntSlice(utils.StrToSlice(input, "\n"))
	count := 0
	for i := 1; i+2 < len(readings); i++ {
		previous := readings[i-1] + readings[i] + readings[i+1]
		current := readings[i] + readings[i+1] + readings[i+2]
		if current > previous {
			count++
		}
	}
	return count
}
