package main

import (
	"bufio"
	"os"
	"strconv"
	"unicode"
)

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
			// compiler complained about for byte value, so I casted it to rune
			// todo: look into this
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
	return 0
}

func main() {
	println("Part 1: ", part1())
	println("Part 2: ", part2())
}
