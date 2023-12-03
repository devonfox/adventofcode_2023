package main

import (
	"bufio"
	"os"
)

func part1() int {
	puzzleInput := "input_sample.txt"

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

	}
	return sum
}

func main() {
	println("Part 1: ", part1())

}
