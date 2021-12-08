package day07

import (
	"testing"

	"github.com/johncalvinroberts/advent-of-code-2021/utils"
)

var fixture = `16,1,2,0,4,2,7,1,2,14`

func TestPart1(t *testing.T) {
	r := Part1(fixture)
	utils.Assert(r, 37, t)
}
