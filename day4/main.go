package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type input struct {
	numbers []int
	boards  [][][]int
}

func getInput() input {
	raw, _ := ioutil.ReadFile("input.txt")
	data := strings.Split(string(raw), "\n\n")

	numbers := []int{}
	boards := [][][]int{}
	for _, v := range data {
		if len(numbers) == 0 {
			for _, x := range strings.Split(v, ",") {
				num, _ := strconv.Atoi(x)
				numbers = append(numbers, num)
			}

			continue
		}

		board := [][]int{}
		for _, vv := range strings.Split(v, "\n") {
			row := []int{}
			for _, vvv := range strings.Split(vv, " ") {
				if len(vvv) > 0 {
					num, _ := strconv.Atoi(vvv)
					row = append(row, num)
				}
			}
			if len(row) > 0 {
				board = append(board, row)
			}
		}

		boards = append(boards, board)
	}

	return input{
		numbers: numbers,
		boards:  boards,
	}
}

func contains(arr []int, val int) bool {
	for _, b := range arr {
		if b == val {
			return true
		}
	}
	return false
}

func indexOf(arr []int, val int) int {
	for i, b := range arr {
		if b == val {
			return i
		}
	}

	return -1
}

func checkBingo(numbers []int, boards [][][]int) []int {
	bingoes := []int{}
	for k, board := range boards {
		for _, row := range board {
			matchesInRow := 0
			for _, value := range row {
				if contains(numbers, value) {
					matchesInRow++
				}
			}
			if matchesInRow == len(row) && !contains(bingoes, k) {
				bingoes = append(bingoes, k)
				break
			}
		}

		// Columns
		for i := 0; i < len(board[0]); i++ {
			col := []int{}
			for j := 0; j < len(board[0]); j++ {
				col = append(col, board[j][i])
			}

			matchesInCol := 0
			for _, value := range col {
				if contains(numbers, value) {
					matchesInCol++
				}
			}

			if matchesInCol == len(col) && !contains(bingoes, k) {
				bingoes = append(bingoes, k)
				break
			}
		}
	}

	return bingoes
}

func main() {
	input := getInput()

	drawn := []int{}
	for _, num := range input.numbers {
		drawn = append(drawn, num)
		bingoes := checkBingo(drawn, input.boards)
		if len(bingoes) > 0 {
			sumUnused := 0
			for _, row := range input.boards[bingoes[0]] {
				for _, value := range row {
					if !contains(drawn, value) {
						sumUnused = sumUnused + value
					}
				}
			}

			fmt.Printf("Part 1 answer: %d \n", sumUnused*drawn[len(drawn)-1])

			break
		}
	}

	drawn = []int{}
	bingoedBoards := []int{}
	var lastBingoAt int

	for _, num := range input.numbers {
		drawn = append(drawn, num)
		bingo := checkBingo(drawn, input.boards)
		for _, v := range bingo {
			if !contains(bingoedBoards, v) {
				lastBingoAt = drawn[len(drawn)-1]
				bingoedBoards = append(bingoedBoards, v)
			}
		}
	}

	finalDrawn := drawn[0 : indexOf(drawn, lastBingoAt)+1]
	lastBoard := input.boards[bingoedBoards[len(bingoedBoards)-1]]
	lastSumUnused := 0

	for _, row := range lastBoard {
		for _, value := range row {
			if !contains(finalDrawn, value) {
				lastSumUnused = lastSumUnused + value
			}
		}
	}

	fmt.Printf("Part 2 answer: %d \n", lastSumUnused*finalDrawn[len(finalDrawn)-1])
}
