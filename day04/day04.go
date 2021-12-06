package day04

import (
	"strings"

	"github.com/johncalvinroberts/advent-of-code-2021/utils"
)

type Board [][]BoardItem

type BoardItem struct {
	value  int
	marked bool
}

const BoardDimension = 5

func Part1(input string) int {
	calls, boards := parse(input)
	var winningBoard Board = nil
	var lastCall int
	for _, n := range calls {
		for _, board := range boards {
			board.mark(n)
			if board.isWin() {
				winningBoard = *board
				lastCall = n
				break
			}
		}
		if winningBoard != nil {
			break
		}
	}
	sum := winningBoard.sumUnmarked()
	return sum * lastCall
}

func Part2(input string) int {
	calls, boards := parse(input)

	for _, n := range calls {
		var j int
		for _, board := range boards {
			board.mark(n)
			if !board.isWin() {
				boards[j] = board
				j++
			}
			if len(boards) == 1 {
				return boards[0].sumUnmarked() * n
			}
		}
		boards = boards[:j]
	}
	return 0
}

func (b Board) isWin() bool {
row:
	for i := 0; i < BoardDimension; i++ {
		markedCount := 0
		for _, item := range b[i] {
			if item.marked {
				markedCount++
			}
		}
		if markedCount == 5 {
			return true
		} else {
			continue row
		}
	}
column:
	for i := 0; i < BoardDimension; i++ {
		markedCount := 0
		for _, row := range b {
			if row[i].marked {
				markedCount++
			}
		}
		if markedCount == 5 {
			return true
		} else {
			continue column
		}
	}
	return false
}

func (b Board) mark(n int) {
	for x, row := range b {
		for y, item := range row {
			if item.value == n {
				b[x][y].marked = true
			}
		}
	}
}

func (b Board) sumUnmarked() int {
	sum := 0
	for _, row := range b {
		for _, item := range row {
			if !item.marked {
				sum += item.value
			}
		}
	}
	return sum
}

func parse(input string) (calls []int, boards []*Board) {
	rows := utils.StrToSlice(input, "\n\n")
	calls = utils.StrSliceToIntSlice(utils.StrToSlice(rows[0], ","))
	// massage boards into map of ints with bool for checked
	boards = make([]*Board, len(rows)-1)

	for i, boardstr := range rows[1:] {
		board := make(Board, 0)
		row := strings.Fields(boardstr)
		for y := 0; y < len(row); y += BoardDimension {
			chunk := row[y : y+BoardDimension]
			chunkInts := utils.StrSliceToIntSlice(chunk)
			rowStructs := make([]BoardItem, 0)
			for _, n := range chunkInts {
				rowStructs = append(rowStructs, BoardItem{value: n, marked: false})
			}
			board = append(board, rowStructs)
		}
		boards[i] = &board
	}
	return calls, boards
}
