package day03

import (
	"strconv"

	"github.com/johncalvinroberts/advent-of-code-2021/utils"
)

func Part1(input string) int {
	rows := utils.StrToSlice(input, "\n")
	var gammaStr string
	var epsilonStr string
	bits := len(rows[0])

	for i := 0; i < bits; i++ {
		count0 := 0
		count1 := 0
		for _, v := range rows {
			if string(v[i]) == "0" {
				count0++
			} else {
				count1++
			}
		}
		if count0 > count1 {
			gammaStr += "0"
			epsilonStr += "1"
		} else {
			gammaStr += "1"
			epsilonStr += "0"
		}
	}

	gamma, err := strconv.ParseInt(gammaStr, 2, 64)
	if err != nil {
		panic("Failed to generate gamma int from binary")
	}
	epsilon, err := strconv.ParseInt(epsilonStr, 2, 64)
	if err != nil {
		panic("Failed to generate epsilon int from binary")
	}

	return int(gamma * epsilon)
}
