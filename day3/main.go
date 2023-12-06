package main

import (
	"bufio"
	"os"
	"strconv"
	"unicode"
)

type Location struct {
	x int
	y int
}

type Number struct {
	value int
	gear  Location
}

type Check struct {
	valid bool
	gear  Location
}

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

func part2() int {
	// 467835 test output
	puzzleInput := "input.txt"
	file, _ := os.Open(puzzleInput)
	buffer := bufio.NewScanner(file)
	sum := 0
	inputArray := []string{}
	gearArray := []Number{}

	for buffer.Scan() {
		line := buffer.Text()
		inputArray = append(inputArray, line)
	}

	for i := 0; i < len(inputArray); i++ {
		var look []byte
		validWord := []Check{}
		for j := 0; j < len(inputArray[i]); j++ {
			if unicode.IsDigit(rune(inputArray[i][j])) {
				look = append(look, inputArray[i][j])
				isGear := checkGear(i, j, inputArray)
				if isGear != nil {
					validWord = append(validWord, Check{valid: true, gear: *isGear})
				} else {
					validWord = append(validWord, Check{valid: false, gear: Location{x: -1, y: -1}})
				}
			} else {
				if len(look) > 0 && checkWordGear(validWord) != nil {
					num, _ := strconv.Atoi(string(look))
					gearArray = append(gearArray, Number{value: num, gear: *checkWordGear(validWord)})
				}
				look = []byte{}
				validWord = []Check{}
			}
		}
		if len(look) > 0 && checkWordGear(validWord) != nil {
			num, _ := strconv.Atoi(string(look))
			gearArray = append(gearArray, Number{value: num, gear: *checkWordGear(validWord)})
		}
	}

	gears := make(map[Location][]int)

	for i := 0; i < len(gearArray); i++ {
		// println(gearArray[i].value, gearArray[i].gear.x, gearArray[i].gear.y)
		gears[gearArray[i].gear] = append(gears[gearArray[i].gear], gearArray[i].value)

	}
	product := 1
	for _, value := range gears {

		if len(value) == 2 {
			// println(key.x, key.y)
			product = value[0] * value[1]
			sum += product
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

func checkGear(i, j int, inputArray []string) *Location {

	check := func(row, col int) bool {
		if row >= 0 && row < len(inputArray) && col >= 0 && col < len(inputArray[row]) {
			char := inputArray[row][col]
			return char != '.' && isAsterisk(char)
		}
		return false
	}

	if check(i-1, j) {
		return &Location{x: i - 1, y: j}
	} else if check(i+1, j) {
		return &Location{x: i + 1, y: j}
	} else if check(i, j-1) {
		return &Location{x: i, y: j - 1}
	} else if check(i, j+1) {
		return &Location{x: i, y: j + 1}
	} else if check(i-1, j-1) {
		return &Location{x: i - 1, y: j - 1}
	} else if check(i-1, j+1) {
		return &Location{x: i - 1, y: j + 1}
	} else if check(i+1, j-1) {
		return &Location{x: i + 1, y: j - 1}
	} else if check(i+1, j+1) {
		return &Location{x: i + 1, y: j + 1}
	}

	return nil
}

func checkWord(valid []bool) bool {
	for i := 0; i < len(valid); i++ {
		if !valid[i] {
			return false
		}
	}
	return true
}

func checkWordGear(check []Check) *Location {
	for i := 0; i < len(check); i++ {
		if check[i].valid {
			return &Location{x: check[i].gear.x, y: check[i].gear.y}
		}
	}
	return nil
}

func isAsterisk(char byte) bool {
	return char == '*'
}

func main() {
	println("\nPart 1: ", part1())
	println("Part 2: ", part2())

}
