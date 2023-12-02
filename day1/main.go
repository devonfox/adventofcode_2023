package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type Index struct {
	index int
	name  string
}

func part1() int {
	puzzleInput := "input.txt"

	// could check err here, skipping for today
	file, _ := os.Open(puzzleInput)

	buffer := bufio.NewScanner(file)
	first := ""
	second := ""
	sum := 0
	// lookin at each line from the file
	for buffer.Scan() {

		word := buffer.Text()
		for i := 0; i < len(word); i++ {
			// compiler complained about byte value, suggested casting to rune
			// todo: look into this 'rune' lol
			if unicode.IsDigit(rune(word[i])) {
				first = string(word[i])
				break
			}
		}
		for i := len(word) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(word[i])) {
				second = string(word[i])
				break
			}
		}
		lineNum, _ := strconv.Atoi(first + second)
		sum += lineNum

	}
	file.Close()
	return sum

}

func part2() int {

	puzzleInput := "input.txt"

	check := [10]string{" ", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	file, _ := os.Open(puzzleInput)

	buffer := bufio.NewScanner(file)

	sum := 0

	for buffer.Scan() {
		var order []Index
		word := buffer.Text()

		for i := 0; i < len(check); i++ {
			start := findBounds(word, check[i])
			if start != -1 {
				order = append(order, Index{start, check[i]})
				// println(check[i], "found at", start)
			}
		}

		for i := 0; i < len(word); i++ {
			if unicode.IsDigit(rune(word[i])) {
				// println(string(word[i]), "found at", i)
				index, _ := strconv.Atoi(string(word[i]))
				order = append(order, Index{i, check[index]})
			}

		}

		sort.Slice(order, func(i, j int) bool {
			return order[i].index < order[j].index
		})
		// for i := 0; i < len(order); i++ {
		// 	println(order[i].index, " ", order[i].name)
		// }

		first, last := "", ""

		for i := 0; i < len(check); i++ {
			if check[i] == order[0].name {
				first = strconv.Itoa(i)
				break
			}
		}
		if len(order) > 1 {
			for i := 0; i < len(check); i++ {
				if check[i] == order[len(order)-1].name {
					last = strconv.Itoa(i)
					// print("last:", last)
					break
				}
			}
		}
		result := ""

		if len(order) > 1 {
			result = first + last
		} else {
			result = first + first
		}
		// println(result)
		lineNum, _ := strconv.Atoi(result)
		// println(lineNum)
		sum += lineNum
		// print("\n")

	}

	file.Close()
	return sum
}

func findBounds(input, toFind string) int {
	start := strings.Index(input, toFind)
	if start == -1 {
		return -1
	}
	return start
}

func main() {
	println("Part 1: ", part1())
	println("Part 2: ", part2())
}
