package day05

import (
	"testing"

	"github.com/johncalvinroberts/advent-of-code-2021/utils"
)

var fixture = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

func TestPart1(t *testing.T) {
	r := Part1(fixture)
	utils.Assert(r, 5, t)
}
