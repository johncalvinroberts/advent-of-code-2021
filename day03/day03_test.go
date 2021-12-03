package day03

import (
	"testing"

	"github.com/johncalvinroberts/advent-of-code-2021/utils"
)

var fixture string = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

func TestPart1(t *testing.T) {
	r := Part1(fixture)
	utils.Assert(r, 198, t)
}
