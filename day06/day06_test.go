package day06

import (
	"testing"

	"github.com/johncalvinroberts/advent-of-code-2021/utils"
)

var fixture = `3,4,3,1,2`

func TestPart1(t *testing.T) {
	r := Part1(fixture)
	utils.Assert(r, 5934, t)
}
