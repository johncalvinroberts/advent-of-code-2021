package day09

import (
	"sort"

	"github.com/johncalvinroberts/advent-of-code-2021/utils"
)

func Part1(input string) int {
	rows := utils.StrToSlice(input, "\n")

	memo := make([][]int, len(rows))

	total := 0

	for i, v := range rows {

		var row []int
		var prevRow []int
		var nextRow []int

		if i == 0 {
			ints := utils.StrSliceToIntSlice(utils.StrToSlice(v, ""))
			memo[i] = ints
			row = ints
		} else {
			row = memo[i]
		}

		if i > 0 {
			prevRow = memo[i-1]
		}

		if i < len(rows)-1 {
			nextRow = utils.StrSliceToIntSlice(utils.StrToSlice(rows[i+1], ""))
			memo[i+1] = nextRow
		}

		for j, v := range row {
			var adjacents []int
			if j != 0 {
				adjacents = append(adjacents, row[j-1])
			}

			if j != len(row)-1 {
				adjacents = append(adjacents, row[j+1])
			}

			if len(prevRow) > 0 {
				adjacents = append(adjacents, prevRow[j])
			}

			if len(nextRow) > 0 {
				adjacents = append(adjacents, nextRow[j])
			}
			sort.Slice(adjacents, func(k, t int) bool {
				return adjacents[k] < adjacents[t]
			})

			if v < adjacents[0] {
				total += v + 1
			}

		}

	}

	return total
}
