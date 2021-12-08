package day05

import (
	"github.com/johncalvinroberts/advent-of-code-2021/utils"
	"github.com/johncalvinroberts/advent-of-code-2021/vector"
)

func Part1(input string) int {
	rowStrs := utils.StrToSlice(input, "\n")
	grid := make(map[vector.XY]int)
	for _, v := range rowStrs {
		coordinates := utils.StrSliceToIntSlice(utils.ExtractIntsToStrSlice(v))
		x1 := coordinates[0]
		y1 := coordinates[1]
		x2 := coordinates[2]
		y2 := coordinates[3]
		if x1 == x2 || y1 == y2 {
			p1, p2 := vector.XY{X: x1, Y: y1}, vector.XY{X: x2, Y: y2}
			direction := p2.Subtract(p1).Sign()
			for p := p1; p != p2; p = p.Add(direction) {
				grid[p]++
			}
			grid[p2]++
		}
	}
	dangerCount := 0
	for _, count := range grid {
		if count > 1 {
			dangerCount++
		}
	}
	return dangerCount
}
