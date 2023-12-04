package main

import (
	"bufio"
	"os"
	"strconv"
	"unicode"
)

func part1() int {
	//4361 test output
	puzzleInput := "input.txt"

	file, _ := os.Open(puzzleInput)

	buffer := bufio.NewScanner(file)

	sum := 0
	inputArray := []string{}

	for buffer.Scan() {
		line := buffer.Text()
		inputArray = append(inputArray, line)
		// println(line)
	}
	for i := 0; i < len(inputArray); i++ {
		var look []byte
		validWord := []bool{}
		// println(inputArray[i])
		for j := 0; j < len(inputArray[i]); j++ {
			if unicode.IsDigit(rune(inputArray[i][j])) {
				look = append(look, inputArray[i][j])
				isValid := checkValid(i, j, inputArray)
				validWord = append(validWord, isValid)
			} else {
				if len(look) > 0 && !checkWord(validWord) {
					// println(string(look))
					num, _ := strconv.Atoi(string(look))
					sum += num
				}
				look = []byte{}
				validWord = []bool{}
			}
		}
		if len(look) > 0 && !checkWord(validWord) {
			// println(string(look))
			num, _ := strconv.Atoi(string(look))
			sum += num
		}
	}

	return sum
}

// This is a monstrosity, but it works lol
func checkValid(i int, j int, inputArray []string) bool {

	if i > 0 {
		if inputArray[i-1][j] != '.' && !unicode.IsDigit(rune(inputArray[i-1][j])) {
			return false
		}
	}
	if i < len(inputArray)-1 {
		if inputArray[i+1][j] != '.' && !unicode.IsDigit(rune(inputArray[i+1][j])) {
			return false
		}
	}
	if j > 0 {
		if inputArray[i][j-1] != '.' && !unicode.IsDigit(rune(inputArray[i][j-1])) {
			return false
		}
	}
	if j < len(inputArray[i])-1 {
		if inputArray[i][j+1] != '.' && !unicode.IsDigit(rune(inputArray[i][j+1])) {
			return false
		}
	}
	if i > 0 && j > 0 {
		if inputArray[i-1][j-1] != '.' && !unicode.IsDigit(rune(inputArray[i-1][j-1])) {
			return false
		}
	}
	if i > 0 && j < len(inputArray[i])-1 {
		if inputArray[i-1][j+1] != '.' && !unicode.IsDigit(rune(inputArray[i-1][j+1])) {
			return false
		}
	}
	if i < len(inputArray)-1 && j > 0 {
		if inputArray[i+1][j-1] != '.' && !unicode.IsDigit(rune(inputArray[i+1][j-1])) {
			return false
		}
	}
	if i < len(inputArray)-1 && j < len(inputArray[i])-1 {
		if inputArray[i+1][j+1] != '.' && !unicode.IsDigit(rune(inputArray[i+1][j+1])) {
			return false
		}
	}

	return true
}

func checkWord(valid []bool) bool {
	for i := 0; i < len(valid); i++ {
		if !valid[i] {
			return false
		}
	}
	return true
}

func main() {
	println("\nPart 1: ", part1())

}
