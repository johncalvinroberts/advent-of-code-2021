package day09

import (
	"testing"

	"github.com/johncalvinroberts/advent-of-code-2021/utils"
)

var fixture = `2199943210
3987894921
9856789892
8767896789
9899965678`

func TestPart1(t *testing.T) {
	r := Part1(fixture)
	utils.Assert(r, 15, t)
}
