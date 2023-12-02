package main

import (
	"bufio"
	"os"
	"strings"
)

type Roll struct {
	red   int
	blue  int
	green int
}

func part1() int {
	puzzleInput := "input.txt"

	file, _ := os.Open(puzzleInput)

	buffer := bufio.NewScanner(file)

	sum := 0

	for buffer.Scan() {
		line := buffer.Text()

		result := strings.Split(line, ":")

		for _, slice := range result {

			println(slice)
		}

		// println(line)
	}
	return sum
}

func main() {
	println("Part 1: ", part1())
}
