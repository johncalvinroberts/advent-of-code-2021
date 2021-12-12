package day08

import (
	"fmt"

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

/*
1,4,7,8  are unique

0:6
1:2
2:5
3:5
4:4
5:5
6:6
7:3
8:7
9:6

map of digit will be 9 blocks
*/

func Part2(input string) int {
	rows := utils.StrToSlice(input, "\n")
	for _, row := range rows {

		strs := utils.StrToSlice(row, " | ")
		readings := utils.StrToSlice(strs[0], " ")
		output := utils.StrToSlice(strs[1], " ")
		fmt.Printf("readings %+q, output %+q\n", readings, output)
	}

	return 0
}
