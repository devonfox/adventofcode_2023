package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	id      int
	winning []int
	board   []int
	matches int
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
		numbers := strings.Split(result[1], " | ")
		numbers[0] = strings.TrimSpace(numbers[0])

		winningStrings := strings.Split(numbers[0], " ")
		winningNumbers := convertToIntArray(winningStrings)

		boardStrings := strings.Split(numbers[1], " ")
		boardNumbers := convertToIntArray(boardStrings)

		id, _ := strconv.Atoi(front[1])

		card := Card{id: id, winning: winningNumbers, board: boardNumbers, matches: 0}

		sum += gradeCard(&card)
	}
	return sum
}

func gradeCard(card *Card) int {
	var matches []int
	wins := 0
	for i := 0; i < len(card.winning); i++ {
		for j := 0; j < len(card.board); j++ {

			if card.winning[i] == card.board[j] {
				_ = append(matches, card.winning[i])
				wins++
				break
			}

		}
	}
	score := 0
	card.matches = wins

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

func part2() int {
	// test expected output: 30
	puzzleInput := "input.txt"

	file, _ := os.Open(puzzleInput)

	buffer := bufio.NewScanner(file)

	sum := 0

	cards := []Card{}

	for buffer.Scan() {
		line := buffer.Text()
		result := strings.Split(line, ": ")
		result[0] = strings.TrimSpace(result[0])
		result[1] = strings.TrimSpace(result[1])
		front := strings.Split(result[0], "d ")
		front[1] = strings.TrimSpace(front[1])

		numbers := strings.Split(result[1], " | ")
		numbers[0] = strings.TrimSpace(numbers[0])
		winningStrings := strings.Split(numbers[0], " ")
		winningNumbers := convertToIntArray(winningStrings)
		boardStrings := strings.Split(numbers[1], " ")
		boardNumbers := convertToIntArray(boardStrings)
		id, _ := strconv.Atoi(front[1])

		card := Card{id: id, winning: winningNumbers, board: boardNumbers, matches: 0}
		_ = gradeCard(&card)
		cards = append(cards, card)

	}

	sum = copySum(cards)

	return sum
}

func copySum(cards []Card) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	copyMap := make(map[int]int)
	sum := 0
	for i := 0; i < len(cards); i++ {
		copyMap[cards[i].id] = 1
	}

	for i := 0; i < len(cards); i++ {
		// println("id:", cards[i].id, "copies:", copyMap[cards[i].id])
		if cards[i].matches != 0 {
			cardCount, ok := copyMap[cards[i].id]
			if ok {
				bound := min(i+cards[i].matches+1, len(cards))
				for j := i + 1; j < bound; j++ {
					copyMap[cards[j].id] += cardCount
				}
			}

		}
	}
	for _, value := range copyMap {
		// println("key:", key, "value:", value)
		sum += value
	}

	return sum
}

func main() {
	println("\nPart 1:", part1())
	println("Part 2:", part2())
}
