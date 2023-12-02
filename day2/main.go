package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	id    int
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

		check := [3]string{"red", "blue", "green"}

		result := strings.Split(line, ":")

		id, _ := strconv.Atoi(strings.TrimSpace(strings.SplitAfter(result[0], " ")[1]))
		// println(id)

		game := Game{id: id, red: 0, blue: 0, green: 0}

		tries := strings.Split(result[1], ";")
		for i := 0; i < len(tries); i++ {
			tries[i] = strings.TrimSpace(tries[i])
			cube := strings.Split(tries[i], ",")
			for i := 0; i < len(cube); i++ {
				cube[i] = strings.TrimSpace(cube[i])
				roll := strings.Split(cube[i], " ")
				for i := 0; i < len(check); i++ {
					if roll[1] == check[i] {
						if check[i] == "red" {
							red, _ := strconv.Atoi(roll[0])
							if game.red < red {
								game.red = red
							}
						} else if check[i] == "blue" {
							blue, _ := strconv.Atoi(roll[0])
							if game.blue < blue {
								game.blue = blue
							}
						} else if check[i] == "green" {
							green, _ := strconv.Atoi(roll[0])
							if game.green < green {
								game.green = green
							}
						}
					}
				}
			}

			// print("\n")
			// println(tries[i])

		}

		// println("\nid:", game.id)
		// println("red:", game.red, "\nblue:", game.blue, "\ngreen:", game.green)

		if valid_check(game) {
			sum += game.id
		}

		// print("\n")
		// println(line)
	}
	return sum
}

func part2() int {
	// Test input should be: 2286
	puzzleInput := "input.txt"

	file, _ := os.Open(puzzleInput)

	buffer := bufio.NewScanner(file)

	sum := 0

	for buffer.Scan() {
		line := buffer.Text()

		check := [3]string{"red", "blue", "green"}

		result := strings.Split(line, ":")

		id, _ := strconv.Atoi(strings.TrimSpace(strings.SplitAfter(result[0], " ")[1]))

		game := Game{id: id, red: 0, blue: 0, green: 0}

		tries := strings.Split(result[1], ";")
		for i := 0; i < len(tries); i++ {
			tries[i] = strings.TrimSpace(tries[i])
			cube := strings.Split(tries[i], ",")
			for i := 0; i < len(cube); i++ {
				cube[i] = strings.TrimSpace(cube[i])
				roll := strings.Split(cube[i], " ")
				for i := 0; i < len(check); i++ {
					if roll[1] == check[i] {
						if check[i] == "red" {
							red, _ := strconv.Atoi(roll[0])
							if game.red < red {
								game.red = red
							}
						} else if check[i] == "blue" {
							blue, _ := strconv.Atoi(roll[0])
							if game.blue < blue {
								game.blue = blue
							}
						} else if check[i] == "green" {
							green, _ := strconv.Atoi(roll[0])
							if game.green < green {
								game.green = green
							}
						}
					}
				}
			}
		}
		// one line change, did the accidental setup in part 1 lol
		sum += game.red * game.blue * game.green
	}
	return sum
}

func valid_check(game Game) bool {
	if game.red <= 12 && game.blue <= 14 && game.green <= 13 {
		return true
	}
	return false

}

func main() {
	println("Part 1: ", part1())
	println("Part 2: ", part2())
}
