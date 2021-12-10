package day08

import (
	"github.com/johncalvinroberts/advent-of-code-2021/utils"
)

// identify how many times 1,4,7,8 appear in total
// 1 is 2, 4 is 4, 7 is 3, 8 is 7
func Part1(input string) int {
	rows := utils.StrToSlice(input, "\n")
	count := 0
	for _, row := range rows {
		numbers := utils.StrToSlice(utils.StrToSlice(row, "|")[1], " ")
		for _, str := range numbers {
			l := len(str)
			if l == 2 || l == 4 || l == 3 || l == 7 {
				count++
			}
		}
	}
	return count
}

func Part2(input string) int {

	return 0
}
