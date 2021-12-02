package day02

import (
	"testing"

	"github.com/johncalvinroberts/advent-of-code-2021/utils"
)

var fixture string = `forward 5
down 5
forward 8
up 3
down 8
forward 2`

func TestPart1(t *testing.T) {
	r := Part1(fixture)
	utils.Assert(r, 150, t)
}

func TestPart2(t *testing.T) {
	r := Part2(fixture)
	utils.Assert(r, 900, t)
}
