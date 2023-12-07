package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	id      int
	winning []int
	board   []int
}

func part1() int {
	// test expected output: 13
	puzzleInput := "input.txt"

	file, _ := os.Open(puzzleInput)

	buffer := bufio.NewScanner(file)

	sum := 0

	for buffer.Scan() {
		line := buffer.Text()
		result := strings.Split(line, ": ")
		result[0] = strings.TrimSpace(result[0])
		result[1] = strings.TrimSpace(result[1])
		front := strings.Split(result[0], "d ")
		front[1] = strings.TrimSpace(front[1])

		// println(result[1])
		numbers := strings.Split(result[1], " | ")
		// println(numbers[0])
		// printlnnumbers[1])
		numbers[0] = strings.TrimSpace(numbers[0])
		// numbers[1] = strings.TrimSpace(numbers[1])

		winningStrings := strings.Split(numbers[0], " ")
		winningNumbers := convertToIntArray(winningStrings)

		boardStrings := strings.Split(numbers[1], " ")
		boardNumbers := convertToIntArray(boardStrings)

		id, _ := strconv.Atoi(front[1])

		card := Card{id: id, winning: winningNumbers, board: boardNumbers}
		print(fmt.Sprintf("Card %d: ", card.id))
		for _, value := range card.winning {
			print(fmt.Sprintf("%d ", value))
		}
		print(" | ")
		for _, value := range card.board {
			print(fmt.Sprintf("%d ", value))
		}
		println()
		sum += gradeCard(card)

	}
	return sum
}

func gradeCard(card Card) int {
	var matches []int
	wins := 0
	for i := 0; i < len(card.winning); i++ {
		for j := 0; j < len(card.board); j++ {

			if card.winning[i] == card.board[j] {
				matches = append(matches, card.winning[i])
				wins++
				break
			}

		}
	}
	score := 0

	for i := 0; i < len(matches); i++ {
		println(matches[i])
	}
	// println("wins:", wins)
	// println("score:", score)
	if wins == 0 {
		return score
	}

	for i := 0; i < wins; i++ {
		if i == 0 {
			score = 1
		} else {
			score = score * 2
		}
	}
	// println("score:", score)
	return score
}

func convertToIntArray(stringArray []string) []int {

	for i := 0; i < len(stringArray); i++ {
		stringArray[i] = strings.TrimSpace(stringArray[i])
	}

	var result []int
	for _, value := range stringArray {
		x, _ := strconv.Atoi(value)
		// println(x)

		// lol this check is lazy af
		// this is just a patch for the trailing space
		if x != 0 {
			result = append(result, x)
		}
	}
	return result
}

func main() {
	println("\nPart 1:", part1())
}
