package day01

import (
	"fmt"

	"github.com/johncalvinroberts/advent-of-code-2021/utils"
)

// find the number of items which are an increase over previous item
func Part1(input string) int {
	readings := utils.StrToSlice(input, "\n")
	count := 0
	for i := 1; i < len(readings); i++ {
		if utils.StrToInt(readings[i-1], 0) < utils.StrToInt(readings[i], 0) {
			count++
		}
	}
	return count
}

// find the number of 3-measurement-windows which are an increase over the previous
func Part2(input string) int {
	readings := utils.StrToSlice(input, "\n")
	count := 0
	fmt.Println("PRIIIIIIINTIIING")
	fmt.Println(len(readings) - 2)
	// key is window, value is the sum
	windows := make(map[int]int)
	for i := 0; i < len(readings)-2; i++ {
		windowStr := readings[i : i+3]
		for j := 0; j < len(windowStr); j++ {
			windows[i] += utils.StrToInt(windowStr[j], 0)
		}
		if i == 0 {
			continue
		}
		if windows[i] > windows[i-1] {
			count++
		}
	}
	return count
}
