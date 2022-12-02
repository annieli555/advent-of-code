package main

import (
	"fmt"
)

func main() {
	input, _ := UrlToLines("https://adventofcode.com/2022/day/2/input")
	// file, _ := os.Open("./input.txt")
	// defer file.Close()

	// scanner := bufio.NewScanner(file)
	// var input []string

	// for scanner.Scan() {
	// 	input = append(input, scanner.Text())
	// }

	totalScore := 0

	for i := 0; i < len(input); i++ {
		// get opponent's move
		opponent := input[i][0] - 'A'

		// get my move
		me := input[i][2] - 'X'

		// add my move
		totalScore += int(me) + 1

		// calculate if i won, draw, or lost
		// me - opponent = 0 --> draw
		// me - opponent = 1 --> win
		// me - opponent = -1 --> lose
		diff := (me - opponent + 1 + 3) % 3

		totalScore += int(diff) * 3
	}
	fmt.Println(totalScore)
}
