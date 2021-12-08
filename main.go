package main

import (
	"fmt"
	"os"

	"github.com/johncalvinroberts/advent-of-code-2021/day01"
	"github.com/johncalvinroberts/advent-of-code-2021/day02"
	"github.com/johncalvinroberts/advent-of-code-2021/day03"
	"github.com/johncalvinroberts/advent-of-code-2021/day04"
	"github.com/johncalvinroberts/advent-of-code-2021/day05"
	"github.com/johncalvinroberts/advent-of-code-2021/day06"
	"github.com/johncalvinroberts/advent-of-code-2021/utils"
)

// usage: go run main.go <NN>
// assumes input is in day<NN>/input.txt
func main() {
	dayStr := os.Args[1]
	day := utils.StrToInt(dayStr, 25)
	fmt.Printf("🌲 Running day %02d\n", day)

	switch day {
	case 1:
		fmt.Printf("part 1: %d\n", day01.Part1(utils.Readfile(day)))
		fmt.Printf("part 2: %d\n", day01.Part2(utils.Readfile(day)))
	case 2:
		fmt.Printf("part 1: %d\n", day02.Part1(utils.Readfile(day)))
		fmt.Printf("part 2: %d\n", day02.Part2(utils.Readfile(day)))
	case 3:
		fmt.Printf("part 1: %d\n", day03.Part1(utils.Readfile(day)))
		// fmt.Printf("part 2: %d\n", day02.Part2(utils.Readfile(day)))
	case 4:
		fmt.Printf("part 1: %d\n", day04.Part1(utils.Readfile(day)))
		fmt.Printf("part 2: %d\n", day04.Part2(utils.Readfile(day)))
	case 5:
		fmt.Printf("part 1: %d\n", day05.Part1(utils.Readfile(day)))
		// fmt.Printf("part 2: %d\n", day04.Part2(utils.Readfile(day)))
	case 6:
		// fmt.Printf("part 1: %d\n", day06.Part1(utils.Readfile(day)))
		fmt.Printf("part 2: %d\n", day06.Part2(utils.Readfile(day)))
	}
}
