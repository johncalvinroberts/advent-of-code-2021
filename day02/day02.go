package day02

import (
	"github.com/johncalvinroberts/advent-of-code-2021/utils"
)

func Part1(input string) int {
	rows := utils.StrToSlice(input, "\n")
	x := 0
	y := 0
	for _, v := range rows {
		direction, amtStr := utils.StrToTuple(v, " ")
		amt := utils.StrToInt(amtStr, 0)

		if direction == "forward" {
			x += amt
		}

		if direction == "down" {
			y += amt
		}

		if direction == "up" {
			y -= amt
		}

	}
	return x * y
}

func Part2(input string) int {
	rows := utils.StrToSlice(input, "\n")
	x := 0 //
	y := 0
	aim := 0
	for _, v := range rows {
		direction, amtStr := utils.StrToTuple(v, " ")
		amt := utils.StrToInt(amtStr, 0)
		if direction == "down" {
			aim += amt
		}

		if direction == "up" {
			aim -= amt
		}

		if direction == "forward" {
			x += amt
			y += aim * amt
		}
	}
	return x * y
}
