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

	for buffer.Scan() {
		line := buffer.Text()
		println(line)
	}
	return sum
}
func main() {
	println("Part 1: ", part1())
}
