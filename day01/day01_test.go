package day01

import (
	"testing"

	"github.com/johncalvinroberts/advent-of-code-2021/utils"
)

var part1Fixture string = `199
200
208
210
200
207
240
269
260
263
`

var part2Fixture string = `
`

func TestPart1(t *testing.T) {
	result := Part1(part1Fixture)
	utils.Assert(result, 7, t)
}

func TestPart2(t *testing.T) {
	result := Part2(part1Fixture)
	utils.Assert(result, 5, t)
}
